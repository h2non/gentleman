package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/query"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Define the base URL to use
	cli.Use(url.BaseURL("http://httpbin.org"))
	cli.Use(url.Path("/get"))

	// Define a custom query params
	cli.Use(query.Set("foo", "bar"))
	cli.Use(query.Set("bar", "baz"))

	// Or multiple values
	cli.Use(query.SetMap(map[string]string{"foo": "bar"}))

	// Remove a query param
	cli.Use(query.Del("bar"))

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
	fmt.Printf("Body: %s", res.String())
}
