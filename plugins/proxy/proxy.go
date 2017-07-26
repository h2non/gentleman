package proxy

import (
	c "gopkg.in/h2non/gentleman.v2/context"
	p "gopkg.in/h2non/gentleman.v2/plugin"
	"net/http"
	"net/url"
)

// Set defines the proxy servers to be used based on the transport scheme
func Set(servers map[string]string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		// Assert http.Transport to work with the instance
		transport, ok := ctx.Client.Transport.(*http.Transport)
		if !ok {
			// If using a custom transport, just ignore it
			h.Next(ctx)
			return
		}

		// Define the proxy function to be used during the transport
		transport.Proxy = func(req *http.Request) (*url.URL, error) {
			if value, ok := servers[req.URL.Scheme]; ok {
				return url.Parse(value)
			}
			return http.ProxyFromEnvironment(req)
		}

		// Override the transport
		ctx.Client.Transport = transport
		h.Next(ctx)
	})
}
