package mux

import (
	c "gopkg.in/h2non/gentleman.v0/context"
)

// Matcher represent the function interface implemented by matchers
type Matcher func(ctx *c.Context) bool

// Method return a new multiplexer who matches an incoming HTTP request based on the given method.
func Method(name string) *Mux {
	return Match(func(ctx *c.Context) bool {
		return ctx.GetString("$phase") == "request" && ctx.Request.Method == name
	})
}
