# gentleman [![Build Status](https://travis-ci.org/h2non/gentleman.svg)](https://travis-ci.org/h2non/gentleman) [![GitHub release](https://img.shields.io/badge/version-2.0.2-orange.svg?style=flat)](https://github.com/h2non/gentleman/releases) [![GoDoc](https://godoc.org/gopkg.in/h2non/gentleman.v2?status.svg)](https://godoc.org/gopkg.in/h2non/gentleman.v2) [![Coverage Status](https://coveralls.io/repos/github/h2non/gentleman/badge.svg?branch=master)](https://coveralls.io/github/h2non/gentleman?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman)](https://goreportcard.com/report/github.com/h2non/gentleman) [![Go Version](https://img.shields.io/badge/go-v2.0+-green.svg?style=flat)](https://github.com/h2non/gentleman)

Full-featured, plugin-driven, middleware-oriented toolkit to easily create rich, versatile and composable HTTP clients in [Go](http://golang.org).

<img src="http://s10.postimg.org/5e31ox1ft/gentleman.png" align="right" height="260" />

gentleman embraces extensibility and composition principles in order to provide a flexible way to easily create featured HTTP client layers based on built-in or third-party plugins that you can register and reuse across HTTP clients.

As an example, you can easily provide retry policy capabilities or dynamic server discovery in your HTTP clients simply attaching the [retry](https://github.com/h2non/gentleman-retry) or [consul](https://github.com/h2non/gentleman-consul) plugins.

Take a look to the [examples](#examples), list of [supported plugins](#plugins), [HTTP entities](#http-entities) or [middleware layer](#middleware) to get started.

For testing purposes, see [baloo](https://github.com/h2non/baloo), an utility library for expressive end-to-end HTTP API testing, built on top of `gentleman` toolkit. For HTTP mocking, see [gentleman-mock](https://github.com/h2non/gentleman-mock), which uses [gock](https://github.com/h2non/gock) under the hood for easy and expressive HTTP client request mocking.

<a target='_blank' rel='nofollow' href='https://app.codesponsor.io/link/1MpD3pzt63uUeG43NP4tDHPY/h2non/gentleman'>  <img alt='Sponsor' width='888' height='68' src='https://app.codesponsor.io/embed/1MpD3pzt63uUeG43NP4tDHPY/h2non/gentleman.svg' /></a>

## Versions

- [v2](https://github.com/h2non/gentleman/) - Latest version. Stable. Recommended.
- [v1](https://github.com/h2non/gentleman/tree/v1) - First version. Stable. Actively maintained.

## Features

- Plugin driven architecture.
- Simple, expressive, fluent API.
- Idiomatic built on top of `net/http` package.
- Context-aware hierarchical middleware layer supporting all the HTTP life cycle.
- Built-in multiplexer for easy composition capabilities.
- Easy to extend via plugins/middleware.
- Ability to easily intercept and modify HTTP traffic on-the-fly.
- Convenient helpers and abstractions over Go's HTTP primitives.
- URL template path params.
- Built-in JSON, XML and multipart bodies serialization and parsing.
- Easy to test via HTTP mocking (e.g: [gentleman-mock](https://github.com/h2non/gentleman-mock)).
- Supports data passing across plugins/middleware via its built-in context.
- Fits good while building domain-specific HTTP API clients.
- Easy to hack.
- Dependency free.

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v2
```

## Requirements

- Go 1.7+

## Plugins

<table>
  <tr>
    <th>Name</th>
    <th>Docs</th>
    <th>Status</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/url">url</a></td>
    <td>
      <a href="https://godoc.org/gopkg.in/h2non/gentleman.v2/plugins/url">
        <img src="https://godoc.org/gopkg.in/h2non/gentleman.v2?status.svg" />
      </a>
    </td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /</a></td>
    <td>Easily declare URL, base URL and path values in HTTP requests</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/auth">auth</a></td>
    <td>
      <a href="https://godoc.org/gopkg.in/h2non/gentleman.v2/plugins/auth">
        <img src="https://godoc.org/gopkg.in/h2non/gentleman.v2?status.svg" />
      </a>
    </td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td>
    <td>Declare authorization headers in your requests</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/body">body</a></td>
    <td>
      <a href="https://godoc.org/gopkg.in/h2non/gentleman.v2/plugins/body">
        <img src="https://godoc.org/gopkg.in/h2non/gentleman.v2?status.svg" />
      </a>
    </td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td>
    <td>Easily define bodies based on JSON, XML, strings, buffers or streams</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/bodytype">bodytype</a></td>
    <td>
      <a href="https://godoc.org/gopkg.in/h2non/gentleman.v2/plugins/bodytype">
        <img src="https://godoc.org/gopkg.in/h2non/gentleman.v2?status.svg" />
      </a>
    </td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td>
    <td>Define body MIME type by alias</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/cookies">cookies</a></td>
    <td>
      <a href="https://godoc.org/gopkg.in/h2non/gentleman.v2/plugins/cookies">
        <img src="https://godoc.org/gopkg.in/h2non/gentleman.v2?status.svg" />
      </a>
    </td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td>
    <td>Declare and store HTTP cookies easily</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/compression">compression</a></td>
    <td>
      <a href="https://godoc.org/gopkg.in/h2non/gentleman.v2/plugins/compression">
        <img src="https://godoc.org/gopkg.in/h2non/gentleman.v2?status.svg" />
      </a>
    </td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td>
    <td>Helpers to define enable/disable HTTP compression</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/headers">headers</a></td>
    <td>
      <a href="https://godoc.org/gopkg.in/h2non/gentleman.v2/plugins/headers">
        <img src="https://godoc.org/gopkg.in/h2non/gentleman.v2?status.svg" />
      </a>
    </td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td>
    <td>Manage HTTP headers easily</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/multipart">multipart</a></td>
    <td>
      <a href="https://godoc.org/gopkg.in/h2non/gentleman.v2/plugins/multipart">
        <img src="https://godoc.org/gopkg.in/h2non/gentleman.v2?status.svg" />
      </a>
    </td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td>
    <td>Create multipart forms easily. Supports files and text fields</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/proxy">proxy</a></td>
    <td>
      <a href="https://godoc.org/gopkg.in/h2non/gentleman.v2/plugins/proxy">
        <img src="https://godoc.org/gopkg.in/h2non/gentleman.v2?status.svg" />
      </a>
    </td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td>
    <td>Configure HTTP proxy servers</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/query">query</a></td>
    <td>
      <a href="https://godoc.org/gopkg.in/h2non/gentleman.v2/plugins/query">
        <img src="https://godoc.org/gopkg.in/h2non/gentleman.v2?status.svg" />
      </a>
    </td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td>
    <td>Easily manage query params</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/redirect">redirect</a></td>
    <td>
      <a href="https://godoc.org/gopkg.in/h2non/gentleman.v2/plugins/redirect">
        <img src="https://godoc.org/gopkg.in/h2non/gentleman.v2?status.svg" />
      </a>
    </td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td>
    <td>Easily configure a custom redirect policy</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/timeout">timeout</a></td>
    <td>
      <a href="https://godoc.org/gopkg.in/h2non/gentleman.v2/plugins/timeout">
        <img src="https://godoc.org/gopkg.in/h2non/gentleman.v2?status.svg" />
      </a>
    </td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td>
    <td>Easily configure the HTTP timeouts (request, dial, TLS...)</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/transport">transport</a></td>
    <td>
      <a href="https://godoc.org/gopkg.in/h2non/gentleman.v2/plugins/transport">
        <img src="https://godoc.org/gopkg.in/h2non/gentleman.v2?status.svg" />
      </a>
    </td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td>
    <td>Define a custom HTTP transport easily</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/tls">tls</a></td>
    <td>
      <a href="https://godoc.org/gopkg.in/h2non/gentleman.v2/plugins/tls">
        <img src="https://godoc.org/gopkg.in/h2non/gentleman.v2?status.svg" />
      </a>
    </td>
    <td><a href="https://travis-ci.org/h2non/gentleman"><img src="https://travis-ci.org/h2non/gentleman.png" /></a></td>
    <td>Configure the TLS options used by the HTTP transport</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman-retry">retry</a></td>
    <td>
      <a href="https://godoc.org/gopkg.in/h2non/gentleman.v2-retry">
        <img src="https://godoc.org/gopkg.in/h2non/gentleman.v2-retry?status.svg" />
      </a>
    </td>
    <td><a href="https://travis-ci.org/h2non/gentleman-retry"><img src="https://travis-ci.org/h2non/gentleman-retry.png" /></a></td>
    <td>Provide retry policy capabilities to your HTTP clients</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman-mock">mock</a></td>
    <td>
      <a href="https://godoc.org/gopkg.in/h2non/gentleman.v2-mock">
        <img src="https://godoc.org/gopkg.in/h2non/gentleman.v2-mock?status.svg" />
      </a>
    </td>
    <td><a href="https://travis-ci.org/h2non/gentleman-mock"><img src="https://travis-ci.org/h2non/gentleman-mock.png" /></a></td>
    <td>Easy HTTP mocking using <a href="https://github.com/h2non/gock">gock</a></td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman-consul">consul</a></td>
    <td>
      <a href="https://godoc.org/gopkg.in/h2non/gentleman.v2-consul">
        <img src="https://godoc.org/gopkg.in/h2non/gentleman.v2-consul?status.svg" />
      </a>
    </td>
    <td><a href="https://travis-ci.org/h2non/gentleman-consul"><img src="https://travis-ci.org/h2non/gentleman-consul.png" /></a></td>
    <td><a href="https://www.consul.io">Consul</a> based server discovery with configurable retry/backoff policy</td>
  </tr>
</table>

### Community plugins

<table>
  <tr>
    <th>Name</th>
    <th>Docs</th>
    <th>Status</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><a href="https://github.com/izumin5210/gentleman-logger">logger</a></td>
    <td>
      <a href="https://godoc.org/github.com/izumin5210/gentleman-logger">
        <img src="https://godoc.org/github.com/izumin5210/gentleman-logger?status.svg" />
      </a>
    </td>
    <td><a href="https://travis-ci.org/izumin5210/gentleman-logger"><img src="https://travis-ci.org/izumin5210/gentleman-logger.png" /></a></td>
    <td>Easily log requests and responses</td>
  </tr>
</table>

[Send](https://github.com/h2non/gentleman/pull/new/master) a PR to add your plugin to the list.

### Creating plugins

You can create your own plugins for a wide variety of purposes, such as server discovery, custom HTTP tranport, modify any request/response param, intercept traffic, authentication and so on.

Plugins are essentially a set of middleware function handlers for one or multiple HTTP life cycle phases exposing [a concrete interface](https://github.com/h2non/gentleman/blob/755d55eef0bd26ae6b4ee19fe59001db2c46a51b/plugin/plugin.go#L16-L35) consumed by gentleman middleware layer.

For more details about plugins see the [plugin](https://github.com/h2non/gentleman/tree/master/plugin) package and [examples](https://github.com/h2non/gentleman/tree/master/_examples/plugin).

Also you can take a look to a plugin [implementation example](https://github.com/h2non/gentleman/blob/master/_examples/plugin/plugin.go).

## HTTP entities

`gentleman` provides two HTTP high level entities: `Client` and `Request`.

Each of these entities provides a common API and are both middleware capable, giving you the ability to plug in custom components with own logic into any of them.

`gentleman` was designed to provide strong reusability capabilities.
This is mostly achieved via its built-in hierarchical, inheritance-based middleware layer.

The following list describes how inheritance hierarchy works and is used across gentleman's entities.

- `Client` entity can inherit from other `Client` entity.
- `Request` entity can inherit from a `Client` entity.
- `Client` entity is mostly designed for reusability.
- `Client` entity can create multiple `Request` entities who implicitly inherits from `Client` entity itself.
- `Request` entity is designed to have specific HTTP request logic that is not typically reused.
- Both `Client` and `Request` entities are full middleware capable interfaces.
- Both `Client` and  `Request` entities can be cloned in order to produce a copy but side-effects free new entity.

You can see an inheritance usage example [here](https://github.com/h2non/gentleman/blob/master/_examples/inheritance/inheritance.go).

## Middleware

gentleman is completely based on a hierarchical middleware layer based on plugins that executes one or multiple function handlers (aka plugin interface) providing a simple way to plug in intermediate custom logic in your HTTP client.

It supports multiple phases which represents the full HTTP request/response life cycle, giving you the ability to perform actions before and after an HTTP transaction happen, even intercepting and stopping it.

The middleware stack chain is executed in FIFO order designed for single thread model.
Plugins can support goroutines, but plugins implementors should prevent data race issues due to concurrency in multithreading programming.

For more implementation details about the middleware layer, see the [middleware](https://github.com/h2non/gentleman/tree/master/middleware) package and [examples](https://github.com/h2non/gentleman/tree/master/_examples/middleware).

#### Middleware phases

Supported middleware phases triggered by gentleman HTTP dispatcher:

- **request** - Executed before a request is sent over the network.
- **response** - Executed when the client receives the response, even if it failed.
- **error** - Executed in case that an error ocurrs, support both injected or native error.
- **stop** - Executed in case that the request has been manually stopped via middleware (e.g: after interception).
- **intercept** - Executed in case that the request has been intercepted before network dialing.
- **before dial** - Executed before a request is sent over the network.
- **after dial** - Executed after the request dialing was done and the response has been received.

Note that the middleware layer has been designed for easy extensibility, therefore new phases may be added in the future and/or the developer could be able to trigger custom middleware phases if needed.

Feel free to fill an issue to discuss this capabilities in detail.

## API

See [godoc reference](https://godoc.org/gopkg.in/h2non/gentleman.v2) for detailed API documentation.

#### Subpackages

- [plugin](https://github.com/h2non/gentleman/tree/master/plugin) - [godoc](https://godoc.org/gopkg.in/h2non/gentleman.v2/plugin) - Plugin layer for gentleman.
- [mux](https://github.com/h2non/gentleman/tree/master/mux) - [godoc](https://godoc.org/gopkg.in/h2non/gentleman.v2/mux) - HTTP client multiplexer with built-in matchers.
- [middleware](https://github.com/h2non/gentleman/tree/master/middleware) - [godoc](https://godoc.org/gopkg.in/h2non/gentleman.v2/middleware) - Middleware layer used by gentleman.
- [context](https://github.com/h2non/gentleman/tree/master/context) - [godoc](https://godoc.org/gopkg.in/h2non/gentleman.v2/context) - HTTP context implementation for gentleman's middleware.
- [utils](https://github.com/h2non/gentleman/tree/master/utils) - [godoc](https://godoc.org/gopkg.in/h2non/gentleman.v2/utils) - HTTP utilities internally used.

## Examples

See [examples](https://github.com/h2non/gentleman/blob/master/_examples) directory for featured examples.

#### Simple request

```go
package main

import (
  "fmt"

  "gopkg.in/h2non/gentleman.v2"
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

  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman.v2/plugins/body"
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

#### Composition via multiplexer

```go
package main

import (
  "fmt"

  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman.v2/mux"
  "gopkg.in/h2non/gentleman.v2/plugins/url"
)

func main() {
  // Create a new client
  cli := gentleman.New()

  // Define the server url (must be first)
  cli.Use(url.URL("http://httpbin.org"))

  // Create a new multiplexer based on multiple matchers
  mx := mux.If(mux.Method("GET"), mux.Host("httpbin.org"))

  // Attach a custom plugin on the multiplexer that will be executed if the matchers passes
  mx.Use(url.Path("/headers"))

  // Attach the multiplexer on the main client
  cli.Use(mx)

  // Perform the request
  res, err := cli.Request().Send()
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
