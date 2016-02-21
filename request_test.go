package gentleman

import (
	"errors"
	"fmt"
	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v0/context"
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

	res, err := req.End()
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

	res, err := req.End()
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

	res, err := req.End()
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode, 200)
	st.Expect(t, res.RawRequest.Header.Get("mux"), "true")
}

/*
func TestIncerceptRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("Server should not be reached!")
		fmt.Fprintln(w, "Hello, world")
	}))
	defer ts.Close()

	client := New()

	handler := func(r *http.Request, c *http.Client, s *http.Response, handler middleware.Handler) {
		r.Header.Set("Client", "pep")

		s.StatusCode = 201
		s.Status = "201 Created"

		s.Header.Set("Server", "pep")
		WriteBodyString(s, "Hello, pep")

		handler.Stop(r, c, s)
	}

	client.UseRequest(handler)

	req := client.NewRequest("GET", ts.URL, nil)
	res, err := req.Do()
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 201 {
		t.Fatal("Invalid status code")
	}

	if res.Header.Get("Server") != "pep" {
		t.Fatal("Invalid server header")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if string(body[:10]) != "Hello, pep" {
		t.Fatal("Invalid body")
	}
}

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

func TestInterceptRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world")
	}))
	defer ts.Close()

	client := New()

	client.UseRequest(func(r *http.Request, c *http.Client, s *http.Response, handler middleware.Handler) {
		s.StatusCode = 400
		s.Header.Set("foo", "bar")
		WriteBodyString(s, "foo, bar")
		handler.Stop(r, c, s)
	})

	req := client.Get(ts.URL)
	res, err := req.Do()
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 400 {
		t.Fatal("Invalid response status")
	}

	if res.Header.Get("foo") != "bar" {
		t.Fatal("Invalid header")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if string(body[:8]) != "foo, bar" {
		t.Fatal("Invalid body")
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
