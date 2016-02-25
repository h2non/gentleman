package gentleman

import (
	"gopkg.in/h2non/gentleman.v0/context"
)

// Dispatcher dispatches a given request triggering the middleware
// layer and handling the request/response state.
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

	// Run the request middleware
	ctx = d.req.Middleware.Run("request", ctx)
	// Handle error
	if ctx.Error != nil {
		ctx = d.req.Middleware.Run("error", ctx)
		if ctx.Error != nil {
			return ctx
		}
	}

	// Verify if the request was intercepted
	if ctx.Response.StatusCode != 0 {
		// Then trigger the response middleware
		ctx = d.req.Middleware.Run("response", ctx)
		if ctx.Error != nil {
			ctx = d.req.Middleware.Run("error", ctx)
		}
		return ctx
	}

	// Perform the request via ctx.Client
	res, err := ctx.Client.Do(ctx.Request)
	ctx.Error = err
	if err != nil {
		ctx = d.req.Middleware.Run("error", ctx)
		if ctx.Error != nil {
			return ctx
		}
	}
	if res != nil {
		ctx.Response = res
	}

	// Run the response middleware
	ctx = d.req.Middleware.Run("response", ctx)
	if ctx.Error != nil {
		ctx = d.req.Middleware.Run("error", ctx)
	}

	return ctx
}
