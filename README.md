# gentleman [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GitHub release](https://img.shields.io/github/tag/h2non/gentleman.svg)](https://github.com/h2non/gentleman/releases) [![GoDoc](https://godoc.org/github.com/h2non/gentleman?status.svg)](https://godoc.org/github.com/h2non/gentleman) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman)](https://goreportcard.com/report/github.com/h2non/gentleman)

<img src="http://s10.postimg.org/5e31ox1ft/gentleman.png" align="right" height="200" />

Plugin-driven library to create rich, versatile and composable HTTP clients in [Go](http://golang.org).

**Note**: work in progress, interface contract may change at this time.

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v0
```

## Goals

- Plugin driven, middleware based.
- Simple and expressive API.
- Idiomatic built on top of Go `net/http` package.
- Strong extensibility capatabilities.
- Multiplexer based on composition.
- Easy to configure and use.
- Convenient helpers and abstractions.
- Dependency free.

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
    <td><img src="https://img.shields.io/badge/api-stable-green.svg?style=flat" /></td>
    <td><img src="https://travis-ci.org/h2non/gentleman.png" /></td>
    <td>Easily declare URL, base URL and path values in HTTP requests</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/auth">auth</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/auth">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/api-stable-green.svg?style=flat" /></td>
    <td><img src="https://travis-ci.org/h2non/gentleman.png" /></td> 
    <td>Declare authorization headers in your requests</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/body">body</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/body">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/api-stable-green.svg?style=flat" /></td>
    <td><img src="https://travis-ci.org/h2non/gentleman.png" /></td> 
    <td>Easily define bodies based on JSON, XML, strings, buffers or streams</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/bodytype">bodytype</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/bodytype">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/api-stable-green.svg?style=flat" /></td>
    <td><img src="https://travis-ci.org/h2non/gentleman.png" /></td> 
    <td>Define body MIME type by alias</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/cookies">cookies</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/cookies">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/api-stable-green.svg?style=flat" /></td>
    <td><img src="https://travis-ci.org/h2non/gentleman.png" /></td> 
    <td>Declare and store HTTP cookies easily</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/compression">compression</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/compression">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/api-stable-green.svg?style=flat" /></td>
    <td><img src="https://travis-ci.org/h2non/gentleman.png" /></td> 
    <td>Helpers to define enable/disable HTTP compression</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/headers">headers</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/headers">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/api-stable-green.svg?style=flat" /></td>
    <td><img src="https://travis-ci.org/h2non/gentleman.png" /></td> 
    <td>Manage HTTP headers easily</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/proxy">proxy</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/proxy">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/api-stable-green.svg?style=flat" /></td>
    <td><img src="https://travis-ci.org/h2non/gentleman.png" /></td> 
    <td>Configure HTTP proxy servers</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/query">query</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/query">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/api-stable-green.svg?style=flat" /></td>
    <td><img src="https://travis-ci.org/h2non/gentleman.png" /></td> 
    <td>Easily manage query params</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/redirect">redirect</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/redirect">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/api-stable-green.svg?style=flat" /></td>
    <td><img src="https://travis-ci.org/h2non/gentleman.png" /></td> 
    <td>Easily configure a custom redirect policy</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/timeout">timeout</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/timeout">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/api-stable-green.svg?style=flat" /></td>
    <td><img src="https://travis-ci.org/h2non/gentleman.png" /></td> 
    <td>Easily configure the HTTP timeouts (request, dial, TLS...)</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/transport">transport</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/transport">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/api-stable-green.svg?style=flat" /></td>
    <td><img src="https://travis-ci.org/h2non/gentleman.png" /></td> 
    <td>Define a custom HTTP transport easily</td>
  </tr>
  <tr>
    <td><a href="https://github.com/h2non/gentleman/tree/master/plugins/tls">tls</a></td>
    <td>
      <a href="https://godoc.org/github.com/h2non/gentleman/plugins/tls">
        <img src="https://godoc.org/github.com/h2non/gentleman?status.svg" />
      </a>
    </td>
    <td><img src="https://img.shields.io/badge/api-stable-green.svg?style=flat" /></td>
    <td><img src="https://travis-ci.org/h2non/gentleman.png" /></td> 
    <td>Configure the TLS options used by the HTTP transport</td>
  </tr>
</table>

[Send](https://github.com/h2non/gentleman/pull/new/master) a PR to add your plugin to the list.

### Creating plugins

`TODO`

## API

See [godoc reference](https://godoc.org/github.com/h2non/gentleman) for detailed API documentation.

## Examples

See [examples](https://github.com/h2non/gentleman/blob/master/_examples) directory for featured use case examples.

<!--
```go
import (
  "fmt"
  "net/http"
  "gopkg.in/h2non/gentleman.v0"
  "gopkg.in/h2non/gentleman.v0/plugins"
)

client := gentleman.New()

client.UseRequest(plugins.Cookies())
client.UseRequest(plugins.Timeout(3000))
client.UseRequest(plugins.Authenticate("foo:s3cr3t"))
client.UseRequest(plugins.ContentType("application/json"))

client.UseResponse(plugins.BodyParser())
client.UseResponse(plugins.ErrorHandler())
client.UseResponse(plugins.Retry(3)))

// Custom middleware
client.UseResponse(func (ctx *gentleman.Context, h gentleman.Handler) {
  // ...
})

res, ctx, err := client.Get("/hello")
if res.StatusCode != 200 {
  fmt.Printf("Error: %s", res.StatusCode)
}
```
-->

## License 

MIT - Tomas Aparicio
