# gentleman/multipart [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GoDoc](https://godoc.org/github.com/h2non/gentleman/plugins/multipart?status.svg)](https://godoc.org/github.com/h2non/gentleman/plugins/multipart) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman)](https://goreportcard.com/report/github.com/h2non/gentleman)

gentleman's plugin to easily define `multipart/form-data` bodies supporting files and string based fields.

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v2/plugins/multipart
```

## API

See [godoc](https://godoc.org/github.com/h2non/gentleman/plugins/multipart) reference.

## Example

```go
package main

import (
  "fmt"
  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman.v2/plugins/multipart"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Create a text based form fields
  fields := map[string]string{"foo": "bar", "bar": "baz"}
  cli.Use(multipart.Fields(fields))

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
