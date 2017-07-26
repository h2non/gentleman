package gentleman

import (
	"net/http"

	"gopkg.in/h2non/gentleman.v2/context"
	"gopkg.in/h2non/gentleman.v2/middleware"
	"gopkg.in/h2non/gentleman.v2/plugin"
	"gopkg.in/h2non/gentleman.v2/plugins/cookies"
	"gopkg.in/h2non/gentleman.v2/plugins/headers"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

// NewContext is a convenient alias to context.New factory.
var NewContext = context.New

// NewHandler is a convenient alias to context.NewHandler factory.
var NewHandler = context.NewHandler

// NewMiddleware is a convenient alias to middleware.New factory.
var NewMiddleware = middleware.New

// Client represents a high-level HTTP client entity capable
// with a built-in middleware and context.
type Client struct {
	// Client entity can inherit behavior from a parent Client.
	Parent *Client

	// Each Client entity has it's own Context that will be inherited by requests or child clients.
	Context *context.Context

	// Client entity has its own Middleware layer to compose and inherit behavior.
	Middleware middleware.Middleware
}

// New creates a new high level client entity
// able to perform HTTP requests.
func New() *Client {
	return &Client{
		Context:    context.New(),
		Middleware: middleware.New(),
	}
}

// Request creates a new Request based on the current Client
func (c *Client) Request() *Request {
	req := NewRequest()
	req.SetClient(c)
	return req
}

// Get creates a new GET request.
func (c *Client) Get() *Request {
	req := c.Request()
	req.Method("GET")
	return req
}

// Post creates a new POST request.
func (c *Client) Post() *Request {
	req := c.Request()
	req.Method("POST")
	return req
}

// Put creates a new PUT request.
func (c *Client) Put() *Request {
	req := c.Request()
	req.Method("PUT")
	return req
}

// Delete creates a new DELETE request.
func (c *Client) Delete() *Request {
	req := c.Request()
	req.Method("DELETE")
	return req
}

// Patch creates a new PATCH request.
func (c *Client) Patch() *Request {
	req := c.Request()
	req.Method("PATCH")
	return req
}

// Head creates a new HEAD request.
func (c *Client) Head() *Request {
	req := c.Request()
	req.Method("HEAD")
	return req
}

// Method defines a the default HTTP method used by outgoing client requests.
func (c *Client) Method(name string) *Client {
	c.Middleware.UseRequest(func(ctx *context.Context, h context.Handler) {
		ctx.Request.Method = name
		h.Next(ctx)
	})
	return c
}

// URL defines the URL for client requests.
// Useful to define at client level the base URL and base path used by child requests.
func (c *Client) URL(uri string) *Client {
	c.Use(url.URL(uri))
	return c
}

// BaseURL defines the URL schema and host for client requests.
// Useful to define at client level the base URL used by client child requests.
func (c *Client) BaseURL(uri string) *Client {
	c.Use(url.BaseURL(uri))
	return c
}

// Path defines the URL base path for client requests.
func (c *Client) Path(path string) *Client {
	c.Use(url.Path(path))
	return c
}

// Param replaces a path param based on the given param name and value.
func (c *Client) Param(name, value string) *Client {
	c.Use(url.Param(name, value))
	return c
}

// Params replaces path params based on the given params key-value map.
func (c *Client) Params(params map[string]string) *Client {
	c.Use(url.Params(params))
	return c
}

// SetHeader sets a new header field by name and value.
// If another header exists with the same key, it will be overwritten.
func (c *Client) SetHeader(key, value string) *Client {
	c.Use(headers.Set(key, value))
	return c
}

// AddHeader adds a new header field by name and value
// without overwriting any existent header.
func (c *Client) AddHeader(name, value string) *Client {
	c.Use(headers.Add(name, value))
	return c
}

// SetHeaders adds new header fields based on the given map.
func (c *Client) SetHeaders(fields map[string]string) *Client {
	c.Use(headers.SetMap(fields))
	return c
}

// AddCookie sets a new cookie field bsaed on the given http.Cookie struct
// without overwriting any existent cookie.
func (c *Client) AddCookie(cookie *http.Cookie) *Client {
	c.Use(cookies.Add(cookie))
	return c
}

// AddCookies sets a new cookie field based on a list of http.Cookie
// without overwriting any existent cookie.
func (c *Client) AddCookies(data []*http.Cookie) *Client {
	c.Use(cookies.AddMultiple(data))
	return c
}

// CookieJar creates a cookie jar to store HTTP cookies when they are sent down.
func (c *Client) CookieJar() *Client {
	c.Use(cookies.Jar())
	return c
}

// Use uses a new plugin to the middleware stack.
func (c *Client) Use(p plugin.Plugin) *Client {
	c.Middleware.Use(p)
	return c
}

// UseRequest uses a new middleware function for request phase.
func (c *Client) UseRequest(fn context.HandlerFunc) *Client {
	c.Middleware.UseRequest(fn)
	return c
}

// UseResponse uses a new middleware function for response phase.
func (c *Client) UseResponse(fn context.HandlerFunc) *Client {
	c.Middleware.UseResponse(fn)
	return c
}

// UseError uses a new middleware function for error phase.
func (c *Client) UseError(fn context.HandlerFunc) *Client {
	c.Middleware.UseError(fn)
	return c
}

// UseHandler uses a new middleware function for the given phase.
func (c *Client) UseHandler(phase string, fn context.HandlerFunc) *Client {
	c.Middleware.UseHandler(phase, fn)
	return c
}

// UseParent uses another Client as parent
// inheriting its middleware stack and configuration.
func (c *Client) UseParent(parent *Client) *Client {
	c.Parent = parent
	c.Context.UseParent(parent.Context)
	c.Middleware.UseParent(parent.Middleware)
	return c
}
