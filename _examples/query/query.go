package main

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v0"
	"gopkg.in/h2non/gentleman.v0/plugins/query"
	"gopkg.in/h2non/gentleman.v0/plugins/url"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Define the base URL to use
	cli.Use(url.BaseURL("http://httpbin.org"))
	cli.Use(url.Path("/get"))

	// Define a custom query param
	cli.Use(query.Set("foo", "bar"))

	// Remove a query param
	cli.Use(query.Del("bar"))

	// Perform the request
	res, err := cli.Request().End()
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
