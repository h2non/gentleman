package timeout

import (
	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v2/context"
	"net/http"
	"testing"
)

func TestTimeout(t *testing.T) {
	ctx := context.New()
	fn := newHandler()
	Request(1000).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	st.Expect(t, ctx.Error, nil)
	st.Expect(t, int(ctx.Client.Timeout), 1000)
}

func TestTimeoutTLS(t *testing.T) {
	ctx := context.New()
	fn := newHandler()
	TLS(1000).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	st.Expect(t, ctx.Error, nil)

	transport := ctx.Client.Transport.(*http.Transport)
	st.Expect(t, int(transport.TLSHandshakeTimeout), 1000)
}

func TestTimeoutAll(t *testing.T) {
	ctx := context.New()
	fn := newHandler()
	All(Timeouts{Request: 1000, Dial: 1000, TLS: 1000}).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	st.Expect(t, ctx.Error, nil)
	transport := ctx.Client.Transport.(*http.Transport)
	st.Expect(t, int(ctx.Client.Timeout), 1000)
	st.Expect(t, int(transport.TLSHandshakeTimeout), 1000)
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
