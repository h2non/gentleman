# gentleman/context [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman) [![GoDoc](https://godoc.org/github.com/h2non/gentleman/context?status.svg)](https://godoc.org/github.com/h2non/gentleman/context) [![API](https://img.shields.io/badge/status-stable-green.svg?style=flat)](https://godoc.org/github.com/h2non/gentleman/context) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman/context)](https://goreportcard.com/report/github.com/h2non/gentleman/context)

Package `context` implements a simple request-aware HTTP context used by plugins and exposed by the middleware layer, designed to share polymorfic data types across plugins in the middleware call chain.

`context` is not thread-safe by default.

## Installation

```bash
go get -u gopkg.in/h2non/gentleman.v1/context
```

## API

See [godoc](https://godoc.org/github.com/h2non/gentleman/context) reference.

## License

MIT - Tomas Aparicio
