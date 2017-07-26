# gentleman/transport [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GoDoc](https://godoc.org/github.com/h2non/gentleman/plugins/transport?status.svg)](https://godoc.org/github.com/h2non/gentleman/plugins/transport) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman/plugins/transport)](https://goreportcard.com/report/github.com/h2non/gentleman/plugins/transport)

gentleman's plugin to easily define the HTTP transport to be used by `http.Client`.

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v2/plugins/transport
```

## API

See [godoc](https://godoc.org/github.com/h2non/gentleman/plugins/transport) reference.

## Example

```go
package main

import (
  "fmt"
  "net/http"
  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman.v2/plugins/transport"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Define the default HTTP transport
  cli.Use(transport.Set(http.DefaultTransport))

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
