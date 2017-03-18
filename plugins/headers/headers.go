package headers

import (
	c "gopkg.in/h2non/gentleman.v2/context"
	p "gopkg.in/h2non/gentleman.v2/plugin"
)

// Set sets the header entries associated with key to the single element value.
// It replaces any existing values associated with key.
func Set(key, value string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		ctx.Request.Header.Set(key, value)
		h.Next(ctx)
	})
}

// Add adds the key, value pair to the header.
// It appends to any existing values associated with key.
func Add(key, value string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		ctx.Request.Header.Add(key, value)
		h.Next(ctx)
	})
}

// Del deletes the header fields associated with key.
func Del(key string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		ctx.Request.Header.Del(key)
		h.Next(ctx)
	})
}

// SetMap sets a map of headers represented by key-value pair.
func SetMap(headers map[string]string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		for k, v := range headers {
			ctx.Request.Header.Set(k, v)
		}
		h.Next(ctx)
	})
}
