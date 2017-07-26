package tls

import (
	"crypto/tls"
	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v2/context"
	"net/http"
	"testing"
)

func TestAuthBasic(t *testing.T) {
	ctx := context.New()
	fn := newHandler()
	config := &tls.Config{InsecureSkipVerify: true}
	Config(config).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)

	transport := ctx.Client.Transport.(*http.Transport)
	st.Expect(t, transport.TLSClientConfig, config)
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
