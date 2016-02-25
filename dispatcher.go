package gentleman

import (
	"gopkg.in/h2non/gentleman.v0/context"
)

// Dispatcher dispatches a given request triggering the middleware
// layer per specific phase and handling the request/response/error
// states accondingly.
type Dispatcher struct {
	req *Request
}

// NewDispatcher creates a new Dispatcher based on the given Context.
func NewDispatcher(req *Request) *Dispatcher {
	return &Dispatcher{req}
}

// Dispatch triggers the middleware chains and performs the HTTP request.
func (d *Dispatcher) Dispatch() *context.Context {
	ctx := d.req.Context
	mw := d.req.Middleware

	// Run the request middleware
	ctx = mw.Run("request", ctx)
	// Handle error
	if ctx.Error != nil {
		ctx = mw.Run("error", ctx)
		if ctx.Error != nil {
			return ctx
		}
	}

	// Verify if the request was intercepted
	if ctx.Response.StatusCode != 0 {
		// Then trigger the response middleware
		ctx = mw.Run("response", ctx)
		if ctx.Error != nil {
			ctx = mw.Run("error", ctx)
		}
		return ctx
	}

	// If manually stopped
	if ctx.Stopped {
		ctx = mw.Run("stopped", ctx)
		if ctx.Error != nil {
			ctx = mw.Run("error", ctx)
			if ctx.Error != nil {
				return ctx
			}
		}
		if ctx.Stopped {
			return ctx
		}
	}

	// Trigger the before dial phase
	ctx = mw.Run("before dial", ctx)
	if ctx.Error != nil {
		ctx = mw.Run("error", ctx)
		if ctx.Error != nil {
			return ctx
		}
	}
	// If manually stopped
	if ctx.Stopped {
		ctx = mw.Run("stopped", ctx)
		if ctx.Error != nil {
			ctx = mw.Run("error", ctx)
			if ctx.Error != nil {
				return ctx
			}
		}
		if ctx.Stopped {
			return ctx
		}
	}

	// Perform the request via ctx.Client
	res, err := ctx.Client.Do(ctx.Request)
	ctx.Error = err
	if err != nil {
		ctx = mw.Run("error", ctx)
		if ctx.Error != nil {
			return ctx
		}
	}
	if res != nil {
		ctx.Response = res
	}

	// Trigger the after dial phase
	ctx = mw.Run("after dial", ctx)
	if ctx.Error != nil {
		ctx = mw.Run("error", ctx)
		if ctx.Error != nil {
			return ctx
		}
	}

	// If manually stopped
	if ctx.Stopped {
		ctx = mw.Run("stopped", ctx)
		if ctx.Error != nil {
			ctx = mw.Run("error", ctx)
			if ctx.Error != nil {
				return ctx
			}
		}
		if ctx.Stopped {
			return ctx
		}
	}

	// Run the response middleware
	ctx = mw.Run("response", ctx)
	if ctx.Error != nil {
		ctx = mw.Run("error", ctx)
	}

	return ctx
}
