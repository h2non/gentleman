package gentleman

import (
	"gopkg.in/h2non/gentleman.v0/context"
	"gopkg.in/h2non/gentleman.v0/middleware"
	"gopkg.in/h2non/gentleman.v0/plugin"
)

// Client represents a high-level HTTP client entity.
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

// Use attaches a new plugin to the middleware stack.
func (c *Client) Use(p plugin.Plugin) {
	c.Middleware.Use(p)
}

// UseRequest attaches a new middleware function for request phase.
func (c *Client) UseRequest(fn context.HandlerFunc) {
	c.Middleware.UseRequest(fn)
}

// UseResponse attaches a new middleware function for response phase.
func (c *Client) UseResponse(fn context.HandlerFunc) {
	c.Middleware.UseResponse(fn)
}

// UseError attaches a new middleware function for error phase.
func (c *Client) UseError(fn context.HandlerFunc) {
	c.Middleware.UseError(fn)
}

// UseParent uses another Client as parent
// inheriting it's middleware stack and configuration.
func (c *Client) UseParent(parent *Client) {
	c.Parent = parent
	c.Context.UseParent(parent.Context)
	c.Middleware.UseParent(parent.Middleware)
}
