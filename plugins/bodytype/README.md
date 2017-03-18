# gentleman/bodytype [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GoDoc](https://godoc.org/github.com/h2non/gentleman/plugins/bodytype?status.svg)](https://godoc.org/github.com/h2non/gentleman/plugins/bodytype) [![API](https://img.shields.io/badge/status-stable-green.svg?style=flat)](https://godoc.org/github.com/h2non/gentleman/plugins/bodytype) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman)](https://goreportcard.com/report/github.com/h2non/gentleman)

gentleman's plugin to easy define HTTP bodies. Supports JSON, XML, strings or streams with interface polymorphism. 

Supported type aliases:

- html       - `text/html`
- json       - `application/json`
- xml        - `application/xml`
- text       - `text/plain`
- urlencoded - `application/x-www-form-urlencoded`
- form       - `application/x-www-form-urlencoded`
- form-data  - `application/x-www-form-urlencoded`

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v2/plugins/bodytype
```

## API

See [godoc](https://godoc.org/github.com/h2non/gentleman/plugins/bodytype) reference.

## Example

```go
package main

import (
  "fmt"
  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman.v2/plugins/body"
  "gopkg.in/h2non/gentleman.v2/plugins/bodytype"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Define the JSON data to send 
  data := `{"foo":"bar"}`
  cli.Use(body.String(data))

  // We're sending a JSON based payload
  cli.Use(bodytype.Type("json"))

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
