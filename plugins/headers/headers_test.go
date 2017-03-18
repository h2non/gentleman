package headers

import (
	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v2/context"
	"testing"
)

func TestHeaderSet(t *testing.T) {
	ctx := context.New()
	fn := newHandler()

	Set("foo", "bar").Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	st.Expect(t, ctx.Request.Header.Get("foo"), "bar")
}

func TestHeaderAdd(t *testing.T) {
	ctx := context.New()
	ctx.Request.Header.Set("foo", "foo")
	fn := newHandler()

	Add("foo", "bar").Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	st.Expect(t, ctx.Request.Header["Foo"][1], "bar")
}

func TestHeaderDel(t *testing.T) {
	ctx := context.New()
	ctx.Request.Header.Set("foo", "foo")
	fn := newHandler()

	Del("foo").Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	st.Expect(t, ctx.Request.Header.Get("foo"), "")
}

func TestHeaderSetMap(t *testing.T) {
	ctx := context.New()
	ctx.Request.Header.Set("foo", "foo")
	fn := newHandler()
	headers := map[string]string{"foo": "bar"}

	SetMap(headers).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
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
