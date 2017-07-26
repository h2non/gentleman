package bodytype

import (
	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v2/context"
	"net/http"
	"testing"
)

func TestBodyTypeDefine(t *testing.T) {
	for name, mime := range Types {
		req := &http.Request{Header: http.Header{}}
		defineType(name, req)
		if req.Header.Get("Content-Type") != mime {
			t.Errorf("Invalid MIME type: %s != %s", name, mime)
		}
	}
}

func TestBodyTypeDefineUnsupported(t *testing.T) {
	req := &http.Request{Header: http.Header{}}
	defineType("text/plain", req)
	if req.Header.Get("Content-Type") != "text/plain" {
		t.Error("Invalid MIME type")
	}
}

func TestBodyType(t *testing.T) {
	ctx := context.New()
	fn := newHandler()
	Set("json").Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	st.Expect(t, ctx.Request.Header.Get("Content-Type"), "application/json")
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
