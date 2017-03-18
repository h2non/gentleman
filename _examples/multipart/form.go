package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/multipart"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Define the generic base URL
	cli.URL("http://httpbin.org/post")

	// Create a new request
	req := cli.Request()

	// Create a text based form fields
	fields := map[string]string{"foo": "bar", "bar": "baz"}

	// Register the multipart plugin at request middleware level
	req.Use(multipart.Fields(fields))

	// Perform the request
	res, err := req.Send()
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
