package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/headers"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Define a global header at client level
	cli.SetHeader("Version", "1.0")

	// Define a custom header (via headers plugin)
	cli.Use(headers.Set("API-Token", "s3cr3t"))

	// Remove a header (via headers plugin)
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
