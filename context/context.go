// Package context implements a request-aware HTTP context used by plugins
// and exposed by the middleware layer, designed to share polymorfic data
// types across plugins in the middleware call chain.
//
// It is built on top of the standard built-in context package in Go:
// https://golang.org/pkg/context
//
// gentleman's Context implements the stdlib `context.Context` interface:
// https://golang.org/pkg/context/#Context
package context

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"gopkg.in/h2non/gentleman.v2/utils"
)

// Key stores the key identifier for the built-in context
var Key interface{} = "$gentleman"

// Store represents the map store for context store.
type Store map[interface{}]interface{}

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

// getStore retrieves the current request context data store.
func (c *Context) getStore() Store {
	store, ok := c.Request.Context().Value(Key).(Store)
	if !ok {
		panic("invalid request context")
	}
	return store
}

// Set sets a value on the current store
func (c *Context) Set(key interface{}, value interface{}) {
	store := c.getStore()
	store[key] = value
}

// Get gets a value by key in the current or parent context
func (c *Context) Get(key interface{}) interface{} {
	store := c.getStore()
	if store == nil {
		return store
	}
	if value, ok := store[key]; ok {
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
	store := c.getStore()
	val, ok := store[key]
	if !ok {
		if c.Parent != nil {
			return c.Parent.GetOk(key)
		}
	}
	return val, ok
}

// GetInt gets an int context value from req.
// Returns an empty string if key not found in the request context,
// or the value does not evaluate to a string
func (c *Context) GetInt(key interface{}) (int, bool) {
	value, ok := c.GetOk(key)
	if !ok {
		if c.Parent != nil {
			return c.Parent.GetInt(key)
		}
	}
	if num, ok := value.(int); ok {
		return num, ok
	}
	return 0, false
}

// GetString gets a string context value from req.
// Returns an empty string if key not found in the request context,
// or the value does not evaluate to a string
func (c *Context) GetString(key interface{}) string {
	store := c.getStore()
	if value, ok := store[key]; ok {
		if typed, ok := value.(string); ok {
			return typed
		}
	}
	if c.Parent != nil {
		return c.Parent.GetString(key)
	}
	return ""
}

// GetAll returns all stored context values for a request.
// Will always return a valid map. Returns an empty map for
// requests context data previously set
func (c *Context) GetAll() Store {
	buf := Store{}
	for key, value := range c.getStore() {
		buf[key] = value
	}
	if c.Parent != nil {
		for key, value := range c.Parent.GetAll() {
			buf[key] = value
		}
	}
	return buf
}

// Delete deletes a stored value from a request’s context
func (c *Context) Delete(key interface{}) {
	delete(c.getStore(), key)
}

// Clear clears all stored values in the current request’s context.
// Parent context store will not be cleaned.
func (c *Context) Clear() {
	store := c.getStore()
	for key := range store {
		delete(store, key)
	}
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
	c.Request = req.WithContext(c.Request.Context())
}

// Clone returns a clone of the current context.
func (c *Context) Clone() *Context {
	ctx := new(Context)
	*ctx = *c

	req := new(http.Request)
	*req = *c.Request
	ctx.Request = req
	c.CopyTo(ctx)

	res := new(http.Response)
	*res = *c.Response
	ctx.Response = res

	return ctx
}

// CopyTo copies the current context store into a new Context.
func (c *Context) CopyTo(newCtx *Context) {
	store := Store{}

	for key, value := range c.getStore() {
		store[key] = value
	}

	ctx := context.WithValue(context.Background(), Key, store)
	newCtx.Request = newCtx.Request.WithContext(ctx)
}

// Deadline returns the time when work done on behalf of this context
// should be canceled. Deadline returns ok==false when no deadline is
// set. Successive calls to Deadline return the same results.
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.Request.Context().Deadline()
}

// Done returns a channel that's closed when work done on behalf of this
// context should be canceled. Done may return nil if this context can
// never be canceled. Successive calls to Done return the same value.
//
// WithCancel arranges for Done to be closed when cancel is called;
// WithDeadline arranges for Done to be closed when the deadline
// expires; WithTimeout arranges for Done to be closed when the timeout
// elapses.
//
// Done is provided for use in select statements:
//
//  // Stream generates values with DoSomething and sends them to out
//  // until DoSomething returns an error or ctx.Done is closed.
//  func Stream(ctx context.Context, out chan<- Value) error {
//  	for {
//  		v, err := DoSomething(ctx)
//  		if err != nil {
//  			return err
//  		}
//  		select {
//  		case <-ctx.Done():
//  			return ctx.Err()
//  		case out <- v:
//  		}
//  	}
//  }
//
// See https://blog.golang.org/pipelines for more examples of how to use
// a Done channel for cancelation.
func (c *Context) Done() <-chan struct{} {
	return c.Request.Context().Done()
}

// Err returns a non-nil error value after Done is closed. Err returns
// Canceled if the context was canceled or DeadlineExceeded if the
// context's deadline passed. No other values for Err are defined.
// After Done is closed, successive calls to Err return the same value.
func (c *Context) Err() error {
	return c.Request.Context().Err()
}

// Value returns the value associated with this context for key, or nil
// if no value is associated with key. Successive calls to Value with
// the same key returns the same result.
//
// Use context values only for request-scoped data that transits
// processes and API boundaries, not for passing optional parameters to
// functions.
//
// A key identifies a specific value in a Context. Functions that wish
// to store values in Context typically allocate a key in a global
// variable then use that key as the argument to context.WithValue and
// Context.Value. A key can be any type that supports equality;
// packages should define keys as an unexported type to avoid
// collisions.
//
// Packages that define a Context key should provide type-safe accessors
// for the values stored using that key:
//
// 	// Package user defines a User type that's stored in Contexts.
// 	package user
//
// 	import "context"
//
// 	// User is the type of value stored in the Contexts.
// 	type User struct {...}
//
// 	// key is an unexported type for keys defined in this package.
// 	// This prevents collisions with keys defined in other packages.
// 	type key int
//
// 	// userKey is the key for user.User values in Contexts. It is
// 	// unexported; clients use user.NewContext and user.FromContext
// 	// instead of using this key directly.
// 	var userKey key = 0
//
// 	// NewContext returns a new Context that carries value u.
// 	func NewContext(ctx context.Context, u *User) context.Context {
// 		return context.WithValue(ctx, userKey, u)
// 	}
//
// 	// FromContext returns the User value stored in ctx, if any.
// 	func FromContext(ctx context.Context) (*User, bool) {
// 		u, ok := ctx.Value(userKey).(*User)
// 		return u, ok
// 	}
func (c *Context) Value(key interface{}) interface{} {
	return c.Request.Context().Value(key)
}

// emptyContext creates a new empty context.Context
func emptyContext() context.Context {
	return context.WithValue(context.Background(), Key, Store{})
}

// createRequest creates a default http.Request instance.
func createRequest() *http.Request {
	// Create HTTP request
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
	// Return shallow copy of Request with the new context
	return req.WithContext(emptyContext())
}

// createResponse creates a default http.Response instance.
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
