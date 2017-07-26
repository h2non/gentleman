package query

import (
	c "gopkg.in/h2non/gentleman.v2/context"
	p "gopkg.in/h2non/gentleman.v2/plugin"
)

// Set sets the query param key and value.
// It replaces any existing values.
func Set(key, value string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		query := ctx.Request.URL.Query()
		query.Set(key, value)
		ctx.Request.URL.RawQuery = query.Encode()
		h.Next(ctx)
	})
}

// Add adds the query param value to key.
// It appends to any existing values associated with key.
func Add(key, value string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		query := ctx.Request.URL.Query()
		query.Add(key, value)
		ctx.Request.URL.RawQuery = query.Encode()
		h.Next(ctx)
	})
}

// Del deletes the query param values associated with key.
func Del(key string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		query := ctx.Request.URL.Query()
		query.Del(key)
		ctx.Request.URL.RawQuery = query.Encode()
		h.Next(ctx)
	})
}

// DelAll deletes all the query params.
func DelAll() p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		ctx.Request.URL.RawQuery = ""
		h.Next(ctx)
	})
}

// SetMap sets a map of query params by key-value pair.
func SetMap(params map[string]string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		query := ctx.Request.URL.Query()
		for k, v := range params {
			query.Set(k, v)
		}
		ctx.Request.URL.RawQuery = query.Encode()
		h.Next(ctx)
	})
}
