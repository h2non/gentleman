package main

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v0"
	"gopkg.in/h2non/gentleman.v0/plugins/timeout"
	"time"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Define the max timeout for the whole HTTP request
	cli.Use(timeout.Request(10 * time.Second))

	// Define dial specific timeouts
	cli.Use(timeout.Dial(5*time.Second, 30*time.Second))

	// Perform the request
	res, err := cli.Request().URL("http://httpbin.org/headers").End()
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
