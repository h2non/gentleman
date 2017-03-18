# gentleman/mux [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GoDoc](https://godoc.org/github.com/h2non/gentleman/mux?status.svg)](https://godoc.org/github.com/h2non/gentleman/mux) [![API](https://img.shields.io/badge/status-stable-green.svg?style=flat)](https://godoc.org/github.com/h2non/gentleman/mux) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman/mux)](https://goreportcard.com/report/github.com/h2non/gentleman/mux)

`mux` package implements a versatile HTTP client multiplexer with built-in matchers for easy plugin composition.

multiplexer can be used to compose plugins for both request/response phases.

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v2/mux
```

## API

See [godoc](https://godoc.org/github.com/h2non/gentleman/mux) reference.

## Example

Create a multiplexer filtered by a custom matcher function:
```go
package main

import (
  "fmt"
  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman.v2/mux"
  "gopkg.in/h2non/gentleman.v2/context"
  "gopkg.in/h2non/gentleman.v2/plugins/url"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Use a custom multiplexer for GET requests
  cli.Use(mux.New().AddMatcher(func (ctx *context.Context) bool {
    return ctx.GetString("$phase") == "request" && ctx.Request.Method == "GET"
  }).Use(url.URL("http://httpbin.org/headers")))

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

Plugin composition via multiplexer:
```go
package main

import (
  "fmt"
  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman.v2/mux"
  "gopkg.in/h2non/gentleman.v2/plugins/url"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Define the server url (must be first)
  cli.Use(url.URL("http://httpbin.org"))

  // Create a new multiplexer based on multiple matchers
  mx := mux.If(mux.Method("GET"), mux.Host("httpbin.org"))

  // Attach a custom plugin on the multiplexer that will be executed if the matchers passes
  mx.Use(url.Path("/headers"))

  // Attach the multiplexer on the main client
  cli.Use(mx)

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
