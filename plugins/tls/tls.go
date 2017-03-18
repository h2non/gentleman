package tls

import (
	"crypto/tls"
	c "gopkg.in/h2non/gentleman.v2/context"
	p "gopkg.in/h2non/gentleman.v2/plugin"
	"net/http"
)

// Config defines the request TLS connection config
func Config(config *tls.Config) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		// Assert http.Transport to work with the instance
		transport, ok := ctx.Client.Transport.(*http.Transport)
		if !ok {
			// If using a custom transport, just ignore it
			h.Next(ctx)
			return
		}

		// Override the http.Client transport
		transport.TLSClientConfig = config
		ctx.Client.Transport = transport

		h.Next(ctx)
	})
}
