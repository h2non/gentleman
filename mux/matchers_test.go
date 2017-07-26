package mux

import (
	"errors"
	"net/url"
	"testing"

	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v2/context"
)

func TestMatchMethod(t *testing.T) {
	mx := New()
	mx.Use(Method("GET").UseRequest(pass))
	ctx := context.New()
	ctx.Request.Method = "GET"
	mx.Run("request", ctx)
	st.Expect(t, ctx.GetString("foo"), "bar")
}

func TestMatchPath(t *testing.T) {
	cases := []struct {
		value   string
		path    string
		matches bool
	}{
		{"baz", "/bar/foo/baz", true},
		{"bar", "/bar/foo/baz", true},
		{"^/bar", "/bar/foo/baz", true},
		{"/foo/", "/bar/foo/baz", true},
		{"f*", "/bar/foo/baz", true},
		{"fo[o]", "/bar/foo/baz", true},
		{"baz$", "/bar/foo/baz", true},
		{"foo$", "/bar/foo/baz", false},
		{"foobar", "/bar/foo/baz", false},
	}

	for _, test := range cases {
		mx := New()
		mx.Use(Path(test.value).UseRequest(pass))
		ctx := context.New()
		ctx.Request.URL.Path = test.path
		mx.Run("request", ctx)
		match(t, ctx, test.matches)
	}
}

func TestMatchURL(t *testing.T) {
	cases := []struct {
		value   string
		url     string
		matches bool
	}{
		{"foo.com", "http://foo.com", true},
		{"foo", "http://foo.com", true},
		{".com", "http://foo.com", true},
		{"^http://foo", "http://foo.com", true},
		{"foo.com$", "http://foo.com", true},
	}

	for _, test := range cases {
		mx := New()
		mx.Use(URL(test.value).UseRequest(pass))
		ctx := context.New()
		u, _ := url.Parse(test.url)
		ctx.Request.URL = u
		mx.Run("request", ctx)
		match(t, ctx, test.matches)
	}
}

func TestMatchHost(t *testing.T) {
	cases := []struct {
		value   string
		url     string
		matches bool
	}{
		{"foo.com", "http://foo.com", true},
		{"foo", "http://foo.com", true},
		{".com", "http://foo.com", true},
		{"foo.com$", "http://foo.com", true},
		{"bar", "http://foo.com", false},
		{"^http://foo", "http://foo.com", false},
		{"^foo.com$", "http://foo.com", true},
	}

	for _, test := range cases {
		mx := New()
		mx.Use(Host(test.value).UseRequest(pass))
		ctx := context.New()
		u, _ := url.Parse(test.url)
		ctx.Request.URL = u
		mx.Run("request", ctx)
		match(t, ctx, test.matches)
	}
}

func TestMatchQuery(t *testing.T) {
	cases := []struct {
		key     string
		value   string
		url     string
		matches bool
	}{
		{"foo", "bar", "http://baz.com?foo=bar", true},
		{"foo", "^bar$", "http://foo.com?foo=bar&baz=foo", true},
		{"foo", "b[a]r", "http://foo.com?foo=bar&baz=foo", true},
		{"foo", "foo", "http://foo.com?foo=bar&baz=foo", false},
		{"foo", "baz", "http://foo.com?foo=bar&baz=foo", false},
		{"baz", "foo", "http://foo.com?foo=bar&baz=foo", true},
		{"foo", "foo", "http://foo.com", false},
	}

	for _, test := range cases {
		mx := New()
		mx.Use(Query(test.key, test.value).UseRequest(pass))
		ctx := context.New()
		u, _ := url.Parse(test.url)
		ctx.Request.URL = u
		mx.Run("request", ctx)
		match(t, ctx, test.matches)
	}
}

func TestMatchRequestHeader(t *testing.T) {
	cases := []struct {
		key     string
		value   string
		headers map[string]string
		matches bool
	}{
		{"foo", "bar", map[string]string{"foo": "bar"}, true},
		{"foo", "bar", map[string]string{"foo": "foobar"}, true},
		{"foo", "bar", map[string]string{"foo": "foo"}, false},
		{"foo", "bar", map[string]string{}, false},
	}

	for _, test := range cases {
		mx := New()
		mx.Use(RequestHeader(test.key, test.value).UseRequest(pass))
		ctx := context.New()
		for key, value := range test.headers {
			ctx.Request.Header.Set(key, value)
		}
		mx.Run("request", ctx)
		match(t, ctx, test.matches)
	}
}

