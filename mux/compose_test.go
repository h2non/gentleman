package mux

import (
	"testing"

	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v2/context"
)

func TestMuxComposeIfMatches(t *testing.T) {
	mx := New()
	mx.Use(If(Method("GET"), Host("foo.com")).UseRequest(func(ctx *context.Context, h context.Handler) {
		ctx.Request.Header.Set("foo", "bar")
		h.Next(ctx)
	}))
	ctx := context.New()
	ctx.Request.Method = "GET"
	ctx.Request.URL.Host = "foo.com"
	mx.Run("request", ctx)
	st.Expect(t, ctx.Request.Header.Get("foo"), "bar")
}

func TestMuxComposeIfUnmatch(t *testing.T) {
	mx := New()
	mx.Use(If(Method("GET"), Host("bar.com")).UseRequest(func(ctx *context.Context, h context.Handler) {
		ctx.Request.Header.Set("foo", "bar")
		h.Next(ctx)
	}))
	ctx := context.New()
	ctx.Request.Method = "GET"
	ctx.Request.URL.Host = "foo.com"
	mx.Run("request", ctx)
	st.Expect(t, ctx.Request.Header.Get("foo"), "")
}

func TestMuxComposeOrMatch(t *testing.T) {
	mx := New()
	mx.Use(Or(Method("GET"), Host("bar.com")).UseRequest(func(ctx *context.Context, h context.Handler) {
		ctx.Request.Header.Set("foo", "bar")
		h.Next(ctx)
	}))
	ctx := context.New()
	ctx.Request.Method = "GET"
	ctx.Request.URL.Host = "foo.com"
	mx.Run("request", ctx)
	st.Expect(t, ctx.Request.Header.Get("foo"), "bar")
}

func TestMuxComposeOrUnMatch(t *testing.T) {
	mx := New()
	mx.Use(Or(Method("GET"), Host("bar.com")).UseRequest(func(ctx *context.Context, h context.Handler) {
		ctx.Request.Header.Set("foo", "bar")
		h.Next(ctx)
	}))
	ctx := context.New()
	ctx.Request.Method = "POST"
	ctx.Request.URL.Host = "foo.com"
	mx.Run("request", ctx)
	st.Expect(t, ctx.Request.Header.Get("foo"), "")
}
