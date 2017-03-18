package gentleman

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v2/context"
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
	st.Expect(t, ctx.Error, nil)
	st.Expect(t, ctx.Response.StatusCode, 200)
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
	st.Reject(t, err, nil)
	st.Reject(t, ctx.Error, nil)
}

func TestDispatcherInterceptor(t *testing.T) {
	req := NewRequest()
	req.UseRequest(func(ctx *context.Context, h context.Handler) {
		ctx.Response.StatusCode = 200
		h.Next(ctx)
	})
	req.UseResponse(func(ctx *context.Context, h context.Handler) {
		ctx.Response.StatusCode = 204
		h.Next(ctx)
	})

	ctx := NewDispatcher(req).Dispatch()
	st.Expect(t, ctx.Error, nil)
	st.Expect(t, ctx.Response.StatusCode, 204)
}

func TestDispatcherResponseError(t *testing.T) {
	req := NewRequest().URL("http://127.0.0.1:9123")
	req.UseError(func(ctx *context.Context, h context.Handler) {
		ctx.Response.StatusCode = 503
		h.Next(ctx)
	})

	ctx := NewDispatcher(req).Dispatch()
	st.Reject(t, ctx.Error, nil)
	st.Expect(t, ctx.Response.StatusCode, 503)
}

func TestDispatcherStopped(t *testing.T) {
	req := NewRequest().URL("http://127.0.0.1:9123")
	req.UseRequest(func(ctx *context.Context, h context.Handler) {
		ctx.Response.StatusCode = 503
		h.Stop(ctx)
	})

	ctx := NewDispatcher(req).Dispatch()
	st.Expect(t, ctx.Stopped, true)
	st.Expect(t, ctx.Error, nil)
	st.Expect(t, ctx.Response.StatusCode, 503)
}

func TestDispatcherStoppedMiddleware(t *testing.T) {
	req := NewRequest().URL("http://127.0.0.1:9123")

	req.UseRequest(func(ctx *context.Context, h context.Handler) {
		ctx.Set("foo", "bar")
		h.Stop(ctx)
	})
	req.UseHandler("stop", func(ctx *context.Context, h context.Handler) {
		h.Error(ctx, errors.New("stop"))
	})

	ctx := NewDispatcher(req).Dispatch()
	st.Expect(t, ctx.Stopped, true)
	st.Reject(t, ctx.Error, nil)
	st.Expect(t, ctx.Error.Error(), "stop")
	st.Expect(t, ctx.GetString("foo"), "bar")
}
