package gentleman

// Originally based on grequests: https://github.com/levigross/grequests
// Apache License Version 2.0

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"os"

	"gopkg.in/h2non/gentleman.v2/context"
	"gopkg.in/h2non/gentleman.v2/utils"
)

// Response provides a more convenient and higher level Response struct.
// Implements an io.ReadCloser interface.
type Response struct {
	// Ok is a boolean flag that validates that the server returned a 2xx code.
	Ok bool

	// This is the Go error flag – if something went wrong within the request, this flag will be set.
	Error error

	// Sugar to check if the response status code is a client error (4xx).
	ClientError bool

	// Sugar to check if the response status code is a server error (5xx).
	ServerError bool

	// StatusCode is the HTTP Status Code returned by the HTTP Response. Taken from resp.StatusCode.
	StatusCode int

	// Header stores the response headers as http.Header interface.
	Header http.Header

	// Cookies stores the parsed response cookies.
	Cookies []*http.Cookie

	// Expose the native Go http.Response object for convenience.
	RawResponse *http.Response

	// Expose the native Go http.Request object for convenience.
	RawRequest *http.Request

	// Expose original request Context for convenience.
	Context *context.Context

	// Internal buffer store
	buffer *bytes.Buffer
}

func buildResponse(ctx *context.Context) (*Response, error) {
	resp := ctx.Response
	statusRange := int(resp.StatusCode / 100)

	res := &Response{
		// If your code is within the 2xx range – the response is considered `Ok`
		Ok:          statusRange >= 2 && statusRange <= 3,
		Error:       ctx.Error,
		ClientError: statusRange == 4,
		ServerError: statusRange == 5,
		Context:     ctx,
		RawResponse: resp,
		RawRequest:  ctx.Request,
		StatusCode:  resp.StatusCode,
		Header:      resp.Header,
		Cookies:     resp.Cookies(),
		buffer:      bytes.NewBuffer([]byte{}),
	}

	return res, res.Error
}

// Read is part of our ability to support io.ReadCloser
// if someone wants to make use of the raw body.
func (r *Response) Read(p []byte) (n int, err error) {
	if r.Error != nil {
		return -1, r.Error
	}
	return r.RawResponse.Body.Read(p)
}

// Close is part of our ability to support io.ReadCloser if
// someone wants to make use of the raw body.
func (r *Response) Close() error {
	if r.Error != nil {
		return r.Error
	}
	return r.RawResponse.Body.Close()
}

// SaveToFile allows you to download the contents
// of the response to a file.
func (r *Response) SaveToFile(fileName string) error {
	if r.Error != nil {
		return r.Error
	}

	fd, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer r.Close() // This is a noop if we use the internal ByteBuffer
	defer fd.Close()

	_, err = io.Copy(fd, r.getInternalReader())
	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

// JSON is a method that will populate a struct that is provided `userStruct`
// with the JSON returned within the response body.
func (r *Response) JSON(userStruct interface{}) error {
	if r.Error != nil {
		return r.Error
	}

	jsonDecoder := json.NewDecoder(r.getInternalReader())
	defer r.Close()

	err := jsonDecoder.Decode(&userStruct)
	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

// XML is a method that will populate a struct that is provided
// `userStruct` with the XML returned within the response body.
func (r *Response) XML(userStruct interface{}, charsetReader utils.XMLCharDecoder) error {
	if r.Error != nil {
		return r.Error
	}

	xmlDecoder := xml.NewDecoder(r.getInternalReader())
	if charsetReader != nil {
		xmlDecoder.CharsetReader = charsetReader
	}

	defer r.Close()
	if err := xmlDecoder.Decode(&userStruct); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Bytes returns the response as a byte array.
func (r *Response) Bytes() []byte {
	if r.Error != nil {
		return nil
	}

	r.populateResponseByteBuffer()

	// Are we still empty?
	if r.buffer.Len() == 0 {
		return nil
	}
	return r.buffer.Bytes()
}

// String returns the response as a string.
func (r *Response) String() string {
	if r.Error != nil {
		return ""
	}

	r.populateResponseByteBuffer()
	return r.buffer.String()
}

// ClearInternalBuffer is a function that will clear the internal buffer that we
// use to hold the .String() and .Bytes() data.
// Once you have used these functions you may want to free up the memory.
func (r *Response) ClearInternalBuffer() {
	if r.Error != nil {
		return // This is a noop as we will be dereferencing a null pointer
	}
	r.buffer.Reset()
}

// createResponseBytesBuffer is a utility method that will populate
// the internal byte reader – this is largely used for .String() and .Bytes()
func (r *Response) populateResponseByteBuffer() {
	// Have I done this already?
	if r.buffer.Len() != 0 {
		return
	}
	defer r.Close()

	// Is there any content?
	if r.RawResponse.ContentLength == 0 {
		return
	}

	// Did the server tell us how big the response is going to be?
	if r.RawResponse.ContentLength > 0 {
		r.buffer.Grow(int(r.RawResponse.ContentLength))
	}

	_, err := io.Copy(r.buffer, r)
	if err != nil && err != io.EOF {
		r.Error = err
		r.RawResponse.Body.Close()
	}
}

// getInternalReader because we implement io.ReadCloser and
// optionally hold a large buffer of the response (created by
// the user's request).
func (r *Response) getInternalReader() io.Reader {
	if r.buffer.Len() != 0 {
		return r.buffer
	}
	return r
}

// isChunkedResponse iterates over the response's transfer encodings
// and returns either true whether 'chunked' is found, or false, otherwise.
func isChunkedResponse(res *http.Response) bool {
	for _, te := range res.TransferEncoding {
		if te == "chunked" {
			return true
		}
	}
	return false
}
