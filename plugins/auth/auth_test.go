package auth

import (
	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v2/context"
	"testing"
)

func TestAuthBasic(t *testing.T) {
	ctx := context.New()
	fn := newHandler()
	Basic("foo", "bar").Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	st.Expect(t, ctx.Request.Header.Get("Authorization"), "Basic Zm9vOmJhcg==")
}

func TestAuthBearer(t *testing.T) {
	ctx := context.New()
	fn := newHandler()
	Bearer("foo").Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	st.Expect(t, ctx.Request.Header.Get("Authorization"), "Bearer foo")
}

func TestAuthCustom(t *testing.T) {
	ctx := context.New()
	fn := newHandler()
	Custom("Token foo").Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	st.Expect(t, ctx.Request.Header.Get("Authorization"), "Token foo")
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
