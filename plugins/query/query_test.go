package query

import (
	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v0/context"
	"testing"
)

func TestQuerySet(t *testing.T) {
	ctx := context.New()
	ctx.Request.URL.RawQuery = "baz=foo&foo=foo"
	fn := newHandler()

	Set("foo", "bar").Request(ctx, fn.fn)
	st.Expect(t, fn.called, true)
	st.Expect(t, ctx.Request.URL.RawQuery, "baz=foo&foo=bar")
}

func TestQueryAdd(t *testing.T) {
	ctx := context.New()
	ctx.Request.URL.RawQuery = "foo=baz"
	fn := newHandler()

	Add("foo", "bar").Request(ctx, fn.fn)
	st.Expect(t, fn.called, true)
	st.Expect(t, ctx.Request.URL.RawQuery, "foo=baz&foo=bar")
}

func TestQueryDel(t *testing.T) {
	ctx := context.New()
	ctx.Request.URL.RawQuery = "foo=baz"
	fn := newHandler()

	Del("foo").Request(ctx, fn.fn)
	st.Expect(t, fn.called, true)
	st.Expect(t, ctx.Request.URL.RawQuery, "")
}

func TestQueryDelAll(t *testing.T) {
	ctx := context.New()
	ctx.Request.URL.RawQuery = "foo=baz&foo=foo"
	fn := newHandler()

	DelAll().Request(ctx, fn.fn)
	st.Expect(t, fn.called, true)
	st.Expect(t, ctx.Request.URL.RawQuery, "")
}

func TestQuerySetMap(t *testing.T) {
	ctx := context.New()
	ctx.Request.URL.RawQuery = "baz=foo&foo=foo"
	fn := newHandler()
	params := map[string]string{"foo": "bar"}

	SetMap(params).Request(ctx, fn.fn)
	st.Expect(t, fn.called, true)
	st.Expect(t, ctx.Request.URL.RawQuery, "baz=foo&foo=bar")
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
