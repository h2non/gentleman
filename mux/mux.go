// Package mux implements an HTTP domain-specific traffic multiplexer
// with built-in matchers and features for easy plugin composition and activable logic.
package mux

import (
	c "gopkg.in/h2non/gentleman.v2/context"
	"gopkg.in/h2non/gentleman.v2/middleware"
	"gopkg.in/h2non/gentleman.v2/plugin"
)

// Mux is a HTTP request/response/error multiplexer who implements both
// middleware and plugin interfaces.
// It has been designed for easy plugin composition based on HTTP matchers/filters.
type Mux struct {
	// Mux also implements a plugin capable interface.
	*plugin.Layer

	// Matchers stores a list of matcher functions.
	Matchers []Matcher

	// Middleware stores the multiplexer middleware layer.
	Middleware middleware.Middleware
}

// New creates a new multiplexer with default settings.
func New() *Mux {
	m := &Mux{Layer: plugin.New()}
	m.Middleware = middleware.New()
	handler := m.Handler()
	m.DefaultHandler = handler
	return m
}

// Match matches the give Context againts a list of matchers and
// returns `true` if all the matchers passed.
func (m *Mux) Match(ctx *c.Context) bool {
	for _, matcher := range m.Matchers {
		if !matcher(ctx) {
			return false
		}
	}
	return true
}

// AddMatcher adds a new matcher function in the current mumultiplexer matchers stack.
func (m *Mux) AddMatcher(matchers ...Matcher) *Mux {
	m.Matchers = append(m.Matchers, matchers...)
	return m
}

// Handler returns the function handler to match an incoming HTTP transacion
// and trigger the equivalent middleware phase.
func (m *Mux) Handler() c.HandlerFunc {
	return func(ctx *c.Context, h c.Handler) {
		if !m.Match(ctx) {
			h.Next(ctx)
			return
		}

		ctx = m.Middleware.Run(ctx.GetString("$phase"), ctx)
		if ctx.Error != nil {
			h.Error(ctx, ctx.Error)
			return
		}
		if ctx.Stopped {
			h.Stop(ctx)
			return
		}

		h.Next(ctx)
	}
}

// Use registers a new plugin in the middleware stack.
func (m *Mux) Use(p plugin.Plugin) *Mux {
	m.Middleware.Use(p)
	return m
}

// UseResponse registers a new response phase middleware handler.
func (m *Mux) UseResponse(fn c.HandlerFunc) *Mux {
	m.Middleware.UseResponse(fn)
	return m
}

// UseRequest registers a new request phase middleware handler.
func (m *Mux) UseRequest(fn c.HandlerFunc) *Mux {
	m.Middleware.UseRequest(fn)
	return m
}

// UseError registers a new error phase middleware handler.
func (m *Mux) UseError(fn c.HandlerFunc) *Mux {
	m.Middleware.UseError(fn)
	return m
}

// UseHandler registers a new error phase middleware handler.
func (m *Mux) UseHandler(phase string, fn c.HandlerFunc) *Mux {
	m.Middleware.UseHandler(phase, fn)
	return m
}

// UseParent attachs a parent middleware.
func (m *Mux) UseParent(parent middleware.Middleware) *Mux {
	m.Middleware.UseParent(parent)
	return m
}

// Flush flushes the plugins stack.
func (m *Mux) Flush() {
	m.Middleware.Flush()
}

// SetStack sets the middleware plugin stack overriding the existent one.
func (m *Mux) SetStack(stack []plugin.Plugin) {
	m.Middleware.SetStack(stack)
}

// GetStack gets the current middleware plugins stack.
func (m *Mux) GetStack() []plugin.Plugin {
	return m.Middleware.GetStack()
}

// Clone creates a new Middleware instance based on the current one.
func (m *Mux) Clone() middleware.Middleware {
	return m.Middleware.Clone()
}

// Run triggers the middleware call chain for the given phase.
func (m *Mux) Run(phase string, ctx *c.Context) *c.Context {
	return m.Middleware.Run(phase, ctx)
}
