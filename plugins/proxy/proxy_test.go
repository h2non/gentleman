package proxy

import (
	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v2/context"
	"net/http"
	"testing"
)

func TestProxy(t *testing.T) {
	ctx := context.New()
	ctx.Request.URL.Scheme = "http"

	fn := newHandler()
	servers := map[string]string{"http": "http://localhost:3128"}

	Set(servers).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)

	transport := ctx.Client.Transport.(*http.Transport)
	url, err := transport.Proxy(ctx.Request)

	st.Expect(t, err, nil)
	st.Expect(t, url.Host, "localhost:3128")
	st.Expect(t, url.Scheme, "http")
}

func TestProxyParseError(t *testing.T) {
	ctx := context.New()
	ctx.Request.URL.Scheme = "http"

	fn := newHandler()
	servers := map[string]string{"http": "://"}

	Set(servers).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)

	transport := ctx.Client.Transport.(*http.Transport)
	_, err := transport.Proxy(ctx.Request)

	st.Expect(t, err.Error(), "parse ://: missing protocol scheme")
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
