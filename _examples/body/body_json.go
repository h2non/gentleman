package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Define the Base URL
	cli.URL("http://httpbin.org/post")

	// Create a new request based on the current client
	req := cli.Request()

	// Method to be used
	req.Method("POST")

	// Define the JSON payload via body plugin
	data := map[string]string{"foo": "bar"}
	req.Use(body.JSON(data))

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
