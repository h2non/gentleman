package middleware

import (
	"errors"
	"testing"
	"time"

	"gopkg.in/h2non/gentleman.v2/context"
	"gopkg.in/h2non/gentleman.v2/plugin"
)

func TestCreateMiddleware(t *testing.T) {
	mw := New()
	mw.UseRequest(forward)
	mw.UseResponse(forward)
	mw.UseError(forward)
	if len(mw.GetStack()) != 3 {
		t.Error("Invalid stack size")
	}
}

func TestMiddlewareFlush(t *testing.T) {
	mw := New()
	mw.UseRequest(forward)
	if len(mw.GetStack()) != 1 {
		t.Error("Invalid stack size")
	}
	mw.Flush()
	if len(mw.GetStack()) != 0 {
		t.Error("Stack must be empty")
	}
}

func TestSimpleMiddlewareCallChain(t *testing.T) {
	mw := New()

	calls := 0
	fn := func(ctx *context.Context, h context.Handler) {
		calls++
		h.Next(ctx)
	}

	mw.UseRequest(fn)
	mw.UseResponse(fn)

	ctx := context.New()
	ctx = mw.Run("request", ctx)
	if ctx.Error != nil || calls != 1 {
		t.Errorf("Invalid middleware calls: %d => %s", calls, ctx.Error)
	}
	ctx = mw.Run("response", ctx)
	if ctx.Error != nil || calls != 2 {
		t.Errorf("Invalid middleware calls: %d => %s", calls, ctx.Error)
	}
}

func TestMultipleMiddlewareCallChain(t *testing.T) {
	mw := New()

	calls := 0
	fn := func(ctx *context.Context, h context.Handler) {
		calls++
		h.Next(ctx)
	}

	mw.UseRequest(fn)
	mw.UseRequest(fn)
	mw.UseRequest(fn)
	mw.UseRequest(fn)

	ctx := context.New()
	ctx = mw.Run("request", ctx)
	if ctx.Error != nil || calls != 4 {
		t.Errorf("Invalid middleware calls: %d => %s", calls, ctx.Error)
	}
}

func TestAsyncMiddlewareCallChain(t *testing.T) {
	mw := New()

	calls := 0
	fn := func(ctx *context.Context, h context.Handler) {
		go func(h context.Handler) {
			calls++
			time.Sleep(time.Millisecond * 25)
			h.Next(ctx)
		}(h)
	}

	mw.UseRequest(fn)
	mw.UseRequest(fn)
	mw.UseRequest(fn)
	mw.UseRequest(fn)

	ctx := context.New()
	ctx = mw.Run("request", ctx)
	if ctx.Error != nil || calls != 4 {
		t.Errorf("Invalid middleware calls: %d => %s", calls, ctx.Error)
	}
}

func TestMultipleHandlerCalls(t *testing.T) {
	mw := New()

	calls := 0
	fn := func(ctx *context.Context, h context.Handler) {
		calls++
		h.Next(ctx)
		h.Next(ctx) // ignored!
	}

	mw.UseRequest(fn)
	mw.UseRequest(fn)
	mw.UseRequest(fn)
	mw.UseRequest(fn)

	ctx := context.New()
	ctx = mw.Run("request", ctx)
	if ctx.Error != nil || calls != 4 {
		t.Errorf("Invalid middleware calls: %d => %s", calls, ctx.Error)
	}
}

func TestMiddlewarePlugin(t *testing.T) {
	calls := 0
	fn := func(ctx *context.Context, h context.Handler) {
		calls++
		h.Next(ctx)
	}
	plugin := plugin.New()
	plugin.SetHandler("request", fn)
	plugin.SetHandler("response", fn)
	plugin.SetHandler("error", fn)

	mw := New()
	mw.Use(plugin)

	ctx := context.New()
	mw.Run("request", ctx)
	mw.Run("error", ctx)
	mw.Run("response", ctx)

	if calls != 3 {
		t.Errorf("Invalid middleware calls: %d", calls)
	}
}

