package context

import (
	"errors"
	"testing"
)

func TestHandler(t *testing.T) {
	ctx := New()
	ctx.Set("foo", "bar")
	h := NewHandler(func(c *Context) {
		if c.Error != nil {
			t.Error("Context error should be empty")
		}
		if c.GetString("foo") != "bar" {
			t.Error("Invalid context value")
		}
	})
	h.Next(ctx)
}

func TestHandlerError(t *testing.T) {
	ctx := New()
	h := NewHandler(func(c *Context) {
		if c.Error == nil {
			t.Error("Context error cannot be empty")
		}
	})
	h.Error(ctx, errors.New("error"))
}

func TestHandlerStop(t *testing.T) {
	ctx := New()
	h := NewHandler(func(c *Context) {
		if c.Error != nil {
			t.Error("Context error should be empty")
		}
		if !c.Stopped {
			t.Error("Should be stopped")
		}
	})
	h.Stop(ctx)
}

func TestHandlerOnceExecution(t *testing.T) {
	ctx := New()
	calls := 0
	h := NewHandler(func(c *Context) { calls++ })
	h.Next(ctx)
	h.Next(ctx) // ignored!
	h.Stop(ctx) // ignored!
	if calls != 1 {
		t.Errorf("Invalid number of calls: %d", calls)
	}
}
