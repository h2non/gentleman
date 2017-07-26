package transport

import (
	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v2/context"
	"net/http"
	"testing"
)

func TestSetTransport(t *testing.T) {
	ctx := context.New()
	fn := newHandler()
	transport := &http.Transport{}
	Set(transport).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	newTransport := ctx.Client.Transport.(*http.Transport)
	st.Expect(t, newTransport, transport)
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
