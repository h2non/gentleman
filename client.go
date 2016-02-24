package gentleman

import (
	"gopkg.in/h2non/gentleman.v0/context"
	"gopkg.in/h2non/gentleman.v0/middleware"
	"gopkg.in/h2non/gentleman.v0/plugin"
	"gopkg.in/h2non/gentleman.v0/plugins/headers"
	"gopkg.in/h2non/gentleman.v0/plugins/url"
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
	c.Use(url.URL(uri))
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

// Set sets a new HTTP header field by key and value in the outgoing client requests.
func (c *Client) Set(key, value string) *Client {
	c.Use(headers.Set(key, value))
	return c
}

// Use attaches a new plugin to the middleware stack.
func (c *Client) Use(p plugin.Plugin) *Client {
	c.Middleware.Use(p)
	return c
}

// UseRequest attaches a new middleware function for request phase.
func (c *Client) UseRequest(fn context.HandlerFunc) *Client {
	c.Middleware.UseRequest(fn)
	return c
}

// UseResponse attaches a new middleware function for response phase.
func (c *Client) UseResponse(fn context.HandlerFunc) *Client {
	c.Middleware.UseResponse(fn)
	return c
}

// UseError attaches a new middleware function for error phase.
func (c *Client) UseError(fn context.HandlerFunc) *Client {
	c.Middleware.UseError(fn)
	return c
}

// UseParent uses another Client as parent
// inheriting it's middleware stack and configuration.
func (c *Client) UseParent(parent *Client) *Client {
	c.Parent = parent
	c.Context.UseParent(parent.Context)
	c.Middleware.UseParent(parent.Middleware)
	return c
}
