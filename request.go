package gentleman

import (
	"errors"
	"gopkg.in/h2non/gentleman.v0/context"
	"gopkg.in/h2non/gentleman.v0/middleware"
	"gopkg.in/h2non/gentleman.v0/mux"
	"gopkg.in/h2non/gentleman.v0/plugin"
	"gopkg.in/h2non/gentleman.v0/plugins/url"
	"gopkg.in/h2non/gentleman.v0/utils"
	"io"
	"net"
	"net/http"
	"time"
)

const (
	// UserAgent represents the static user agent name and version
	UserAgent = "gentleman/" + Version
)

var (
	// DialTimeout represents the maximum amount of time the network dialer can take
	DialTimeout = 30 * time.Second

	// DialKeepAlive represents the maximum amount of time too keep alive the socket
	DialKeepAlive = 30 * time.Second

	// TLSHandshakeTimeout represents the maximum amount of time that
	// TLS handshake can take defined in the default http.Transport
	TLSHandshakeTimeout = 10 * time.Second

	// RequestTimeout represents the maximum about of time that
	// a request can take, including dial / request / redirect processes
	RequestTimeout = 60 * time.Second

	// DefaultDialer defines the default network dialer
	DefaultDialer = &net.Dialer{
		Timeout:   DialTimeout,
		KeepAlive: DialKeepAlive,
	}

	// DefaultTransport stores the default HTTP transport to be used
	DefaultTransport = NewDefaultTransport(DefaultDialer)
)

// Request represents an basic HTTP structure entity
type Request struct {
	// Stores if the request was already dispatched
	dispatched bool

	// Optional reference to the gentleman.Client instance
	Client *Client

	// Request scope Context instance
	Context *context.Context

	// Request scope Middleware instance
	Middleware middleware.Middleware
}

// NewRequest creates a new Request entity
func NewRequest() *Request {
	ctx := context.New()
	ctx.Client.Transport = DefaultTransport
	ctx.Request.Header.Set("User-Agent", UserAgent)

	return &Request{
		Context:    ctx,
		Middleware: middleware.New(),
	}
}

// SetClient Attach a client to the current Request
// This is mostly done internally.
func (r *Request) SetClient(cli *Client) *Request {
	r.Client = cli
	r.Context.UseParent(cli.Context)
	r.Middleware.UseParent(cli.Middleware)
	return r
}

// Use attaches a new plugin in the middleware stack
func (r *Request) Use(p plugin.Plugin) *Request {
	r.Middleware.Use(p)
	return r
}

// UseRequest attaches a request middleware handler
func (r *Request) UseRequest(fn context.HandlerFunc) *Request {
	r.Middleware.UseRequest(fn)
	return r
}

// UseResponse attaches a response middleware handler
func (r *Request) UseResponse(fn context.HandlerFunc) *Request {
	r.Middleware.UseResponse(fn)
	return r
}

// UseError attaches an error middleware handler
func (r *Request) UseError(fn context.HandlerFunc) *Request {
	r.Middleware.UseError(fn)
	return r
}

// Mux is a middleware multiplexer for easy plugin composition
func (r *Request) Mux() *mux.Mux {
	mx := mux.New()
	r.Use(mx)
	return mx
}

// Method defines the HTTP verb to be used
func (r *Request) Method(method string) *Request {
	r.Context.Request.Method = method
	return r
}

// URL parses and defines the URL to be used in the HTTP request.
func (r *Request) URL(uri string) *Request {
	r.Use(url.URL(uri))
	return r
}

// Path defines the request URL path to be used in the HTTP request.
func (r *Request) Path(uri string) *Request {
	r.Use(url.URL(uri))
	return r
}

// Set sets a new header field by name and value.
func (r *Request) Set(name, value string) *Request {
	r.Context.Request.Header.Set(name, value)
	return r
}

// Body defines the HTTP request body data based on a io.Reader stream.
func (r *Request) Body(body io.Reader) *Request {
	return r
}

// Send is an alias to Do(), which executes the current request.
func (r *Request) Send() (*Response, error) {
	return r.Do()
}

// Do performs the HTTP request and returns the HTTP response
func (r *Request) Do() (*Response, error) {
	if r.dispatched {
		return nil, errors.New("gentleman: Request was already dispatched")
	}

	r.dispatched = true
	ctx := NewDispatcher(r).Dispatch()

	return buildResponse(ctx)
}

// Clone creates a new side-effects free Request based on the current one
func (r *Request) Clone() *Request {
	req := NewRequest()
	req.Client = r.Client
	req.Context = r.Context.Clone()
	req.Middleware = r.Middleware.Clone()
	return req
}

// NewDefaultTransport returns a new http.Transport with default values
// based on the given net.Dialer.
func NewDefaultTransport(dialer *net.Dialer) *http.Transport {
	transport := &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		Dial:                dialer.Dial,
		TLSHandshakeTimeout: TLSHandshakeTimeout,
	}
	utils.SetTransportFinalizer(transport)
	return transport
}
