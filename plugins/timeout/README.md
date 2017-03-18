# gentleman/timeout [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GoDoc](https://godoc.org/github.com/h2non/gentleman/plugins/timeout?status.svg)](https://godoc.org/github.com/h2non/gentleman/plugins/timeout) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman/plugins/timeout)](https://goreportcard.com/report/github.com/h2non/gentleman/plugins/timeout)
 
gentleman's plugin to easily define HTTP timeouts per specific phase in the HTTP connection live cycle.

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v2/plugins/timeout
```

## API

See [godoc](https://godoc.org/github.com/h2non/gentleman/plugins/timeout) reference.

## Example

```go
package main

import (
  "fmt"
  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman.v2/plugins/timeout"
  "time"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Define the max timeout for the whole HTTP request
  cli.Use(timeout.Request(10 * time.Second))

  // Define dial specific timeouts
  cli.Use(timeout.Dial(5*time.Second, 30*time.Second))

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
