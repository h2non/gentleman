package main

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v1"
	"gopkg.in/h2non/gentleman.v1/plugins/headers"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Define a custom header
	cli.Use(headers.Set("API-Token", "s3cr3t"))

	// Remove a header
	cli.Use(headers.Del("User-Agent"))

	// Perform the request
	res, err := cli.Request().URL("http://httpbin.org/headers").Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body: %s", res.String())
}
