package main

import (
	"fmt"
	"time"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/timeout"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Define the max timeout for the whole HTTP request
	cli.Use(timeout.Request(10 * time.Second))

	// Define dial specific timeouts
	cli.Use(timeout.Dial(5*time.Second, 30*time.Second))

	// Perform the request
	res, err := cli.Request().URL("http://httpbin.org/headers").Send()
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
