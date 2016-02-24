package main

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v0"
	"gopkg.in/h2non/gentleman.v0/plugins/body"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Attach the plugin at client level
	data := map[string]string{"foo": "bar"}
	cli.Use(body.JSON(data))

	// Perform the request
	res, err := cli.Request().Method("POST").URL("http://httpbin.org/post").Send()
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
