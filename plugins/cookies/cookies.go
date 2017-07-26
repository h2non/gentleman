package cookies

import (
	"golang.org/x/net/publicsuffix"
	c "gopkg.in/h2non/gentleman.v2/context"
	p "gopkg.in/h2non/gentleman.v2/plugin"
	"net/http"
	"net/http/cookiejar"
)

// Add adds a cookie to the request. Per RFC 6265 section 5.4, AddCookie does not
// attach more than one Cookie header field.
// That means all cookies, if any, are written into the same line, separated by semicolon.
func Add(cookie *http.Cookie) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		ctx.Request.AddCookie(cookie)
		h.Next(ctx)
	})
}

// Set sets a new cookie field by key and value.
func Set(key, value string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		cookie := &http.Cookie{Name: key, Value: value}
		ctx.Request.AddCookie(cookie)
		h.Next(ctx)
	})
}

// DelAll deletes all the cookies by deleting the Cookie header field.
func DelAll() p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		ctx.Request.Header.Del("Cookie")
		h.Next(ctx)
	})
}

// SetMap sets a map of cookies represented by key-value pair.
func SetMap(cookies map[string]string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		for k, v := range cookies {
			cookie := &http.Cookie{Name: k, Value: v}
			ctx.Request.AddCookie(cookie)
		}
		h.Next(ctx)
	})
}

// AddMultiple adds a list of cookies.
func AddMultiple(cookies []*http.Cookie) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		for _, cookie := range cookies {
			ctx.Request.AddCookie(cookie)
		}
		h.Next(ctx)
	})
}

// Jar creates a cookie jar to store HTTP cookies when they are sent down.
func Jar() p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
		ctx.Client.Jar = jar
		h.Next(ctx)
	})
}
