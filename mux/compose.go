package mux

import (
	c "gopkg.in/h2non/gentleman.v2/context"
)

// If creates a new multiplexer that will be executed if all the mux matchers passes.
func If(muxes ...*Mux) *Mux {
	mx := New()
	for _, mm := range muxes {
		mx.AddMatcher(mm.Matchers...)
	}
	return mx
}

// Or creates a new multiplexer that will be executed if at least one mux matcher passes.
func Or(muxes ...*Mux) *Mux {
	return Match(func(ctx *c.Context) bool {
		for _, mm := range muxes {
			if mm.Match(ctx) {
				return true
			}
		}
		return false
	})
}
