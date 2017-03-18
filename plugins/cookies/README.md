# gentleman/cookies [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GoDoc](https://godoc.org/github.com/h2non/gentleman/plugins/cookies?status.svg)](https://godoc.org/github.com/h2non/gentleman/plugins/cookies) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman)](https://goreportcard.com/report/github.com/h2non/gentleman)

gentleman's plugin to easily deal and manage cookies HTTP clients.

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v2/plugins/cookies
```

## API

See [godoc](https://godoc.org/github.com/h2non/gentleman/plugins/cookies) reference.

## Example

```go
package main

import (
  "fmt"
  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman.v2/plugins/cookies"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Define cookies
  cli.Use(cookies.Set("foo", "bar"))

  // Configure cookie jar store
  cli.Use(cookies.Jar())

  // Perform the request
  res, err := cli.Request().URL("http://httpbin.org/cookies").Send()
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
