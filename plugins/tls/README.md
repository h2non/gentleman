# gentleman/tls [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GoDoc](https://godoc.org/github.com/h2non/gentleman/plugins/tls?status.svg)](https://godoc.org/github.com/h2non/gentleman/plugins/tls) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman)](https://goreportcard.com/report/github.com/h2non/gentleman)

gentleman's plugin to easily define TLS config used by `http.Transport`/`RoundTripper` interface.

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v2/plugins/tls
```

## API

See [godoc](https://godoc.org/github.com/h2non/gentleman/plugins/tls) reference.

## Example

```go
package main

import (
  "fmt"
  "crypto/tls"
  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman.v2/plugins/tls"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Define a custom header
  cli.Use(tls.Config(&tls.Config{ServerName: "foo.com"}))

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
