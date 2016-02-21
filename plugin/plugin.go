package plugin

import (
	"gopkg.in/h2non/gentleman.v0/context"
)

// Plugin interface that must be implemented by plugins
type Plugin interface {
	// Enable enabled the plugin
	Enable()

	// Disable disables the plugin
	Disable()

	// Disabled returns true if the plugin is enabled
	Disabled() bool

	// Remove will remove the plugin from the middleware stack
	Remove()

	// Enabled returns true if the plugin was removed
	Removed() bool

	// Error phase middleware handler
	Error(*context.Context, context.Handler)

	// Request phase middleware handler
	Request(*context.Context, context.Handler)

	// Response phase middleware handler
	Response(*context.Context, context.Handler)
}

// Layer encapsulates an Error, Request and Response function handlers
type Layer struct {
	// removed stores if the plugin was removed
	removed bool

	// disabled stores if the plugin was disabled
	disabled bool

	// ErrorHandler function handler for the error phase
	ErrorHandler context.HandlerFunc

	// RequestHandler function handler for the request phase
	RequestHandler context.HandlerFunc

	// ResponseHandler function handler for the response phase
	ResponseHandler context.HandlerFunc
}

// Disable will disable the current plugin
func (p *Layer) Disable() {
	p.disabled = true
}

// Enable will enable the current plugin
func (p *Layer) Enable() {
	p.disabled = false
}

// Disabled returns true if the plugin is enabled
func (p *Layer) Disabled() bool {
	return p.disabled
}

// Remove will remove the plugin from the middleware stack
func (p *Layer) Remove() {
	p.removed = true
}

// Removed returns true if the plugin Was removed
func (p *Layer) Removed() bool {
	return p.removed
}

// Response triggers the response middleware phase
func (p *Layer) Response(ctx *context.Context, handler context.Handler) {
	p.call(ctx, handler, p.ResponseHandler)
}

// Request triggers the request middleware phase
func (p *Layer) Request(ctx *context.Context, handler context.Handler) {
	p.call(ctx, handler, p.RequestHandler)
}

// Error triggers the error middleware phase
func (p *Layer) Error(ctx *context.Context, handler context.Handler) {
	p.call(ctx, handler, p.ErrorHandler)
}

func (p *Layer) call(ctx *context.Context, h context.Handler, fn context.HandlerFunc) {
	if fn == nil {
		h.Next(ctx)
		return
	}
	if p.disabled || p.removed {
		h.Next(ctx)
		return
	}
	fn(ctx, h)
}

// NewResponsePlugin creates a new plugin layer
// to handle response middleware phase
func NewResponsePlugin(handler context.HandlerFunc) Plugin {
	return &Layer{ResponseHandler: handler}
}

// NewRequestPlugin creates a new plugin layer
// to handle request middleware phase
func NewRequestPlugin(handler context.HandlerFunc) Plugin {
	return &Layer{RequestHandler: handler}
}

// NewErrorPlugin creates a new plugin layer
// to handle error middleware phase
func NewErrorPlugin(handler context.HandlerFunc) Plugin {
	return &Layer{ErrorHandler: handler}
}

// New creates a new plugin based on phase-specific function handlers
func New(res, req, err context.HandlerFunc) Plugin {
	return &Layer{
		ResponseHandler: res,
		RequestHandler:  req,
		ErrorHandler:    err,
	}
}
