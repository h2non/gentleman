package auth

import (
	c "gopkg.in/h2non/gentleman.v2/context"
	p "gopkg.in/h2non/gentleman.v2/plugin"
)

// Basic defines an authorization basic header in the outgoing request
func Basic(username, password string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		ctx.Request.SetBasicAuth(username, password)
		h.Next(ctx)
	})
}

// Bearer defines an authorization bearer token header in the outgoing request
func Bearer(token string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		ctx.Request.Header.Set("Authorization", "Bearer "+token)
		h.Next(ctx)
	})
}

// Custom defines a custom authorization header field in the outgoing request
func Custom(value string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		ctx.Request.Header.Set("Authorization", value)
		h.Next(ctx)
	})
}
