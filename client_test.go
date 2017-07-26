package gentleman

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v2/context"
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
	st.Expect(t, ctx.GetString("foo"), "barbar")
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
	st.Expect(t, ctx.Request.Header.Get("Client"), "gogo")
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
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode, 200)
	st.Expect(t, res.Header.Get("Server"), "go")
	st.Expect(t, res.Header.Get("Agent"), "gentleman")
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
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode, 200)
	st.Expect(t, res.Header.Get("Server"), "go")
}

func TestClientErrorMiddleware(t *testing.T) {
	client := New()
	client.UseRequest(func(c *context.Context, h context.Handler) {
		c.Error = errors.New("foo error")
		h.Next(c)
	})
	client.UseError(func(c *context.Context, h context.Handler) {
		c.Error = errors.New("error: " + c.Error.Error())
		h.Next(c)
	})

	req := client.Request()
	res, err := req.Do()
	st.Expect(t, err.Error(), "error: foo error")
	st.Expect(t, res.Ok, false)
	st.Expect(t, res.StatusCode, 0)
}

func TestClientCustomPhaseMiddleware(t *testing.T) {
	client := New()
	client.UseRequest(func(c *context.Context, h context.Handler) {
		c.Error = errors.New("foo error")
		h.Next(c)
	})
	client.UseHandler("error", func(c *context.Context, h context.Handler) {
		c.Error = errors.New("error: " + c.Error.Error())
		h.Next(c)
	})

	req := client.Request()
	res, err := req.Do()
	st.Expect(t, err.Error(), "error: foo error")
	st.Expect(t, res.Ok, false)
	st.Expect(t, res.StatusCode, 0)
}

func TestClientMethod(t *testing.T) {
	cli := New()
	cli.Method("POST")
	cli.Middleware.Run("request", cli.Context)
	st.Expect(t, cli.Context.Request.Method, "POST")
}

func TestClientURL(t *testing.T) {
	url := "http://foo.com"
	cli := New()
	cli.URL(url)
	cli.Middleware.Run("request", cli.Context)
	st.Expect(t, cli.Context.Request.URL.String(), url)
}

func TestClientBaseURL(t *testing.T) {
	url := "http://foo.com/bar/baz"
	cli := New()
	cli.BaseURL(url)
	cli.Middleware.Run("request", cli.Context)
	st.Expect(t, cli.Context.Request.URL.String(), "http://foo.com")
}

func TestClientPath(t *testing.T) {
	url := "http://foo.com/bar/baz"
	path := "/foo/baz"
	cli := New()
	cli.URL(url)
	cli.Path(path)
	cli.Middleware.Run("request", cli.Context)
	st.Expect(t, cli.Context.Request.URL.String(), "http://foo.com/foo/baz")
}

func TestClientPathParam(t *testing.T) {
	url := "http://foo.com/bar/baz"
	path := "/:foo/bar/:baz"
	cli := New()
	cli.URL(url)
	cli.Path(path)
	cli.Param("foo", "baz")
	cli.Param("baz", "foo")
	cli.Middleware.Run("request", cli.Context)
	st.Expect(t, cli.Context.Request.URL.String(), "http://foo.com/baz/bar/foo")
}

func TestClientPathParams(t *testing.T) {
	url := "http://foo.com/bar/baz"
	path := "/:foo/bar/:baz"
	cli := New()
	cli.URL(url)
	cli.Path(path)
	cli.Params(map[string]string{"foo": "baz", "baz": "foo"})
	cli.Middleware.Run("request", cli.Context)
	st.Expect(t, cli.Context.Request.URL.String(), "http://foo.com/baz/bar/foo")
}

func TestClientSetHeader(t *testing.T) {
	cli := New()
	cli.SetHeader("foo", "bar")
	cli.Middleware.Run("request", cli.Context)
	st.Expect(t, cli.Context.Request.Header.Get("foo"), "bar")
}

func TestClientAddHeader(t *testing.T) {
	cli := New()
	cli.AddHeader("foo", "baz")
	cli.AddHeader("foo", "bar")
	cli.Middleware.Run("request", cli.Context)
	st.Expect(t, cli.Context.Request.Header.Get("foo"), "baz")
}

func TestClientSetHeaders(t *testing.T) {
	cli := New()
	cli.SetHeaders(map[string]string{"foo": "baz", "baz": "foo"})
	cli.Middleware.Run("request", cli.Context)
	st.Expect(t, cli.Context.Request.Header.Get("foo"), "baz")
	st.Expect(t, cli.Context.Request.Header.Get("baz"), "foo")
}

func TestClientAddCookie(t *testing.T) {
	cli := New()
	cookie := &http.Cookie{Name: "foo", Value: "bar"}
	cli.AddCookie(cookie)
	cli.Middleware.Run("request", cli.Context)
	st.Expect(t, cli.Context.Request.Header.Get("Cookie"), "foo=bar")
}

func TestClientAddCookies(t *testing.T) {
	cli := New()
	cookies := []*http.Cookie{{Name: "foo", Value: "bar"}}
	cli.AddCookies(cookies)
	cli.Middleware.Run("request", cli.Context)
	st.Expect(t, cli.Context.Request.Header.Get("Cookie"), "foo=bar")
}

func TestClientCookieJar(t *testing.T) {
	cli := New()
	cli.CookieJar()
	cli.Middleware.Run("request", cli.Context)
	st.Reject(t, cli.Context.Client.Jar, nil)
}

func TestClientVerbMethods(t *testing.T) {
	cli := New()
	req := cli.Get()
	req.Middleware.Run("request", req.Context)
	if req.Context.Request.Method != "GET" {
		t.Errorf("Invalid request method: %s", req.Context.Request.Method)
	}

	cli = New()
	req = cli.Post()
	req.Middleware.Run("request", req.Context)
	if req.Context.Request.Method != "POST" {
		t.Errorf("Invalid request method: %s", req.Context.Request.Method)
	}

	cli = New()
	req = cli.Put()
	req.Middleware.Run("request", req.Context)
	if req.Context.Request.Method != "PUT" {
		t.Errorf("Invalid request method: %s", req.Context.Request.Method)
	}

	cli = New()
	req = cli.Delete()
	req.Middleware.Run("request", req.Context)
	if req.Context.Request.Method != "DELETE" {
		t.Errorf("Invalid request method: %s", req.Context.Request.Method)
	}

	cli = New()
	req = cli.Patch()
	req.Middleware.Run("request", req.Context)
	if req.Context.Request.Method != "PATCH" {
		t.Errorf("Invalid request method: %s", req.Context.Request.Method)
	}

	cli = New()
	req = cli.Head()
	req.Middleware.Run("request", req.Context)
	if req.Context.Request.Method != "HEAD" {
		t.Errorf("Invalid request method: %s", req.Context.Request.Method)
	}
}
