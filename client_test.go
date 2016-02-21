package gentleman

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v0/context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientMiddlewareContext(t *testing.T) {
	fn := func(ctx *context.Context, h context.Handler) {
		ctx.Set("foo", ctx.GetString("foo")+"bar")
		h.Next(ctx)
	}

	cli := New()
	cli.UseRequest(fn)
	cli.UseResponse(fn)

	if len(cli.Middleware.GetStack()) != 2 {
		t.Error("Invalid middleware stack length")
	}

	ctx := NewContext()
	cli.Middleware.Run("request", ctx)
	cli.Middleware.Run("response", ctx)

	if ctx.GetString("foo") != "barbar" {
		t.Error("Invalid context value")
	}
}

func TestClientInheritance(t *testing.T) {
	parent := New()
	cli := New()
	cli.UseParent(parent)

	parent.UseRequest(func(ctx *context.Context, h context.Handler) {
		ctx.Request.Header.Set("Client", "go")
		h.Next(ctx)
	})
	cli.UseRequest(func(ctx *context.Context, h context.Handler) {
		ctx.Request.Header.Set("Client", ctx.Request.Header.Get("Client")+"go")
		h.Next(ctx)
	})

	ctx := NewContext()
	cli.Middleware.Run("request", ctx)
	if header := ctx.Request.Header.Get("Client"); header != "gogo" {
		t.Errorf("Invalid client header: %s", header)
	}
}

func TestClientRequestMiddleware(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", r.Header.Get("Client"))
		w.Header().Set("Agent", r.Header.Get("Agent"))
		fmt.Fprintln(w, "Hello, world")
	}))
	defer ts.Close()

	client := New()

	client.UseRequest(func(ctx *context.Context, h context.Handler) {
		ctx.Request.Header.Set("Client", "go")
		h.Next(ctx)
	})

	req := client.Request()
	req.URL(ts.URL)
	req.UseRequest(func(ctx *context.Context, h context.Handler) {
		ctx.Request.Header.Set("Agent", "gentleman")
		h.Next(ctx)
	})

	res, err := req.Do()
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Invalid status code: %s", res.StatusCode)
	}
	if res.Header.Get("Server") != "go" {
		t.Error("Invalid server header")
	}
	if res.Header.Get("Agent") != "gentleman" {
		t.Error("Invalid agent header")
	}
}

func TestClientRequestResponseMiddleware(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world")
	}))
	defer ts.Close()

	client := New()

	client.UseRequest(func(c *context.Context, h context.Handler) {
		c.Request.Header.Set("Client", "go")
		h.Next(c)
	})

	client.UseResponse(func(c *context.Context, h context.Handler) {
		c.Response.Header.Set("Server", c.Request.Header.Get("Client"))
		h.Next(c)
	})

	req := client.Request()
	req.URL(ts.URL)
	res, err := req.Do()
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Error("Invalid status code")
	}
	if res.Header.Get("Server") != "go" {
		t.Error("Invalid response header")
	}
}
