package redirect

import (
	"net/http"
	"testing"

	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v2/context"
)

func TestRedirectPolicy(t *testing.T) {
	headers := make(http.Header)
	headers.Set("foo", "bar")

	req := &http.Request{Header: make(http.Header)}
	prevReq := &http.Request{Header: headers}
	pool := []*http.Request{prevReq}
	opts := Options{}

	err := redirectPolicy(opts, req, pool)
	st.Expect(t, err, nil)
	st.Expect(t, req.Header.Get("foo"), "bar")
}

func TestRedirectPolicyAllowsSensitiveHeadersWhenTrusted(t *testing.T) {
	headers := http.Header{}
	headers.Set("foo", "bar")
	headers.Set("Authorization", "bar")

	req := &http.Request{Header: make(http.Header)}
	prevReq := &http.Request{Header: headers}
	pool := []*http.Request{prevReq}
	opts := Options{SensitiveHeaders: []string{"Authorization"}, Trusted: true}

	err := redirectPolicy(opts, req, pool)
	st.Expect(t, err, nil)
	st.Expect(t, req.Header.Get("foo"), "bar")
	st.Expect(t, req.Header.Get("Authorization"), "bar")
}

func TestRedirectPolicyAllowsSensitiveHeadersForTrustedHosts(t *testing.T) {
	headers := http.Header{}
	headers.Set("foo", "bar")
	headers.Set("Authorization", "bar")

	req := &http.Request{Host: "www.trusted.com", Header: make(http.Header)}
	prevReq := &http.Request{Header: headers}
	pool := []*http.Request{prevReq}
	opts := Options{SensitiveHeaders: []string{"Authorization"}, TrustedHostSuffixes: []string{".trusted.com"}}

	err := redirectPolicy(opts, req, pool)
	st.Expect(t, err, nil)
	st.Expect(t, req.Header.Get("foo"), "bar")
	st.Expect(t, req.Header.Get("Authorization"), "bar")
}

func TestRedirectPolicyRemovesSensitiveHeadersForNonTrustedHosts(t *testing.T) {
	headers := http.Header{}
	headers.Set("foo", "bar")
	headers.Set("Authorization", "bar")

	req := &http.Request{Host: "untrusted.com", Header: make(http.Header)}
	prevReq := &http.Request{Header: headers}
	pool := []*http.Request{prevReq}
	opts := Options{SensitiveHeaders: []string{"Authorization"}, TrustedHostSuffixes: []string{".trusted.com"}}

	err := redirectPolicy(opts, req, pool)
	st.Expect(t, err, nil)
	st.Expect(t, req.Header.Get("foo"), "bar")
	st.Expect(t, req.Header.Get("Authorization"), "")
}

func TestRedirectPolicyRemoveSensitiveHeaders(t *testing.T) {
	headers := http.Header{}
	headers.Set("foo", "bar")
	headers.Set("Authorization", "bar")

	req := &http.Request{Header: make(http.Header)}
	prevReq := &http.Request{Header: headers}
	pool := []*http.Request{prevReq}
	opts := Options{SensitiveHeaders: []string{"Authorization"}}

	err := redirectPolicy(opts, req, pool)
	st.Expect(t, err, nil)
	st.Expect(t, req.Header.Get("foo"), "bar")
	st.Expect(t, req.Header.Get("Authorization"), "")
}

func TestRedirectPolicyLimit(t *testing.T) {
	req := &http.Request{Header: make(http.Header)}
	prevReq := &http.Request{Header: make(http.Header)}
	pool := []*http.Request{prevReq, prevReq}
	opts := Options{Limit: 1}

	err := redirectPolicy(opts, req, pool)
	st.Expect(t, err, ErrRedirectLimitExceeded)
}

func TestRedirectPlugin(t *testing.T) {
	ctx := context.New()
	fn := newHandler()
	Config(Options{}).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
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
