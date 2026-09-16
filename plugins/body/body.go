package body

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	c "gopkg.in/h2non/gentleman.v2/context"
	p "gopkg.in/h2non/gentleman.v2/plugin"
	"gopkg.in/h2non/gentleman.v2/utils"
)

// String defines the HTTP request body based on the given string.
func String(data string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		ctx.Request.Method = getMethod(ctx)
		setReplayableBody(ctx.Request, []byte(data), func() io.ReadCloser {
			return utils.StringReader(data)
		})
		h.Next(ctx)
	})
}

// JSON defines a JSON body in the outgoing request.
// Supports strings, array of bytes or buffer.
func JSON(data interface{}) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		buf := &bytes.Buffer{}

		switch data.(type) {
		case string:
			buf.WriteString(data.(string))
		case []byte:
			buf.Write(data.([]byte))
		default:
			if err := json.NewEncoder(buf).Encode(data); err != nil {
				h.Error(ctx, err)
				return
			}
		}

		ctx.Request.Method = getMethod(ctx)
		setReplayableBody(ctx.Request, buf.Bytes(), nil)
		ctx.Request.Header.Set("Content-Type", "application/json")

		h.Next(ctx)
	})
}

// XML defines a XML body in the outgoing request.
// Supports strings, array of bytes or buffer.
func XML(data interface{}) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		buf := &bytes.Buffer{}

		switch data.(type) {
		case string:
			buf.WriteString(data.(string))
		case []byte:
			buf.Write(data.([]byte))
		default:
			if err := xml.NewEncoder(buf).Encode(data); err != nil {
				h.Error(ctx, err)
				return
			}
		}

		ctx.Request.Method = getMethod(ctx)
		setReplayableBody(ctx.Request, buf.Bytes(), nil)
		ctx.Request.Header.Set("Content-Type", "application/xml")

		h.Next(ctx)
	})
}

// Reader defines a io.Reader stream as request body.
// Content-Type header won't be defined automatically, you have to declare it manually.
func Reader(body io.Reader) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		req := ctx.Request
		if body != nil {
			switch v := body.(type) {
			case *bytes.Buffer:
				setReplayableBody(req, v.Bytes(), nil)
			case *bytes.Reader:
				snapshot := *v
				req.ContentLength = int64(v.Len())
				req.GetBody = func() (io.ReadCloser, error) {
					r := snapshot
					return ioutil.NopCloser(&r), nil
				}
				req.Body = ioutil.NopCloser(v)
			case *strings.Reader:
				snapshot := *v
				req.ContentLength = int64(v.Len())
				req.GetBody = func() (io.ReadCloser, error) {
					r := snapshot
					return ioutil.NopCloser(&r), nil
				}
				req.Body = ioutil.NopCloser(v)
			default:
				rc, ok := body.(io.ReadCloser)
				if !ok {
					rc = ioutil.NopCloser(body)
				}
				req.Body = rc
			}
		}

		ctx.Request.Method = getMethod(ctx)

		h.Next(ctx)
	})
}

// setReplayableBody attaches a body that net/http can resend on 307/308 redirects.
func setReplayableBody(req *http.Request, data []byte, open func() io.ReadCloser) {
	if open == nil {
		open = func() io.ReadCloser {
			return ioutil.NopCloser(bytes.NewReader(data))
		}
	}
	req.Body = open()
	req.ContentLength = int64(len(data))
	req.GetBody = func() (io.ReadCloser, error) {
		return open(), nil
	}
}

func getMethod(ctx *c.Context) string {
	method := ctx.Request.Method
	if method == "" {
		return "POST"
	}
	return method
}
