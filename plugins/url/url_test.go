package url

import (
	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v2/context"
	"net/url"
	"testing"
)

type test struct {
	value string
	url   *url.URL
}

func TestURL(t *testing.T) {
	cases := []test{
		{"http://foo", &url.URL{Host: "foo", Scheme: "http"}},
		{"http://foo:123", &url.URL{Host: "foo:123", Scheme: "http"}},
		{"https://127.0.0.1", &url.URL{Host: "127.0.0.1", Scheme: "https"}},
		{"https://foo/bar", &url.URL{Host: "foo", Path: "/bar", Scheme: "https"}},
		{"foo/bar", &url.URL{Host: "foo", Path: "/bar", Scheme: "http"}},
	}

	ctx := context.New()
	fn := newHandler()

	for _, test := range cases {
		URL(test.value).Exec("request", ctx, fn.fn)
		assert(t, fn, ctx, test)
	}
}

func TestBaseURL(t *testing.T) {
	cases := []test{
		{"http://foo", &url.URL{Host: "foo", Scheme: "http"}},
		{"http://foo:123", &url.URL{Host: "foo:123", Scheme: "http"}},
		{"https://127.0.0.1", &url.URL{Host: "127.0.0.1", Scheme: "https"}},
		{"https://foo/bar", &url.URL{Host: "foo", Scheme: "https"}},
		{"foo/bar", &url.URL{Host: "foo", Scheme: "http"}},
	}

	ctx := context.New()
	fn := newHandler()

	for _, test := range cases {
		BaseURL(test.value).Exec("request", ctx, fn.fn)
		assert(t, fn, ctx, test)
	}
}

func TestPath(t *testing.T) {
	cases := []test{
		{"/", &url.URL{Path: ""}},
		{"/foo", &url.URL{Path: "/foo"}},
		{"/foo/bar", &url.URL{Path: "/foo/bar"}},
	}

	ctx := context.New()
	fn := newHandler()

	ctx.Request.URL.Path = "/baz"
	for _, test := range cases {
		Path(test.value).Exec("request", ctx, fn.fn)
		assert(t, fn, ctx, test)
	}
}

func TestAddPath(t *testing.T) {
	cases := []test{
		{"/", &url.URL{Path: "/baz"}},
		{"/foo", &url.URL{Path: "/baz/foo"}},
		{"/foo/bar", &url.URL{Path: "/baz/foo/bar"}},
	}

	for _, test := range cases {
		ctx := context.New()
		fn := newHandler()
		ctx.Request.URL.Path = "/baz"

		AddPath(test.value).Exec("request", ctx, fn.fn)
		assert(t, fn, ctx, test)
	}
}

func TestPathPrefix(t *testing.T) {
	cases := []test{
		{"/", &url.URL{Path: "/baz"}},
		{"/foo", &url.URL{Path: "/foo/baz"}},
		{"/foo/bar", &url.URL{Path: "/foo/bar/baz"}},
	}

	for _, test := range cases {
		ctx := context.New()
		fn := newHandler()
		ctx.Request.URL.Path = "/baz"

		PathPrefix(test.value).Exec("request", ctx, fn.fn)
		assert(t, fn, ctx, test)
	}
}

func TestPathParam(t *testing.T) {
	cases := []struct {
		key   string
		value string
		path  string
		url   *url.URL
	}{
		{"", "bar", "/bar", &url.URL{Path: "/bar"}},
		{"foo", "bar", "/bar", &url.URL{Path: "/bar"}},
		{"foo", "bar", "/:foo", &url.URL{Path: "/bar"}},
		{"foo", "bar", "/foo/:foo", &url.URL{Path: "/foo/bar"}},
		{"foo", "bar", "/:foo/bar", &url.URL{Path: "/bar/bar"}},
		{"foo", "bar", "/:foo/:foo", &url.URL{Path: "/bar/bar"}},
		{"foo", "bar", "/:foo/bar/:foo", &url.URL{Path: "/bar/bar/bar"}},
	}

	for _, test := range cases {
		ctx := context.New()
		fn := newHandler()
		ctx.Request.URL.Path = test.path
		Param(test.key, test.value).Exec("request", ctx, fn.fn)
		st.Expect(t, ctx.Request.URL.Path, test.url.Path)
	}
}

func TestPathParams(t *testing.T) {
	cases := []struct {
		list map[string]string
		path string
		url  *url.URL
	}{
		{map[string]string{"": "bar"}, "/bar", &url.URL{Path: "/bar"}},
		{map[string]string{"bar": "bar"}, "/bar", &url.URL{Path: "/bar"}},
		{map[string]string{"foo": "bar"}, "/:foo", &url.URL{Path: "/bar"}},
		{map[string]string{"foo": "bar"}, "/foo/:foo", &url.URL{Path: "/foo/bar"}},
		{map[string]string{"foo": "bar"}, "/:foo/bar/:foo", &url.URL{Path: "/bar/bar/bar"}},
		{map[string]string{"foo": "bar", "baz": "foo"}, "/:foo/bar/:baz", &url.URL{Path: "/bar/bar/foo"}},
	}

	for _, test := range cases {
		ctx := context.New()
		fn := newHandler()
		ctx.Request.URL.Path = test.path
		Params(test.list).Exec("request", ctx, fn.fn)
		st.Expect(t, ctx.Request.URL.Path, test.url.Path)
	}
}

func assert(t *testing.T, fn *handler, ctx *context.Context, test test) {
	st.Expect(t, fn.called, true)
	st.Expect(t, ctx.Error, nil)
	st.Expect(t, ctx.Request.URL.Host, test.url.Host)
	st.Expect(t, ctx.Request.URL.Scheme, test.url.Scheme)
	st.Expect(t, ctx.Request.URL.Path, test.url.Path)
	st.Expect(t, ctx.Request.URL.RawQuery, test.url.RawQuery)
	st.Expect(t, ctx.Request.URL.Fragment, test.url.Fragment)
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
