# gentleman/proxy [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GoDoc](https://godoc.org/github.com/h2non/gentleman/plugins/proxy?status.svg)](https://godoc.org/github.com/h2non/gentleman/plugins/proxy) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman)](https://goreportcard.com/report/github.com/h2non/gentleman)

gentleman's plugin to easily manage HTTP proxies used by clients.

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v2/plugins/proxy
```

## API

See [godoc](https://godoc.org/github.com/h2non/gentleman/plugins/proxy) reference.

## Example

```go
package main

import (
  "fmt"
  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman.v2/plugins/proxy"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Define a list of HTTP proxies to be used
  cli.Use(proxy.Set([]string{"http://proxy:3128", "http://proxy2:3128"}))

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
