package compression

import (
	c "gopkg.in/h2non/gentleman.v2/context"
	p "gopkg.in/h2non/gentleman.v2/plugin"
	"net/http"
)

// Disable disables the authorization basic header in the outgoing request
func Disable() p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		// Assert http.Transport to work with the instance
		transport, ok := ctx.Client.Transport.(*http.Transport)
		if !ok {
			h.Next(ctx)
			return
		}

		// Override the http.Client transport
		transport.DisableCompression = true
		ctx.Client.Transport = transport

		h.Next(ctx)
	})
}
