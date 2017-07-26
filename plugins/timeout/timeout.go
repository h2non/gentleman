package timeout

import (
	g "gopkg.in/h2non/gentleman.v2"
	c "gopkg.in/h2non/gentleman.v2/context"
	p "gopkg.in/h2non/gentleman.v2/plugin"
	"net"
	"net/http"
	"time"
)

// Timeouts represents the supported timeouts
type Timeouts struct {
	// Request represents the total timeout including dial / request / redirect steps
	Request time.Duration

	// TLS represents the maximum amount of time for TLS handshake process
	TLS time.Duration

	// Dial represents the maximum amount of time for dialing process
	Dial time.Duration

	// KeepAlive represents the maximum amount of time to keep alive the socket
	KeepAlive time.Duration
}

// Request defines the maximum amount of time a whole request process
// (including dial / request / redirect) can take.
func Request(timeout time.Duration) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		ctx.Client.Timeout = timeout
		h.Next(ctx)
	})
}

// TLS defines the maximum amount of time waiting for a TLS handshake
func TLS(timeout time.Duration) p.Plugin {
	return All(Timeouts{TLS: timeout})
}

// Dial defines the maximum amount of time waiting for network dialing
func Dial(timeout, keepAlive time.Duration) p.Plugin {
	return All(Timeouts{Dial: timeout, KeepAlive: keepAlive})
}

// All defines all the timeout types for the outgoing request
func All(timeouts Timeouts) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		defineTimeouts(timeouts, ctx)
		h.Next(ctx)
	})
}

func defineTimeouts(timeouts Timeouts, ctx *c.Context) {
	if timeouts.Request == 0 {
		timeouts.Request = g.RequestTimeout
	}
	ctx.Client.Timeout = timeouts.Request

	// Assert http.Transport to work with the instance
	transport, ok := ctx.Client.Transport.(*http.Transport)
	if !ok {
		// If using custom transport, just ignore it
		return
	}

	if timeouts.TLS == 0 {
		timeouts.TLS = g.TLSHandshakeTimeout
	}
	transport.TLSHandshakeTimeout = timeouts.TLS

	if timeouts.Dial == 0 {
		timeouts.Dial = g.DialTimeout
	}
	if timeouts.KeepAlive == 0 {
		timeouts.KeepAlive = g.DialKeepAlive
	}

	transport.Dial = (&net.Dialer{
		Timeout:   timeouts.Dial,
		KeepAlive: timeouts.KeepAlive,
	}).Dial

	// Finally expose the transport to be used
	ctx.Client.Transport = transport
}
