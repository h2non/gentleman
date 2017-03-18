package url

import (
	c "gopkg.in/h2non/gentleman.v2/context"
	p "gopkg.in/h2non/gentleman.v2/plugin"
	"net/url"
	"regexp"
	"strings"
)

// URL parses and defines a new URL in the outgoing request
func URL(uri string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		u, err := url.Parse(normalize(uri))
		if err != nil {
			h.Error(ctx, err)
			return
		}

		ctx.Request.URL = u
		h.Next(ctx)
	})
}

// BaseURL parses and defines a schema and host URL values in the outgoing request
func BaseURL(uri string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		u, err := url.Parse(normalize(uri))
		if err != nil {
			h.Error(ctx, err)
			return
		}

		ctx.Request.URL.Scheme = u.Scheme
		ctx.Request.URL.Host = u.Host
		h.Next(ctx)
	})
}

// Path defines a new URL path in the outgoing request
func Path(path string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		ctx.Request.URL.Path = normalizePath(path)
		h.Next(ctx)
	})
}

// AddPath concatenates a path slice to the existent path in the outgoing request
func AddPath(path string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		ctx.Request.URL.Path += normalizePath(path)
		h.Next(ctx)
	})
}

// PathPrefix defines a path prefix to the existent path in the outgoing request
func PathPrefix(path string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		ctx.Request.URL.Path = normalizePath(path) + ctx.Request.URL.Path
		h.Next(ctx)
	})
}

// Param replaces one or multiple path param expressions by the given value
func Param(key, value string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		ctx.Request.URL.Path = replace(ctx.Request.URL.Path, key, value)
		h.Next(ctx)
	})
}

// Params replaces one or multiple path param expressions by the given map of key-value pairs
func Params(params map[string]string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		for key, value := range params {
			ctx.Request.URL.Path = replace(ctx.Request.URL.Path, key, value)
		}
		h.Next(ctx)
	})
}

func replace(str, key, value string) string {
	return strings.Replace(str, ":"+key, value, -1)
}

func normalizePath(path string) string {
	if path == "/" {
		return ""
	}
	return path
}

func normalize(uri string) string {
	match, _ := regexp.MatchString("^http[s]?://", uri)
	if match {
		return uri
	}
	return "http://" + uri
}
