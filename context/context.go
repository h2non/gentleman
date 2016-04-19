// Package context implements a simple request-aware HTTP context used by plugins
// and exposed by the middleware layer, designed to share polymorfic data
// types across plugins in the middleware call chain.
//
// Context is not thread-safe by default.
// In case that you support multithread programming in plugins you have to use locks/mutex accordingly.
package context

import (
	"io"
	"net/http"
	"net/url"

	"gopkg.in/h2non/gentleman.v1/utils"
)

// Context encapsulates required domain-specific HTTP entities
// to share data and entities for HTTP transactions in a middleware chain
type Context struct {
	// Stores the last error for the current Context
	Error error

	// Flag if the HTTP transaction was explicitly stopped in some phase
	Stopped bool

	// Context can inherit behavior and data from another Context
	Parent *Context

	// Reference to the http.Client used in the current HTTP transaction
	Client *http.Client

	// Reference to the http.Request used in the current HTTP transaction
	Request *http.Request

	// Reference to the http.Response used in the current HTTP transaction
	Response *http.Response
}

// New creates an empty default Context
func New() *Context {
	req := createRequest()
	res := createResponse(req)
	cli := &http.Client{Transport: http.DefaultTransport}
	return &Context{Request: req, Response: res, Client: cli}
}

// Set sets a value on the current store
func (c *Context) Set(key interface{}, value interface{}) {
	crc := getContextReadCloser(c.Request)
	crc.Context()[key] = value
}

// Get gets a value by key in the current or parent context
func (c *Context) Get(key interface{}) interface{} {
	crc := getContextReadCloser(c.Request)
	value := crc.Context()[key]
	if value != nil {
		return value
	}
	if c.Parent != nil {
		return c.Parent.Get(key)
	}
	return nil
}

// GetOk gets a context value from req.
// Returns (nil, false) if key not found in the request context.
func (c *Context) GetOk(key interface{}) (interface{}, bool) {
	crc := getContextReadCloser(c.Request)
	val, ok := crc.Context()[key]
	return val, ok
}

// GetString gets a string context value from req.
// Returns an empty string if key not found in the request context,
// or the value does not evaluate to a string
func (c *Context) GetString(key interface{}) string {
	crc := getContextReadCloser(c.Request)
	if value, ok := crc.Context()[key]; ok {
		if typed, ok := value.(string); ok {
			return typed
		}
	}
	return ""
}

// GetAll returns all stored context values for a request.
// Will always return a valid map. Returns an empty map for
// requests context data previously set
func (c *Context) GetAll() map[interface{}]interface{} {
	crc := getContextReadCloser(c.Request)
	return crc.Context()
}

// Delete deletes a stored value from a request’s context
func (c *Context) Delete(key interface{}) {
	crc := getContextReadCloser(c.Request)
	delete(crc.Context(), key)
}

// Clear clears all stored values from a request’s context
func (c *Context) Clear() {
	crc := getContextReadCloser(c.Request).(*contextReadCloser)
	crc.store = map[interface{}]interface{}{}
}

// UseParent uses a new parent Context
func (c *Context) UseParent(ctx *Context) {
	c.Parent = ctx
}

// Root returns the root Context looking in the parent contexts recursively.
// If the current context has no parent context, it will return the Context itself.
func (c *Context) Root() *Context {
	if c.Parent != nil {
		return c.Parent.Root()
	}
	return c
}

// SetRequest replaces the context http.Request
func (c *Context) SetRequest(req *http.Request) {
	c.Copy(req)
	c.Request = req
}

// Clone returns a clone of the current context.
func (c *Context) Clone() *Context {
	ctx := new(Context)
	*ctx = *c

	req := new(http.Request)
	*req = *c.Request
	ctx.Request = req
	c.Copy(ctx.Request)

	res := new(http.Response)
	*res = *c.Response
	ctx.Response = res

	return ctx
}

// Copy copies the current context data into a new http.Request
func (c *Context) Copy(req *http.Request) {
	req.Body = getContextReadCloser(c.Request).Clone()
}

// ReadCloser augments the io.ReadCloser interface
// with a Context() method
type ReadCloser interface {
	io.ReadCloser
	Clone() ReadCloser
	Context() map[interface{}]interface{}
}

type contextReadCloser struct {
	io.ReadCloser
	store map[interface{}]interface{}
}

func (crc *contextReadCloser) Context() map[interface{}]interface{} {
	return crc.store
}

func (crc *contextReadCloser) Clone() ReadCloser {
	clone := &contextReadCloser{
		ReadCloser: crc.ReadCloser,
		store:      make(map[interface{}]interface{}),
	}
	for key, value := range crc.store {
		clone.store[key] = value
	}
	return clone
}

func getContextReadCloser(req *http.Request) ReadCloser {
	crc, ok := req.Body.(ReadCloser)
	if !ok {
		crc = &contextReadCloser{
			ReadCloser: req.Body,
			store:      make(map[interface{}]interface{}),
		}
		req.Body = crc
	}
	return crc
}

func createRequest() *http.Request {
	req := &http.Request{
		Method:     "GET",
		URL:        &url.URL{},
		Host:       "",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Proto:      "HTTP/1.1",
		Header:     make(http.Header),
		Body:       utils.NopCloser(),
	}
	return req
}

func createResponse(req *http.Request) *http.Response {
	return &http.Response{
		ProtoMajor: 1,
		ProtoMinor: 1,
		Proto:      "HTTP/1.1",
		Request:    req,
		Header:     make(http.Header),
		Body:       utils.NopCloser(),
	}
}
