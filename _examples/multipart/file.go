package main

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v0"
	"gopkg.in/h2non/gentleman.v0/plugins/multipart"
	"os"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Read a file from disk and post it
	file, _ := os.Open("LICENSE")
	cli.Use(multipart.File("license", file))
	defer file.Close()

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
