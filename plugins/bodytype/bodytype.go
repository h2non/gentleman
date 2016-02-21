package bodytype

import (
	c "gopkg.in/h2non/gentleman.v0/context"
	p "gopkg.in/h2non/gentleman.v0/plugin"
	"net/http"
)

// Types is a map of MIME type aliases
var Types = map[string]string{
	"html":       "text/html",
	"json":       "application/json",
	"xml":        "application/xml",
	"text":       "text/plain",
	"urlencoded": "application/x-www-form-urlencoded",
	"form":       "application/x-www-form-urlencoded",
	"form-data":  "application/x-www-form-urlencoded",
}

// Type defines an authorization basic header in the outgoing request
func Type(name string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		defineType(name, ctx.Request)
		h.Next(ctx)
	})
}

func defineType(name string, req *http.Request) {
	match := Types[name]
	if match == "" {
		match = name
	}
	req.Header.Set("Content-Type", match)
}
