package mux

import (
	c "gopkg.in/h2non/gentleman.v0/context"
)

// If creates a new multiplexer that will be executed if all the matchers passes.
func If(matchers ...Matcher) *Mux {
	return Match(matchers...)
}

// Or creates a new multiplexer that will be executed if at least one matcher passes.
func Or(matchers ...Matcher) *Mux {
	return Match(func(ctx *c.Context) bool {
		for _, matcher := range matchers {
			if matcher(ctx) {
				return true
			}
		}
		return false
	})
}
