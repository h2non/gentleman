package gentleman

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v0/context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestDispatcher(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world")
	}))
	defer ts.Close()

	req := NewRequest()
	req.UseRequest(func(ctx *context.Context, h context.Handler) {
		u, _ := url.Parse(ts.URL)
		ctx.Request.URL = u
		h.Next(ctx)
	})

	ctx := NewDispatcher(req).Dispatch()
	if ctx.Error != nil {
		t.Errorf("Dispatcher error: %s", ctx.Error)
	}
	if ctx.Response.StatusCode != 200 {
		t.Errorf("Invalid status code: %d", ctx.Response.StatusCode)
	}
}

func TestDispatcherError(t *testing.T) {
	req := NewRequest()
	req.UseRequest(func(ctx *context.Context, h context.Handler) {
		u, _ := url.Parse("http://127.0.0.1:9123")
		ctx.Request.URL = u
		h.Next(ctx)
	})

	var err error
	req.UseError(func(ctx *context.Context, h context.Handler) {
		err = ctx.Error
		h.Next(ctx)
	})

	ctx := NewDispatcher(req).Dispatch()
	if ctx.Error == nil || err == nil {
		t.Error("Error cannot be empty")
	}
}
