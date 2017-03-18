package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/headers"
)

func main() {
	// Create a parent client
	parent := gentleman.New()

	// Define default URL
	parent.URL("http://httpbin.org")

	// Define a custom header via parent client
	parent.Use(headers.Set("API-Token", "s3cr3t"))

	// Create a new client
	cli := gentleman.New()

	// Bind parent client
	cli.UseParent(parent)

	// Create a new request based on the current client
	req := cli.Request()

	// Perform the request
	res, err := req.Path("/post").JSON(map[string]string{"foo": "bar"}).Send()
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
