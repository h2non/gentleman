package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/bodytype"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Define the JSON data to send
	data := `{"foo":"bar"}`
	cli.Use(body.String(data))

	// We're sending a JSON based payload
	cli.Use(bodytype.Type("json"))

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
