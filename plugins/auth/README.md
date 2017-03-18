# gentleman/auth [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GoDoc](https://godoc.org/github.com/h2non/gentleman?status.svg)](https://godoc.org/github.com/h2non/gentleman/plugins/auth) [![API](https://img.shields.io/badge/status-stable-green.svg?style=flat)](https://godoc.org/github.com/h2non/gentleman/plugins/auth) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman)](https://goreportcard.com/report/github.com/h2non/gentleman)

gentleman's plugin to easily define HTTP authorization headers based on multiple schemas.

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v2/plugins/auth
```

## API

See [godoc](https://godoc.org/github.com/h2non/gentleman/plugins/auth) reference.

## Example

```go
package main

import (
  "fmt"
  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman.v2/plugins/auth"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Attach the plugin at client level
  cli.Use(auth.Basic("user", "pas$w0rd"))

  // Perform the request
  res, err := cli.Request().Method().URL("http://httpbin.org/headers").Send()
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