func TestMatchResponseHeader(t *testing.T) {
	cases := []struct {
		key     string
		value   string
		headers map[string]string
		matches bool
	}{
		{"foo", "bar", map[string]string{"foo": "bar"}, true},
		{"foo", "bar", map[string]string{"foo": "foobar"}, true},
		{"foo", "bar", map[string]string{"foo": "foo"}, false},
		{"foo", "bar", map[string]string{}, false},
	}

	for _, test := range cases {
		mx := New()
		mx.Use(ResponseHeader(test.key, test.value).UseResponse(pass))
		ctx := context.New()
		for key, value := range test.headers {
			ctx.Response.Header.Set(key, value)
		}
		mx.Run("response", ctx)
		match(t, ctx, test.matches)
	}
}

func TestMatchType(t *testing.T) {
	cases := []struct {
		value   string
		content string
		matches bool
	}{
		{"html", "text/html", true},
		{"plain", "text/plain", true},
		{"json", "application/json", true},
		{"xml", "text/xml", false},
		{"xml", "application/xml", true},
		{"text", "text/xml", false},
		{"urlencoded", "application/x-www-form-urlencoded", true},
		{"form-data", "application/x-www-form-urlencoded", true},
	}

	for _, test := range cases {
		mx := New()
		mx.Use(Type(test.value).UseResponse(pass))
		ctx := context.New()
		ctx.Response.Header.Set("Content-Type", test.content)
		mx.Run("response", ctx)
		match(t, ctx, test.matches)
	}
}

func TestMatchStatus(t *testing.T) {
	cases := []struct {
		code     int
		response int
		matches  bool
	}{
		{200, 200, true},
		{204, 204, true},
		{204, 301, false},
		{200, 400, false},
		{200, 500, false},
	}

	for _, test := range cases {
		mx := New()
		mx.Use(Status(test.code).UseResponse(pass))
		ctx := context.New()
		ctx.Response.StatusCode = test.response
		mx.Run("response", ctx)
		match(t, ctx, test.matches)
	}
}

func TestMatchStatusRange(t *testing.T) {
	cases := []struct {
		start    int
		end      int
		response int
		matches  bool
	}{
		{200, 200, 200, true},
		{200, 204, 200, true},
		{200, 204, 204, true},
		{200, 300, 300, true},
		{200, 300, 301, false},
		{200, 300, 500, false},
		{200, 204, 400, false},
		{200, 300, 500, false},
		{200, 500, 500, true},
	}

	for _, test := range cases {
		mx := New()
		mx.Use(StatusRange(test.start, test.end).UseResponse(pass))
		ctx := context.New()
		ctx.Response.StatusCode = test.response
		mx.Run("response", ctx)
		match(t, ctx, test.matches)
	}
}

func TestMatchError(t *testing.T) {
	cases := []struct {
		err     error
		matches bool
	}{
		{errors.New("foo error"), true},
		{nil, false},
	}

	for _, test := range cases {
		mx := New()
		mx.Use(Error().UseError(pass))
		ctx := context.New()
		ctx.Error = test.err
		mx.Run("error", ctx)
		match(t, ctx, test.matches)
	}
}

func TestMatchServerError(t *testing.T) {
	cases := []struct {
		status  int
		matches bool
	}{
		{200, false},
		{400, false},
		{500, true},
		{503, true},
	}

	for _, test := range cases {
		mx := New()
		mx.Use(ServerError().UseResponse(pass))
		ctx := context.New()
		ctx.Response.StatusCode = test.status
		mx.Run("response", ctx)
		match(t, ctx, test.matches)
	}
}

func pass(ctx *context.Context, h context.Handler) {
	ctx.Set("foo", "bar")
	h.Next(ctx)
}

func match(t *testing.T, ctx *context.Context, shouldMatch bool) {
	if shouldMatch {
		st.Expect(t, ctx.GetString("foo"), "bar")
	} else {
		st.Expect(t, ctx.GetString("foo"), "")
	}
}
