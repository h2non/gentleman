// Package plugin implements a plugin layer for gentleman components.
// Exports the required interface that must be implemented by plugins.
//
// Plugins are phase-oriented middleware function handlers encapsulated
// in a simple interface that will be consumed by the middleware layer in
// order to trigger the plugin handler.
//
// Plugin implementors can decide to build a plugin to handle a unique
// middleware phase or instead handle multiple phases: request, response, error...
package plugin

import "gopkg.in/h2non/gentleman.v2/context"

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

	// Exec executes the plugin handler for a specific middleware phase.
	Exec(string, *context.Context, context.Handler)
}

// Handlers represents a map to store middleware handler functions per phase.
type Handlers map[string]context.HandlerFunc

// Layer encapsulates an Error, Request and Response function handlers
type Layer struct {
	// removed stores if the plugin was removed
	removed bool

	// disabled stores if the plugin was disabled
	disabled bool

	// Handlers defines the required handlers
	Handlers Handlers

	// DefaultHandler is an optional field used to store
	// a default handler for any middleware phase.
	DefaultHandler context.HandlerFunc
}

// New creates a new plugin layer.
func New() *Layer {
	return &Layer{Handlers: make(Handlers)}
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

// SetHandler uses a new handler function for the given middleware phase.
func (p *Layer) SetHandler(phase string, handler context.HandlerFunc) {
	p.Handlers[phase] = handler
}

// SetHandlers uses a new map of handler functions.
func (p *Layer) SetHandlers(handlers Handlers) {
	p.Handlers = handlers
}

// Exec executes the plugin handler for the given middleware phase passing the given context.
func (p *Layer) Exec(phase string, ctx *context.Context, h context.Handler) {
	if p.disabled || p.removed {
		h.Next(ctx)
		return
	}

	fn := p.Handlers[phase]
	if fn == nil {
		fn = p.DefaultHandler
	}
	if fn == nil {
		h.Next(ctx)
		return
	}

	fn(ctx, h)
}

// NewPhasePlugin creates a new plugin layer
// to handle a given middleware phase.
func NewPhasePlugin(phase string, handler context.HandlerFunc) Plugin {
	return &Layer{Handlers: Handlers{phase: handler}}
}

// NewResponsePlugin creates a new plugin layer
// to handle response middleware phase
func NewResponsePlugin(handler context.HandlerFunc) Plugin {
	return NewPhasePlugin("response", handler)
}

// NewRequestPlugin creates a new plugin layer
// to handle request middleware phase
func NewRequestPlugin(handler context.HandlerFunc) Plugin {
	return NewPhasePlugin("request", handler)
}

// NewErrorPlugin creates a new plugin layer
// to handle error middleware phase
func NewErrorPlugin(handler context.HandlerFunc) Plugin {
	return NewPhasePlugin("error", handler)
}
