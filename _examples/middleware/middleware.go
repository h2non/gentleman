package main

import (
	"fmt"
	"net/url"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/context"
	"gopkg.in/h2non/gentleman.v2/plugins/headers"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Define a custom header
	cli.Use(headers.Set("Token", "s3cr3t"))

	// Attach a new midddleware function for request phase.
	cli.UseRequest(func(ctx *context.Context, h context.Handler) {
		u, _ := url.Parse("http://httpbin.org/headers")
		ctx.Request.URL = u
		h.Next(ctx)
	})

	// Attach a phase-specific middleware function.
	cli.UseHandler("after dial", func(ctx *context.Context, h context.Handler) {
		ctx.Response.Header.Set("Server", "go")
		h.Next(ctx)
	})

	// Attach a new midddleware function for response phase.
	cli.UseResponse(func(ctx *context.Context, h context.Handler) {
		ctx.Response.Header.Set("Server", "go "+ctx.Response.Header.Get("Server"))
		h.Next(ctx)
	})

	// Perform the request
	res, err := cli.Request().Send()
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
