# gentleman/compression [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GoDoc](https://godoc.org/github.com/h2non/gentleman/plugins/compression?status.svg)](https://godoc.org/github.com/h2non/gentleman/plugins/compression) [![API](https://img.shields.io/badge/status-beta-green.svg?style=flat)](https://godoc.org/github.com/h2non/gentleman/plugins/compression) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman)](https://goreportcard.com/report/github.com/h2non/gentleman)

gentleman's plugin to disable and customize data compression in HTTP requests/responses.

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v2/plugins/compression
```

## API

See [godoc](https://godoc.org/github.com/h2non/gentleman/plugins/compression) reference.

## Example

```go
package main

import (
  "fmt"
  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman.v2/plugins/compression"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Disable HTTP compression
  cli.Use(compression.Disable())

  // Perform the request
  res, err := cli.Request().URL("http://httpbin.org/gzip").Send()
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