func TestMiddlewareContextSharing(t *testing.T) {
	fn := func(ctx *context.Context, h context.Handler) {
		ctx.Set("foo", ctx.GetString("foo")+"bar")
		h.Next(ctx)
	}

	mw := New()
	mw.UseRequest(fn)
	mw.UseRequest(fn)
	mw.UseRequest(fn)

	ctx := context.New()
	ctx = mw.Run("request", ctx)
	if val := ctx.GetString("foo"); val != "barbarbar" {
		t.Errorf("Invalid context value: %s", val)
	}
}

func TestMiddlewareInheritance(t *testing.T) {
	parent := New()
	child := New()
	child.UseParent(parent)
	mw := New()
	mw.UseParent(child)

	fn := func(c *context.Context, h context.Handler) {
		c.Set("foo", c.GetString("foo")+"bar")
		h.Next(c)
	}

	child.UseRequest(fn)
	parent.UseRequest(fn)
	mw.UseRequest(fn)

	ctx := context.New()
	ctx = mw.Run("request", ctx)

	if ctx.GetString("foo") != "barbarbar" {
		t.Error("Invalid context value")
	}
}

func TestMiddlewareNewContextPassing(t *testing.T) {
	mw := New()

	mw.UseRequest(func(c *context.Context, h context.Handler) {
		c.Set("foo", "bar")
		h.Next(c)
	})
	mw.UseRequest(func(c *context.Context, h context.Handler) {
		ctx := c.Clone()
		h.Next(ctx)
	})

	ctx := context.New()
	newCtx := mw.Run("request", ctx)

	if ctx.Error != nil {
		t.Error("Error must be empty")
	}
	if newCtx == ctx {
		t.Error("Context should not be equal")
	}
	if newCtx.Get("foo") == "foo" {
		t.Error("foo context value is not valid")
	}
}

func TestMiddlewareError(t *testing.T) {
	mw := New()

	fn := func(c *context.Context, h context.Handler) {
		t.Error("Should not call the handler")
		h.Next(c)
	}
	fnError := func(c *context.Context, h context.Handler) {
		h.Error(c, errors.New("Error"))
	}

	mw.UseRequest(fnError)
	mw.UseRequest(fn)

	ctx := context.New()
	ctx = mw.Run("request", ctx)

	if ctx.Error == nil {
		t.Error("Error must exists")
	}
}

func TestMiddlewareErrorPassing(t *testing.T) {
	mw := New()
	mw.UseRequest(func(c *context.Context, h context.Handler) {
		h.Error(c, errors.New("foo"))
	})
	mw.UseError(func(c *context.Context, h context.Handler) {
		h.Error(c, errors.New("Error: "+c.Error.Error()))
	})

	ctx := mw.Run("request", context.New())
	if ctx.Error.Error() != "foo" {
		t.Errorf("Invalid error value: %s", ctx.Error)
	}

	ctx = mw.Run("error", ctx)
	if ctx.Error.Error() != "Error: foo" {
		t.Errorf("Invalid error value: %s", ctx.Error)
	}
}

func TestMiddlewareStop(t *testing.T) {
	mw := New()

	fn := func(c *context.Context, h context.Handler) {
		t.Error("Should not call the handler")
		h.Next(c)
	}
	stop := func(c *context.Context, h context.Handler) {
		h.Stop(c)
	}

	mw.UseRequest(stop)
	mw.UseRequest(fn)

	ctx := context.New()
	ctx = mw.Run("request", ctx)

	if ctx.Error != nil {
		t.Error("Error must be empty")
	}
	if !ctx.Stopped {
		t.Error("Call chain should be stopped")
	}
}

func TestMiddlewareErrorWithInheritance(t *testing.T) {
	parent := New()
	mw := New()
	mw.UseParent(parent)

	fn := func(c *context.Context, h context.Handler) {
		t.Error("Should not call the handler")
		h.Next(c)
	}
	stop := func(c *context.Context, h context.Handler) {
		h.Stop(c)
	}

	mw.UseRequest(stop)
	mw.UseRequest(fn)

	ctx := context.New()
	ctx = mw.Run("request", ctx)

	if ctx.Error != nil {
		t.Error("Error must be empty")
	}
	if !ctx.Stopped {
		t.Error("Call chain should be stopped")
	}
}

func forward(ctx *context.Context, h context.Handler) {
	h.Next(ctx)
}
