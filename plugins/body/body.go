package body

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"strings"

	c "gopkg.in/h2non/gentleman.v2/context"
	p "gopkg.in/h2non/gentleman.v2/plugin"
	"gopkg.in/h2non/gentleman.v2/utils"
)

// String defines the HTTP request body based on the given string.
func String(data string) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		ctx.Request.Method = getMethod(ctx)
		ctx.Request.Body = utils.StringReader(data)
		ctx.Request.ContentLength = int64(bytes.NewBufferString(data).Len())
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
		ctx.Request.Body = ioutil.NopCloser(buf)
		ctx.Request.ContentLength = int64(buf.Len())
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
		ctx.Request.Body = ioutil.NopCloser(buf)
		ctx.Request.ContentLength = int64(buf.Len())
		ctx.Request.Header.Set("Content-Type", "application/xml")

		h.Next(ctx)
	})
}

// Reader defines a io.Reader stream as request body.
// Content-Type header won't be defined automatically, you have to declare it manually.
func Reader(body io.Reader) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		rc, ok := body.(io.ReadCloser)
		if !ok && body != nil {
			rc = ioutil.NopCloser(body)
		}

		req := ctx.Request
		if body != nil {
			switch v := body.(type) {
			case *bytes.Buffer:
				req.ContentLength = int64(v.Len())
			case *bytes.Reader:
				req.ContentLength = int64(v.Len())
			case *strings.Reader:
				req.ContentLength = int64(v.Len())
			}
		}

		req.Body = rc
		ctx.Request.Method = getMethod(ctx)

		h.Next(ctx)
	})
}

func getMethod(ctx *c.Context) string {
	method := ctx.Request.Method
	if method == "" {
		return "POST"
	}
	return method
}
