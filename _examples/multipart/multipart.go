package main

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v0"
	"gopkg.in/h2non/gentleman.v0/plugins/multipart"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Create a text based form fields
	fields := map[string]string{"foo": "bar", "bar": "baz"}
	cli.Use(multipart.Fields(fields))

	// Perform the request
	res, err := cli.Request().Method("POST").URL("http://httpbin.org/post").End()
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
