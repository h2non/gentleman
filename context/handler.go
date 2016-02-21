package context

// HandlerCtx represents a Context only function handler
type HandlerCtx func(*Context)

// HandlerFunc represents a middleware function handler interface
type HandlerFunc func(*Context, Handler)

// Handler exposes a simple interface providing
// control flow capabilities to middleware functions
type Handler interface {
	// Next handler invokes the next plugin in the middleware
	// call chain with the given Context
	Next(*Context)

	// Stop handler stops the current middleware call chain
	// with the given Context
	Stop(*Context)

	// Error handler reports an error and stops the middleware
	// call chain with the given Context
	Error(*Context, error)
}

// handler encapsulates a HandlerCtx function
type handler struct {
	next HandlerCtx
}

// NewHandler creates a new Handler based
// on a given HandlerCtx function
func NewHandler(fn HandlerCtx) Handler {
	return handler{once(fn)}
}

// Next continues executing the next middleware
// function in the call chain
func (h handler) Next(ctx *Context) {
	h.next(ctx)
}

// Error reports an error and stops the middleware call chain
func (h handler) Error(ctx *Context, err error) {
	ctx.Error = err
	h.next(ctx)
}

// Stop stops the middleware call chain with a custom Context
func (h handler) Stop(ctx *Context) {
	ctx.Stopped = true
	h.next(ctx)
}

// once returns a function that can be executed once
// Subsequent calls will be no-op
func once(fn HandlerCtx) HandlerCtx {
	called := false
	return func(ctx *Context) {
		if called {
			return
		}
		called = true
		fn(ctx)
	}
}
