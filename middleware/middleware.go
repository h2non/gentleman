// Package middleware implements an HTTP client domain-specific phase-oriented
// middleware layer used internally by gentleman packages.
package middleware

import (
	"sync"

	c "gopkg.in/h2non/gentleman.v1/context"
	"gopkg.in/h2non/gentleman.v1/plugin"
)

// Middleware especifies the required interface that must be
// implemented by middleware capable interfaces.
type Middleware interface {
	// Use method is used to register a new plugin in the middleware stack.
	Use(plugin.Plugin) Middleware

	// UseError is used to register a new error phase middleware function handler.
	UseError(c.HandlerFunc) Middleware

	// UseRequest is used to register a new request phase middleware function handler.
	UseRequest(c.HandlerFunc) Middleware

	// UseResposne is used to register a new resposne phase middleware function handler.
	UseResponse(c.HandlerFunc) Middleware

	// UseHandler is used to register a new phase specific middleware function handler.
	UseHandler(string, c.HandlerFunc) Middleware

	// Run is used to dispatch the middleware call chain for a specific phase.
	Run(string, *c.Context) *c.Context

	// UseParent defines a parent middleware for easy inheritance.
	UseParent(Middleware) Middleware

	// Clone is used to created a new clone of the existent middleware.
	Clone() Middleware

	// Flush is used to flush the middleware stack.
	Flush()

	// GetStack is used to retrieve an array of registered plugins.
	GetStack() []plugin.Plugin

	// SetStack is used to override the stack of registered plugins.
	SetStack([]plugin.Plugin)
}

// Layer type represent an HTTP domain
// specific middleware layer with inheritance support.
type Layer struct {
	// stack stores the plugins registered in the current middleware instance.
	stack []plugin.Plugin

	// parent points to a parent middleware for behavior inheritance.
	parent Middleware
}

// New creates a new middleware layer.
func New() *Layer {
	return &Layer{}
}

// Use registers a new plugin to the middleware stack.
func (s *Layer) Use(plugin plugin.Plugin) Middleware {
	s.stack = append(s.stack, plugin)
	return s
}

// UseHandler registers a phase specific plugin handler in the middleware stack.
func (s *Layer) UseHandler(phase string, fn c.HandlerFunc) Middleware {
	s.stack = append(s.stack, plugin.NewPhasePlugin(phase, fn))
	return s
}

// UseResponse registers a new response phase middleware handler.
func (s *Layer) UseResponse(fn c.HandlerFunc) Middleware {
	s.stack = append(s.stack, plugin.NewResponsePlugin(fn))
	return s
}

// UseRequest registers a new request phase middleware handler.
func (s *Layer) UseRequest(fn c.HandlerFunc) Middleware {
	s.stack = append(s.stack, plugin.NewRequestPlugin(fn))
	return s
}

// UseError registers a new error phase middleware handler.
func (s *Layer) UseError(fn c.HandlerFunc) Middleware {
	s.stack = append(s.stack, plugin.NewErrorPlugin(fn))
	return s
}

// UseParent attachs a parent middleware.
func (s *Layer) UseParent(parent Middleware) Middleware {
	s.parent = parent
	return s
}

// Flush flushes the plugins stack.
func (s *Layer) Flush() {
	s.stack = s.stack[:0]
}

// SetStack sets the middleware plugin stack overriding the existent one.
func (s *Layer) SetStack(stack []plugin.Plugin) {
	s.stack = stack
}

// GetStack gets the current middleware plugins stack.
func (s *Layer) GetStack() []plugin.Plugin {
	return s.stack
}

// Clone creates a new Middleware instance based on the current one.
func (s *Layer) Clone() Middleware {
	mw := New()
	mw.parent = s.parent
	mw.stack = append([]plugin.Plugin(nil), s.stack...)
	return mw
}

// Run triggers the middleware call chain for the given phase.
func (s *Layer) Run(phase string, ctx *c.Context) *c.Context {
	if s.parent != nil {
		ctx = s.parent.Run(phase, ctx)
		if ctx.Error != nil || ctx.Stopped {
			return ctx
		}
	}

	s.stack = filter(s.stack)
	return trigger(phase, s.stack, ctx)
}

func filter(stack []plugin.Plugin) []plugin.Plugin {
	buf := []plugin.Plugin{}
	for _, plugin := range stack {
		if !plugin.Removed() {
			buf = append(buf, plugin)
		}
	}
	return buf
}

// Note: this implementation may change in the future
func trigger(phase string, stack []plugin.Plugin, ctx *c.Context) *c.Context {
	var wg sync.WaitGroup
	wg.Add(1)

	// Finisher function
	done := func(_ctx *c.Context) {
		ctx = _ctx
		wg.Done()
	}

	i := len(stack)
	if i == 0 {
		wg.Done()
		return ctx
	}

	next := done
	for i > 0 {
		i--
		next = nextHandler(phase, stack[i], next, done)
	}

	// Exposes current middleware phase via context
	ctx.Set("$phase", phase)

	// Triggers the middleware call chain
	next(ctx)

	wg.Wait()
	return ctx
}

func nextHandler(phase string, plugin plugin.Plugin, next c.HandlerCtx, done c.HandlerCtx) c.HandlerCtx {
	return func(ctx *c.Context) {
		handler := c.NewHandler(eval(phase, next, done))
		plugin.Exec(phase, ctx, handler)
	}
}

func eval(phase string, next c.HandlerCtx, done c.HandlerCtx) c.HandlerCtx {
	return func(ctx *c.Context) {
		if phase == "error" {
			if ctx.Error == nil {
				done(ctx)
				return
			}
			next(ctx)
			return
		}

		if ctx.Error != nil || (ctx.Stopped && phase != "stop") {
			done(ctx)
			return
		}

		next(ctx)
	}
}
