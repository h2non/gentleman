package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/auth"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Attach the plugin
	cli.Use(auth.Basic("user", "pas$w0rd"))

	// Perform the request
	res, err := cli.Request().URL("http://httpbin.org/headers").Send()
	if err != nil {
		fmt.Printf("Request error: %s", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d", res.StatusCode)
		return
	}

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body: %s", res.String())
}
