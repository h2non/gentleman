# gentleman/headers [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GitHub release](https://img.shields.io/github/tag/h2non/gentleman.svg)](https://github.com/h2non/gentleman/releases) [![GoDoc](https://godoc.org/github.com/h2non/gentleman/plugins/headers?status.svg)](https://godoc.org/github.com/h2non/gentleman/plugins/headers) [![API](https://img.shields.io/badge/api-beta-green.svg?style=flat)](https://godoc.org/github.com/h2non/gentleman/plugins/headers)

gentleman's plugin to easily manage HTTP headers.

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v0/plugins/headers
```

## API

See [godoc](https://godoc.org/github.com/h2non/gentleman/plugins/headers) reference.

## Example

```go
package main

import (
  "fmt"
  "gopkg.in/h2non/gentleman.v0"
  "gopkg.in/h2non/gentleman.v0/plugins/headers"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Define a custom header
  cli.Use(headers.Set("API-Token", "s3cr3t"))

  // Remove a header
  cli.Use(headers.Del("User-Agent"))

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
```

## License

MIT - Tomas Aparicio
