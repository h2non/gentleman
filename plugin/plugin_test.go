package plugin

import (
	"testing"

	"gopkg.in/h2non/gentleman.v2/context"
)

func TestPluginLayer(t *testing.T) {
	phase := ""
	fn := func(c *context.Context, h context.Handler) {
		phase = c.GetString("$phase")
		h.Next(c)
	}

	ctx := context.New()
	plugin := New()
	plugin.SetHandler("request", fn)
	plugin.SetHandler("response", fn)
	plugin.SetHandler("error", fn)

	calls := 0
	createHandler := func() context.Handler {
		return context.NewHandler(func(c *context.Context) { calls++ })
	}

	ctx.Set("$phase", "request")
	plugin.Exec("request", ctx, createHandler())
	if phase != "request" {
		t.Errorf("Invalid phase: %s", phase)
	}

	ctx.Set("$phase", "response")
	plugin.Exec("response", ctx, createHandler())
	if phase != "response" {
		t.Errorf("Invalid phase: %s", phase)
	}

	ctx.Set("$phase", "error")
	plugin.Exec("error", ctx, createHandler())
	if phase != "error" {
		t.Errorf("Invalid phase: %s", phase)
	}

	if calls != 3 {
		t.Errorf("Invalid number of calls: %d", calls)
	}
}

func TestNewPhasePlugin(t *testing.T) {
	called := false
	plugin := NewPhasePlugin("foo", func(c *context.Context, h context.Handler) { h.Next(c) })
	plugin.Exec("foo", context.New(), context.NewHandler(func(c *context.Context) { called = true }))
	if !called {
		t.Errorf("Handler not called")
	}
}

func TestNewResponsePlugin(t *testing.T) {
	called := false
	plugin := NewResponsePlugin(func(c *context.Context, h context.Handler) { h.Next(c) })
	plugin.Exec("response", context.New(), context.NewHandler(func(c *context.Context) { called = true }))
	if !called {
		t.Errorf("Handler not called")
	}
}

func TestNewRequestPlugin(t *testing.T) {
	called := false
	plugin := NewRequestPlugin(func(c *context.Context, h context.Handler) { h.Next(c) })
	plugin.Exec("request", context.New(), context.NewHandler(func(c *context.Context) { called = true }))
	if !called {
		t.Errorf("Handler not called")
	}
}

func TestNewErrorPlugin(t *testing.T) {
	called := false
	plugin := NewErrorPlugin(func(c *context.Context, h context.Handler) { h.Next(c) })
	plugin.Exec("error", context.New(), context.NewHandler(func(c *context.Context) { called = true }))
	if !called {
		t.Errorf("Handler not called")
	}
}
