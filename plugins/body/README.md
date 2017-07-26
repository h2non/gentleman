# gentleman/body [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GoDoc](https://godoc.org/github.com/h2non/gentleman/plugins/body?status.svg)](https://godoc.org/github.com/h2non/gentleman/plugins/body) [![API](https://img.shields.io/badge/status-stable-green.svg?style=flat)](https://godoc.org/github.com/h2non/gentleman) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman)](https://goreportcard.com/report/github.com/h2non/gentleman)

gentleman's plugin to easy define HTTP bodies. Supports JSON, XML, strings or streams with interface polymorphism. 

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v2/plugins/body
```

## API

See [godoc](https://godoc.org/github.com/h2non/gentleman/plugins/body) reference.

## Example

```go
package main

import (
  "fmt"
  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman.v2/plugins/body"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Define the body we're going to send
  data := map[string]string{"foo": "bar"}
  cli.Use(body.JSON(data))

  // Perform the request
  res, err := cli.Request().Method("POST").URL("http://httpbin.org/post").Send()
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
