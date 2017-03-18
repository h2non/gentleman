package main

import (
	"fmt"
	"sync"

	"gopkg.in/h2non/gentleman.v2"
)

func main() {
	reqs := []string{
		"/headers",
		"/delay/1",
		"/get",
		"/ip",
	}

	// Create a new client
	cli := gentleman.New()

	// Define the base URL
	cli.BaseURL("http://httpbin.org")

	// Create a sync group to wait for goroutines
	var wg sync.WaitGroup
	wg.Add(len(reqs))

	// Fetch resources in parallel
	for _, path := range reqs {
		go fetch(cli, path, &wg)
	}

	wg.Wait()
	fmt.Printf("Done!\n")
}

func fetch(cli *gentleman.Client, path string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Perform the request
	res, err := cli.Request().Path(path).Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	fmt.Printf("Path: %s => Response: %d\n", path, res.StatusCode)
}
