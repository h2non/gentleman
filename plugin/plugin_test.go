package plugin

import (
	"gopkg.in/h2non/gentleman.v0/context"
	"testing"
)

func TestPluginLayer(t *testing.T) {
	phase := ""
	fn := func(c *context.Context, h context.Handler) {
		phase = c.GetString("phase")
		h.Next(c)
	}

	ctx := context.New()
	plugin := &Layer{false, false, fn, fn, fn}

	calls := 0
	createHandler := func() context.Handler {
		return context.NewHandler(func(c *context.Context) { calls++ })
	}

	ctx.Set("phase", "request")
	plugin.Request(ctx, createHandler())
	if phase != "request" {
		t.Errorf("Invalid phase: %s", phase)
	}

	ctx.Set("phase", "response")
	plugin.Response(ctx, createHandler())
	if phase != "response" {
		t.Errorf("Invalid phase: %s", phase)
	}

	ctx.Set("phase", "error")
	plugin.Error(ctx, createHandler())
	if phase != "error" {
		t.Errorf("Invalid phase: %s", phase)
	}

	if calls != 3 {
		t.Errorf("Invalid number of calls: %d", calls)
	}
}

func TestNewResponsePlugin(t *testing.T) {
	called := false
	plugin := NewResponsePlugin(func(c *context.Context, h context.Handler) { h.Next(c) })
	plugin.Response(context.New(), context.NewHandler(func(c *context.Context) { called = true }))
	if !called {
		t.Errorf("Handler not called")
	}
}

func TestNewRequestPlugin(t *testing.T) {
	called := false
	plugin := NewRequestPlugin(func(c *context.Context, h context.Handler) { h.Next(c) })
	plugin.Request(context.New(), context.NewHandler(func(c *context.Context) { called = true }))
	if !called {
		t.Errorf("Handler not called")
	}
}

func TestNewErrorPlugin(t *testing.T) {
	called := false
	plugin := NewErrorPlugin(func(c *context.Context, h context.Handler) { h.Next(c) })
	plugin.Error(context.New(), context.NewHandler(func(c *context.Context) { called = true }))
	if !called {
		t.Errorf("Handler not called")
	}
}
