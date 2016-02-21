package mux

import (
	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v0/context"
	"testing"
)

func TestMuxSimple(t *testing.T) {
	mx := New()
	mx.UseRequest(func(ctx *context.Context, h context.Handler) {
		ctx.Request.Header.Set("foo", "bar")
		h.Next(ctx)
	})
	ctx := context.New()
	mx.Run("request", ctx)
	st.Expect(t, ctx.Request.Header.Get("foo"), "bar")
}

func TestMuxMethodMatcher(t *testing.T) {
	mx := New()
	mx.Use(Method("GET").UseRequest(func(ctx *context.Context, h context.Handler) {
		ctx.Request.Header.Set("foo", "bar")
		h.Next(ctx)
	}))
	ctx := context.New()
	ctx.Request.Method = "GET"
	mx.Run("request", ctx)
	st.Expect(t, ctx.Request.Header.Get("foo"), "bar")
}

type handler struct {
	fn     context.Handler
	called bool
}

func newHandler() *handler {
	h := &handler{}
	h.fn = context.NewHandler(func(c *context.Context) {
		h.called = true
	})
	return h
}
