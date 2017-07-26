# gentleman/headers [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GoDoc](https://godoc.org/github.com/h2non/gentleman/plugins/headers?status.svg)](https://godoc.org/github.com/h2non/gentleman/plugins/headers) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman)](https://goreportcard.com/report/github.com/h2non/gentleman)

gentleman's plugin to easily manage HTTP headers.

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v2/plugins/headers
```

## API

See [godoc](https://godoc.org/github.com/h2non/gentleman/plugins/headers) reference.

## Example

```go
package main

import (
  "fmt"
  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman.v2/plugins/headers"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Define a custom header
  cli.Use(headers.Set("API-Token", "s3cr3t"))

  // Remove a header
  cli.Use(headers.Del("User-Agent"))

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
```

## License

MIT - Tomas Aparicio
