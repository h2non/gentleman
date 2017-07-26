# gentleman/context [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GoDoc](https://godoc.org/github.com/h2non/gentleman/context?status.svg)](https://godoc.org/github.com/h2non/gentleman/context) [![API](https://img.shields.io/badge/status-stable-green.svg?style=flat)](https://godoc.org/github.com/h2non/gentleman/context) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman/context)](https://goreportcard.com/report/github.com/h2non/gentleman/context)

Package `context` implements a simple request-aware HTTP context used by plugins and exposed by the middleware layer, designed to share
polymorphic data across plugins in the middleware call chain.

It is built on top of the standard built-in [context](https://golang.org/pkg/context) package in Go.

gentleman's `context` also implements a valid stdlid `context.Context` interface:
https://golang.org/pkg/context/#Context

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v2/context
```

## API

See [godoc](https://godoc.org/github.com/h2non/gentleman/context) reference.

## License

MIT - Tomas Aparicio
