package mux

import (
	c "gopkg.in/h2non/gentleman.v0/context"
	"regexp"
)

// Matcher represent the function interface implemented by matchers
type Matcher func(ctx *c.Context) bool

// Method returns a new multiplexer who matches an HTTP request based on the given method.
func Method(name string) *Mux {
	return Match(func(ctx *c.Context) bool {
		return ctx.GetString("$phase") == "request" && ctx.Request.Method == name
	})
}

// Methods returns a new multiplexer who matches an HTTP request based on the given methods.
func Methods(methods []string) *Mux {
	return Match(func(ctx *c.Context) bool {
		if ctx.GetString("$phase") != "request" {
			return false
		}
		for _, method := range methods {
			if ctx.Request.Method == method {
				return true
			}
		}
		return false
	})
}

// Path returns a new multiplexer who matches an HTTP request
// path based on the given regexp pattern.
func Path(pattern string) *Mux {
	return Match(func(ctx *c.Context) bool {
		if ctx.GetString("$phase") != "request" {
			return false
		}
		matched, _ := regexp.MatchString(pattern, ctx.Request.URL.Path)
		return matched
	})
}

// URL returns a new multiplexer who matches an HTTP request
// URL based on the given regexp pattern.
func URL(pattern string) *Mux {
	return Match(func(ctx *c.Context) bool {
		if ctx.GetString("$phase") != "request" {
			return false
		}
		matched, _ := regexp.MatchString(pattern, ctx.Request.URL.String())
		return matched
	})
}

// Host returns a new multiplexer who matches an HTTP request
// URL host based on the given regexp pattern.
func Host(pattern string) *Mux {
	return Match(func(ctx *c.Context) bool {
		if ctx.GetString("$phase") != "request" {
			return false
		}
		matched, _ := regexp.MatchString(pattern, ctx.Request.URL.Host)
		return matched
	})
}

// Query returns a new multiplexer who matches an HTTP request
// query param based on the given key and regexp pattern.
func Query(key, pattern string) *Mux {
	return Match(func(ctx *c.Context) bool {
		if ctx.GetString("$phase") != "request" {
			return false
		}
		matched, _ := regexp.MatchString(pattern, ctx.Request.URL.Query().Get(key))
		return matched
	})
}

// RequestHeader returns a new multiplexer who matches an HTTP request
// header field based on the given key and regexp pattern.
func RequestHeader(key, pattern string) *Mux {
	return Match(func(ctx *c.Context) bool {
		if ctx.GetString("$phase") != "request" {
			return false
		}
		matched, _ := regexp.MatchString(pattern, ctx.Request.Header.Get(key))
		return matched
	})
}

// ResponseHeader returns a new multiplexer who matches an HTTP response
// header field based on the given key and regexp pattern.
func ResponseHeader(key, pattern string) *Mux {
	return Match(func(ctx *c.Context) bool {
		if ctx.GetString("$phase") != "response" {
			return false
		}
		matched, _ := regexp.MatchString(pattern, ctx.Response.Header.Get(key))
		return matched
	})
}

// Status returns a new multiplexer who matches an HTTP response
// status code based on the given status.
func Status(code int) *Mux {
	return Match(func(ctx *c.Context) bool {
		return ctx.GetString("$phase") == "response" && ctx.Response.StatusCode == code
	})
}

// StatusRange returns a new multiplexer who matches an HTTP response
// status code based on the given status range, including both numbers.
func StatusRange(start, end int) *Mux {
	return Match(func(ctx *c.Context) bool {
		return ctx.GetString("$phase") == "response" && ctx.Response.StatusCode >= start && ctx.Response.StatusCode <= end
	})
}
