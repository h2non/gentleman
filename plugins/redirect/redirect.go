package redirect

import (
	"errors"
	"net/http"
	"strings"

	c "gopkg.in/h2non/gentleman.v2/context"
	p "gopkg.in/h2non/gentleman.v2/plugin"
)

var (
	// ErrRedirectLimitExceeded is the error returned when the request responded
	// with too many redirects
	ErrRedirectLimitExceeded = errors.New("gentleman: Request exceeded redirect count")

	// RedirectLimit defines the maximum number of redirects to follow in a request
	RedirectLimit = 10

	// SensitiveHeaders is a map of sensitive HTTP headers that a user
	// doesn't want passed on a redirect. This is the global variable
	SensitiveHeaders = []string{
		"WWW-Authenticate",
		"Authorization",
		"Proxy-Authorization",
	}
)

// Options store the redirect policy options
type Options struct {
	// Limit is the acceptable amount of redirects that we should expect
	// before returning an error be default this is set to 30. You can change this
	// globally by modifying the `Limit` variable
	Limit int

	// Trusted is a flag that will enable all headers to be
	// forwarded to the redirect location. Otherwise, the headers specified in
	// `SensitiveHeaders` will be removed from the request
	Trusted bool

	// SensitiveHeaders is a map of sensitive HTTP headers that a user
	// doesn't want passed on a redirect
	SensitiveHeaders []string
}

// Config defines in the request http.Client the redirect
// policy based on the given options.
func Config(opts Options) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		ctx.Client.CheckRedirect = func(req *http.Request, pool []*http.Request) error {
			return redirectPolicy(opts, req, pool)
		}
		h.Next(ctx)
	})
}

// Limit defines in the maximum number of redirects that http.Client should follow.
func Limit(limit int) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		ctx.Client.CheckRedirect = func(req *http.Request, pool []*http.Request) error {
			return redirectPolicy(Options{Limit: limit}, req, pool)
		}
		h.Next(ctx)
	})
}

func redirectPolicy(opts Options, req *http.Request, pool []*http.Request) error {
	if opts.Limit == 0 {
		opts.Limit = RedirectLimit
	}

	if len(pool) >= opts.Limit {
		return ErrRedirectLimitExceeded
	}

	if opts.SensitiveHeaders == nil {
		opts.SensitiveHeaders = SensitiveHeaders
	}

	for k, vv := range pool[0].Header {
		copyHeaders(k, vv, opts, req)
	}

	return nil
}

func copyHeaders(k string, vv []string, opts Options, req *http.Request) {
	if !opts.Trusted {
		for _, v := range opts.SensitiveHeaders {
			if strings.ToLower(k) == strings.ToLower(v) {
				return
			}
		}
	}

	req.Header[k] = vv
}
