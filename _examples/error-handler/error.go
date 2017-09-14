package main

import (
	"errors"
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/context"
	"gopkg.in/h2non/gentleman.v2/plugins/headers"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Define a custom header
	cli.Use(headers.Set("Token", "s3cr3t"))

	// Declare first error phase middleware handler
	cli.UseError(func(ctx *context.Context, h context.Handler) {
		fmt.Printf("1) Handling error: %s\n", ctx.Error)
		h.Next(ctx)
	})

	// Declare second error phase middleware handler
	cli.UseError(func(ctx *context.Context, h context.Handler) {
		fmt.Printf("2) Handling error: %s\n", ctx.Error)
		// Overwrite error with wrapped message
		ctx.Error = errors.New("wrapped error: " + ctx.Error.Error())
		h.Next(ctx)
	})

	// Attach a phase-specific middleware function.
	cli.UseHandler("after dial", func(ctx *context.Context, h context.Handler) {
		ctx.Error = errors.New("simulated error")
		h.Next(ctx)
	})

	// Perform the request
	req := cli.Request()

	// Declare request-level error phase middleware handler
	req.UseError(func(ctx *context.Context, h context.Handler) {
		fmt.Printf("3) Handling request-level error: %s\n", ctx.Error)
		h.Next(ctx)
	})

	res, err := req.URL("http://httpbin.org/ip").Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Header: %s\n", res.Header.Get("Server"))
	fmt.Printf("Body: %s", res.String())
}
