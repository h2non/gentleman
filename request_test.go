package gentleman

import (
	"errors"
	"fmt"
	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v0/context"
	"gopkg.in/h2non/gentleman.v0/utils"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world")
	}))
	defer ts.Close()

	req := NewRequest()
	req.UseRequest(func(ctx *context.Context, h context.Handler) {
		h.Next(ctx)
	})

	req.UseRequest(func(ctx *context.Context, h context.Handler) {
		u, _ := url.Parse(ts.URL)
		ctx.Request.URL = u
		h.Next(ctx)
	})

	res, err := req.Send()
	if err != nil {
		t.Errorf("Request error: %s", err)
	}
	if res.RawRequest.URL == nil {
		t.Error("Invalid context")
	}
	if res.StatusCode != 200 {
		t.Errorf("Invalid response status: %d", res.StatusCode)
	}
}

func TestMiddlewareErrorInjectionAndInterception(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world")
	}))
	defer ts.Close()

	req := NewRequest()
	req.UseRequest(func(ctx *context.Context, h context.Handler) {
		h.Next(ctx)
	})

	req.UseRequest(func(ctx *context.Context, h context.Handler) {
		u, _ := url.Parse(ts.URL)
		ctx.Request.URL = u
		h.Error(ctx, errors.New("Error"))
	})

	req.UseError(func(ctx *context.Context, h context.Handler) {
		ctx.Error = nil
		h.Next(ctx)
	})

	res, err := req.Send()
	if err != nil {
		t.Errorf("Request error: %s", err)
	}
	if res.RawRequest.URL == nil {
		t.Error("Invalid context")
	}
	if res.StatusCode != 200 {
		t.Errorf("Invalid response status: %d", res.StatusCode)
	}
}

func TestRequestResponseMiddleware(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world")
	}))
	defer ts.Close()

	req := NewRequest()
	req.URL(ts.URL)
	req.UseRequest(func(c *context.Context, h context.Handler) {
		c.Request.Header.Set("Client", "go")
		h.Next(c)
	})
	req.UseResponse(func(c *context.Context, h context.Handler) {
		c.Response.Header.Set("Server", c.Request.Header.Get("Client"))
		h.Next(c)
	})

	res, err := req.Do()
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode, 200)
	st.Expect(t, res.Header.Get("Server"), "go")
}

func TestRequestMux(t *testing.T) {
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

	req.Mux().AddMatcher(func(ctx *context.Context) bool {
		return ctx.GetString("$phase") == "request" && ctx.Request.Method == "GET"
	}).UseRequest(func(ctx *context.Context, h context.Handler) {
		ctx.Request.Header.Set("mux", "true")
		h.Next(ctx)
	})

	res, err := req.Send()
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode, 200)
	st.Expect(t, res.RawRequest.Header.Get("mux"), "true")
}

func TestRequestInterceptor(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("Server should not be reached!")
		fmt.Fprintln(w, "Hello, world")
	}))
	defer ts.Close()

	req := NewRequest()
	req.UseRequest(func(ctx *context.Context, h context.Handler) {
		ctx.Request.Header.Set("Client", "gentleman")

		ctx.Response.StatusCode = 201
		ctx.Response.Status = "201 Created"

		ctx.Response.Header.Set("Server", "gentleman")
		utils.WriteBodyString(ctx.Response, "Hello, gentleman")

		h.Stop(ctx)
	})
	req.UseRequest(func(ctx *context.Context, h context.Handler) {
		t.Fatal("middleware should not be called")
		h.Next(ctx)
	})

	res, err := req.Do()
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode, 201)
	st.Expect(t, res.RawRequest.Header.Get("Client"), "gentleman")
	st.Expect(t, res.RawResponse.Header.Get("Server"), "gentleman")
	st.Expect(t, res.String(), "Hello, gentleman")
}

func TestRequestMethod(t *testing.T) {
	req := NewRequest()
	req.Method("POST")
	req.Middleware.Run("request", req.Context)
	st.Expect(t, req.Context.Request.Method, "POST")
}

/*
func TestCancelRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world")
	}))
	defer ts.Close()

	client := New()

	client.UseRequest(func(r *http.Request, c *http.Client, s *http.Response, handler middleware.Handler) {
		time.Sleep(500 * time.Millisecond)
		if tra, ok := c.Transport.(*http.Transport); !ok {
			t.Fatal("transport does not implement CancelRequest(*http.Request)")
		} else {
			tra.CancelRequest(r)
		}
		handler.Stop(r, c, s)
	})

	req := client.Get(ts.URL)
	res, err := req.Do()
	if err != nil {
		t.Fatal(err)
	}

	//debug("> %#v => %#v", res, err)
	if res.StatusCode != 0 {
		t.Fatal("Invalid status code")
	}
}

func TestTimeoutRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		fmt.Fprintln(w, "Hello, world")
	}))
	defer ts.Close()

	client := New()

	client.UseRequest(func(r *http.Request, c *http.Client, s *http.Response, handler middleware.Handler) {
		c.Timeout = 50 * time.Millisecond
		handler.Next(r, c, s)
	})

	req := client.Get(ts.URL)
	res, err := req.Do()
	if err == nil {
		t.Fatal(err)
	}

	if res != nil {
		t.Fatal("Invalid response")
	}
}

func TestOverwriteTargetURL(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world")
	}))
	defer ts.Close()

	client := New()

	client.UseRequest(func(r *http.Request, c *http.Client, s *http.Response, handler middleware.Handler) {
		r.URL, _ = url.Parse(ts.URL)
		handler.Next(r, c, s)
	})

	req := client.Get("http://foo")
	res, err := req.Do()
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 200 {
		t.Fatal("Invalid response status")
	}
}

func TestRetryOnFailure(t *testing.T) {
	failed := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !failed {
			failed = true
			w.WriteHeader(500)
			fmt.Fprintln(w, "Error")
		} else {
			fmt.Fprintln(w, "Hello, world")
		}
	}))
	defer ts.Close()

	client := New()

	client.UseResponse(func(r *http.Request, c *http.Client, s *http.Response, handler middleware.Handler) {
		if s.StatusCode == 200 {
			handler.Next(r, c, s)
			return
		}

		res, err := c.Do(r)
		if err != nil {
			handler.Error(err)
			return
		}

		handler.Next(r, c, res)
	})

	req := client.Get(ts.URL)
	res, err := req.Do()
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 200 {
		t.Fatal("Invalid response status")
	}
}
*/
