package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Define the base URL
	cli.Use(url.BaseURL("http://httpbin.org"))

	// Define the path with dynamic params
	// Use the :<name> notation
	cli.Use(url.Path("/:action/:subaction"))

	// Define the path params to replace
	cli.Use(url.Param("action", "delay"))
	cli.Use(url.Param("subaction", "1"))

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
