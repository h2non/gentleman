// Package utils provides a set of reusable HTTP client utilities used internally
// in gentleman for required functionality and testing.
package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

// XMLCharDecoder is a helper type that takes a stream of bytes (not encoded in
// UTF-8) and returns a reader that encodes the bytes into UTF-8. This is done
// because Go's XML library only supports XML encoded in UTF-8.
type XMLCharDecoder func(charset string, input io.Reader) (io.Reader, error)

// ReplyWithStatus helper to write the http.Response status code and text.
func ReplyWithStatus(res *http.Response, code int) {
	res.StatusCode = code
	res.Status = strconv.Itoa(code) + " " + http.StatusText(code)
}

// WriteBodyString writes a string based body in a given http.Response.
func WriteBodyString(res *http.Response, body string) {
	res.Body = StringReader(body)
	res.ContentLength = int64(len(body))
}

// StringReader creates an io.ReadCloser interface from a string.
func StringReader(body string) io.ReadCloser {
	b := bytes.NewReader([]byte(body))

	rc, ok := io.Reader(b).(io.ReadCloser)
	if !ok && b != nil {
		rc = ioutil.NopCloser(b)
	}

	return rc
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

// NopCloser returns a ReadCloser with a no-op Close
//  method wrapping the provided Reader r.
func NopCloser() io.ReadCloser {
	return nopCloser{bytes.NewBuffer([]byte{})}
}
