package transport

import (
	c "gopkg.in/h2non/gentleman.v2/context"
	p "gopkg.in/h2non/gentleman.v2/plugin"
	"net/http"
)

// Set sets a new HTTP transport for the outgoing request
func Set(transport http.RoundTripper) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		// Override the http.Client transport
		ctx.Client.Transport = transport
		h.Next(ctx)
	})
}
