package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/redirect"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Set authorization header
	cli.SetHeader("Authorization", "s3cr3t!")

	// Define the maximum number of redirects
	cli.Use(redirect.Limit(10))

	// Perform the request
	res, err := cli.Request().URL("http://httpbin.org/relative-redirect/3").Send()
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
