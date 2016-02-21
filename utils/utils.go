package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
)

// ReplyWithStatus helper to write the http.Response status code and text
func ReplyWithStatus(res *http.Response, code int) {
	res.StatusCode = code
	res.Status = strconv.Itoa(code) + " " + http.StatusText(code)
}

// WriteBodyString writes a string based body in a given http.Response
func WriteBodyString(res *http.Response, body string) {
	res.Body = StringReader(body)
	res.ContentLength = int64(len(body))
}

// StringReader creates an io.ReadCloser interface from a string
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

// EnsureTransporterFinalized will ensure that when the HTTP client is GCed
// the runtime will close the idle connections (so that they won't leak)
// this function was adopted from Hashicorp's go-cleanhttp package
func EnsureTransporterFinalized(httpTransport *http.Transport) {
	runtime.SetFinalizer(&httpTransport, func(transportInt **http.Transport) {
		(*transportInt).CloseIdleConnections()
	})
}

// SetTransportFinalizer sets a finalizer on the transport to ensure that
// idle connections are closed prior to garbage collection; otherwise
// these may leak.
func SetTransportFinalizer(transport *http.Transport) {
	runtime.SetFinalizer(&transport, func(t **http.Transport) {
		(*t).CloseIdleConnections()
	})
}
