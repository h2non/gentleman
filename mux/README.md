# gentleman/mux [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GoDoc](https://godoc.org/github.com/h2non/gentleman/mux?status.svg)](https://godoc.org/github.com/h2non/gentleman/mux) [![API](https://img.shields.io/badge/status-stable-green.svg?style=flat)](https://godoc.org/github.com/h2non/gentleman/mux) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman/mux)](https://goreportcard.com/report/github.com/h2non/gentleman/mux)

`mux` package implements a versatile HTTP client multiplexer with built-in matchers for easy plugin composition.

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v0/mux
```

## API

See [godoc](https://godoc.org/github.com/h2non/gentleman/mux) reference.

## Example

```go
package main

import (
  "fmt"
  "gopkg.in/h2non/gentleman.v0"
  "gopkg.in/h2non/gentleman.v0/mux"
  "gopkg.in/h2non/gentleman.v0/context"
  "gopkg.in/h2non/gentleman.v0/plugins/url"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Use a custom multiplexer for GET requests
  cli.Use(mux.New().AddMatcher(func (ctx *context.Context) bool {
    return ctx.GetString("$phase") == "request" && ctx.Request.Method == "GET"
  }).Use(url.URL("http://httpbin.org/headers")))

  // Perform the request
  res, err := cli.Request().End()
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
