# gentleman/url [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GoDoc](https://godoc.org/github.com/h2non/gentleman/plugins/url?status.svg)](https://godoc.org/github.com/h2non/gentleman/plugins/url) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman/plugins/url)](https://goreportcard.com/report/github.com/h2non/gentleman/plugins/url)

gentleman's plugin to easily define URL fields in HTTP requests.

Supports full URL parsing, base URL, base path, full path and dynamic path params templating.

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v2/plugins/url
```

## API

See [godoc](https://godoc.org/github.com/h2non/gentleman/plugins/url) reference.

## Example

```go
package main

import (
  "fmt"
  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman.v2/plugins/url"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Define the base URL
  cli.Use(url.BaseURL("http://httpbin.org"))

  // Define the path with dynamic value
  cli.Use(url.Path("/:resource"))

  // Define the path value to be replaced
  cli.Use(url.Param("resource", "get"))

  // Perform the request
  res, err := cli.Request().Send()
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
```

## License

MIT - Tomas Aparicio
