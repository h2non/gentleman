# gentleman [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GitHub release](https://img.shields.io/github/tag/h2non/gentleman.svg)](https://github.com/h2non/gentleman/releases) [![GoDoc](https://godoc.org/github.com/h2non/gentleman?status.svg)](https://godoc.org/github.com/h2non/gentleman) [![Coverage Status](https://coveralls.io/repos/github/h2non/gentleman/badge.svg?branch=master)](https://coveralls.io/github/h2non/gentleman?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman)](https://goreportcard.com/report/github.com/h2non/gentleman)

Full-featyred, plugin-driven, middleware-oriented library to easily create rich, versatile and composable HTTP clients in [Go](http://golang.org).

Take a look to the [examples](#examples) to get started.

<img src="http://s10.postimg.org/5e31ox1ft/gentleman.png" align="center" height="320" />

## Goals

- Plugin driven.
- Simple, expressive, fluent API.
- Idiomatic built on top of `net/http` package.
- Middleware layer that supports the whole HTTP life cycle.
- Built-in multiplexer with easy composition features.
- Strong extensibility capabilities via plugins.
- Easy to configure and use.
- Convenient helpers and abstractions over Go's HTTP primitives.
- Built-in JSON, XML and multipart bodies serialization and parsing.
- Dependency free.

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v0
```

Note: I strongly recommend you to use `gopkg.in` when depending on third-party packages to prevent unexpected breaks of the interface contract in upcoming major versions of the package.

## Plugins

<table>
  <tr>
    <th>Name</th>
    <th>Docs</th>
    <th>API</th>
    <th>Status</th> 
    <th>Description</th>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/url">url</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/url">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/status-stable-green.svg?style=flat" /></td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /</a></td>
    <td>Easily declare URL, base URL and path values in HTTP requests</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/auth">auth</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/auth">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/status-stable-green.svg?style=flat" /></td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td> 
    <td>Declare authorization headers in your requests</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/body">body</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/body">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/status-stable-green.svg?style=flat" /></td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td> 
    <td>Easily define bodies based on JSON, XML, strings, buffers or streams</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/bodytype">bodytype</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/bodytype">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/status-stable-green.svg?style=flat" /></td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td> 
    <td>Define body MIME type by alias</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/cookies">cookies</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/cookies">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/status-stable-green.svg?style=flat" /></td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td> 
    <td>Declare and store HTTP cookies easily</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/compression">compression</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/compression">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/status-beta-green.svg?style=flat" /></td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td> 
    <td>Helpers to define enable/disable HTTP compression</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/headers">headers</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/headers">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/status-stable-green.svg?style=flat" /></td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td> 
    <td>Manage HTTP headers easily</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/multipart">multipart</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/multipart">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/status-stable-green.svg?style=flat" /></td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td> 
    <td>Create multipart forms easily. Supports files and text fields</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/proxy">proxy</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/proxy">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/status-stable-green.svg?style=flat" /></td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td> 
    <td>Configure HTTP proxy servers</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/query">query</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/query">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/status-stable-green.svg?style=flat" /></td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td> 
    <td>Easily manage query params</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/redirect">redirect</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/redirect">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/status-stable-green.svg?style=flat" /></td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td> 
    <td>Easily configure a custom redirect policy</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/timeout">timeout</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/timeout">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/status-stable-green.svg?style=flat" /></td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td> 
    <td>Easily configure the HTTP timeouts (request, dial, TLS...)</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/transport">transport</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/transport">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/status-stable-green.svg?style=flat" /></td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td> 
    <td>Define a custom HTTP transport easily</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/tls">tls</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/tls">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/status-stable-green.svg?style=flat" /></td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td> 
    <td>Configure the TLS options used by the HTTP transport</td>
  </tr>
</table>

[Send](https://github.com/h2non/gentleman/pull/new/master) a PR to add your plugin to the list.

### Creating plugins

You can create your own plugins for a variety of purposes, such as server discovery, custom HTTP tranport, modify any request/response param, intercept traffic, authentication and so on.

You can easily create custom plugins via [plugin](https://github.com/h2non/gentleman/tree/master/plugin) package.

See [plugin example](https://github.com/h2non/gentleman/blob/master/_examples/plugin/plugin.go).

## Middleware

gentleman is completely based on a hierarchical middleware layer based on simple plugin interfaces.

The middleware layer allows you to plug in intermediate custom logic in for any HTTP traffic handled by gentleman. It supports multiple phases, which represents the full request/response life cycle, giving you the ability to perform actions before and after an HTTP transaction is done.

#### Middleware phases

gentleman's dispatcher provides built-in support to the following middleware phases:

- **request** - Executed before a request is sent over the network.
- **response** - Executed when the client receives the response, even if it failed.
- **error** - Executed in case that an error ocurrs, support both injected or native error.
- **stop** - Executed in case that the request has been manually stopped via middleware (e.g: after interception).
- **intercept** - Executed in case that the request has been intercepted before network dialing.
- **before dial** - Executed before a request is sent over the network.
- **after dial** - Executed after the request dialing is done and the response has been received.

## API

See [godoc reference](https://godoc.org/github.com/h2non/gentleman) for detailed API documentation.

#### Subpackages

- [plugin](https://github.com/h2non/gentleman/tree/master/plugin) - [reference](https://godoc.org/github.com/h2non/gentleman/plugin) - Plugin layer for gentleman.
- [mux](https://github.com/h2non/gentleman/tree/master/mux) - [reference](https://godoc.org/github.com/h2non/gentleman/mux) - HTTP client multiplexer with built-in matchers.
- [middleware](https://github.com/h2non/gentleman/tree/master/middleware) - [reference](https://godoc.org/github.com/h2non/gentleman/middleware) - Middleware layer used by gentleman.
- [context](https://github.com/h2non/gentleman/tree/master/context) - [reference](https://godoc.org/github.com/h2non/gentleman/context) - HTTP context implementation for gentleman's middleware.
- [utils](https://github.com/h2non/gentleman/tree/master/utils) - [reference](https://godoc.org/github.com/h2non/gentleman/utils) - HTTP utilities internally used.

## Examples

See [examples](https://github.com/h2non/gentleman/blob/master/_examples) directory for featured examples.

#### Simple request

```go
package main

import (
  "fmt"
  "gopkg.in/h2non/gentleman.v0"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Define base URL
  cli.URL("http://httpbin.org")

  // Create a new request based on the current client
  req := cli.Request()

  // Define the URL path at request level
  req.Path("/headers")

  // Set a new header field
  req.SetHeader("Client", "gentleman")

  // Perform the request
  res, err := req.Send()
  if err != nil {
    fmt.Printf("Request error: %s\n", err)
    return
  }
  if !res.Ok {
    fmt.Printf("Invalid server response: %d\n", res.StatusCode)
    return
  }

  // Reads the whole body and returns it as string
  fmt.Printf("Body: %s", res.String())
}
```

#### Send JSON body

```go
package main

import (
  "fmt"
  "gopkg.in/h2non/gentleman.v0"
  "gopkg.in/h2non/gentleman.v0/plugins/body"
)

func main() {
  // Create a new client
  cli := gentleman.New()
  
  // Define the Base URL
  cli.URL("http://httpbin.org/post")

  // Create a new request based on the current client
  req := cli.Request()
  
  // Method to be used
  req.Method("POST")

  // Define the JSON payload via body plugin
  data := map[string]string{"foo": "bar"}
  req.Use(body.JSON(data))

  // Perform the request
  res, err := req.Send()
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
