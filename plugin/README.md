# gentleman/plugin [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GoDoc](https://godoc.org/github.com/h2non/gentleman/plugin?status.svg)](https://godoc.org/github.com/h2non/gentleman/plugin) [![API](https://img.shields.io/badge/status-stable-green.svg?style=flat)](https://godoc.org/github.com/h2non/gentleman/plugin) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman/plugin)](https://goreportcard.com/report/github.com/h2non/gentleman/plugin)

`plugin` package implements a simple middleware-based plugin layer especially designed for HTTP clients and complete HTTP request/response live cycle control.

The package exposes a simple API with multiple factory helper functions to create plugins more easily.

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v2/plugin
```

## API

See [godoc](https://godoc.org/github.com/h2non/gentleman/plugin) reference.

## Examples

#### Create a request plugin

```go
package main

import (
  "fmt"
  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman.v2/context"
  "gopkg.in/h2non/gentleman.v2/plugin"
  "net/url"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Create a request plugin to define the URL
  cli.Use(plugin.NewRequestPlugin(func(ctx *context.Context, h context.Handler) {
    u, _ := url.Parse("http://httpbin.org/headers")
    ctx.Request.URL = u
    h.Next(ctx)
  }))

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
