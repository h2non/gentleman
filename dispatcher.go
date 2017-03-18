package gentleman

import (
	c "gopkg.in/h2non/gentleman.v2/context"
)

// Dispatcher dispatches a given request triggering the middleware
// layer per specific phase and handling the request/response/error
// states accondingly.
type Dispatcher struct {
	req *Request
}

// task function represents the required interface for dispatcher pipeline tasks.
type task func(*c.Context) (*c.Context, bool)

// NewDispatcher creates a new Dispatcher based on the given Context.
func NewDispatcher(req *Request) *Dispatcher {
	return &Dispatcher{req}
}

// Dispatch triggers the middleware chains and performs the HTTP request.
func (d *Dispatcher) Dispatch() *c.Context {
	// Pipeline of tasks to execute in FIFO order
	pipeline := []task{
		func(ctx *c.Context) (*c.Context, bool) {
			return d.runBefore("request", ctx)
		},
		func(ctx *c.Context) (*c.Context, bool) {
			return d.runBefore("before dial", ctx)
		},
		func(ctx *c.Context) (*c.Context, bool) {
			return d.doDial(ctx)
		},
		func(ctx *c.Context) (*c.Context, bool) {
			return d.runAfter("after dial", ctx)
		},
		func(ctx *c.Context) (*c.Context, bool) {
			return d.runAfter("response", ctx)
		},
	}

	// Reference to initial context
	ctx := d.req.Context

	// Execute tasks in order, stopping in case of error or explicit stop.
	for _, task := range pipeline {
		var stop bool
		if ctx, stop = task(ctx); stop {
			break
		}
	}

	return ctx
}

func (d *Dispatcher) doDial(ctx *c.Context) (*c.Context, bool) {
	// Perform the request via ctx.Client
	res, err := ctx.Client.Do(ctx.Request)
	ctx.Error = err
	if err != nil {
		ctx = d.req.Middleware.Run("error", ctx)
		if ctx.Error != nil {
			return ctx, true
		}
	}

	// Assign response if present
	if res != nil {
		ctx.Response = res
	}

	return ctx, false
}

func (d *Dispatcher) runBefore(phase string, ctx *c.Context) (*c.Context, bool) {
	// Run the request middleware
	ctx, stop := d.run(phase, ctx)
	if stop {
		return ctx, true
	}

	// Verify if the response was injected
	ctx, stop = d.intercepted(ctx)
	if stop {
		return ctx, true
	}

	// Verify if the should stop
	return d.stop(ctx)
}

func (d *Dispatcher) runAfter(phase string, ctx *c.Context) (*c.Context, bool) {
	// Trigger the after dial phase
	ctx, stop := d.run(phase, ctx)
	if stop {
		return ctx, true
	}

	// Verify if the should stop
	return d.stop(ctx)
}

func (d *Dispatcher) intercepted(ctx *c.Context) (*c.Context, bool) {
	// Verify if the request was intercepted
	if ctx.Response.StatusCode == 0 {
		return ctx, false
	}

	// Trigger the intercept middleware
	ctx, stop := d.run("intercept", ctx)
	if stop {
		return ctx, true
	}

	// Finally trigger the response middleware
	ctx, _ = d.run("response", ctx)
	return ctx, true
}

func (d *Dispatcher) stop(ctx *c.Context) (*c.Context, bool) {
	if !ctx.Stopped {
		return ctx, false
	}

	mw := d.req.Middleware
	ctx = mw.Run("stop", ctx)
	if ctx.Error != nil {
		ctx = mw.Run("error", ctx)
		if ctx.Error != nil {
			return ctx, true
		}
	}

	return ctx, ctx.Stopped
}

func (d *Dispatcher) run(phase string, ctx *c.Context) (*c.Context, bool) {
	mw := d.req.Middleware

	// Run the middleware by phase
	ctx = mw.Run(phase, ctx)
	if ctx.Error == nil {
		return ctx, false
	}

	// Run error middleware
	ctx = mw.Run("error", ctx)
	if ctx.Error != nil {
		return ctx, true
	}

	return ctx, false
}
