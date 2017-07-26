package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Create a new request based on the current client
	req := cli.Request().URL("http://httpbin.org/get")

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

	type jsonData struct {
		URL     string            `json:url`
		Origin  string            `json:origin`
		Headers map[string]string `json:headers`
	}
	json := &jsonData{}

	// Parse the body and map into a struct
	res.JSON(json)
	fmt.Printf("Body: %#v\n", json)
}
