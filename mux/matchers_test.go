package mux

import (
	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v0/context"
	"testing"
)

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
