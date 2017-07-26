package gentleman

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v2/utils"
)

func TestResponseBuild(t *testing.T) {
	ctx := NewContext()
	ctx.Response.StatusCode = 200
	ctx.Response.Header.Set("foo", "bar")

	res, err := buildResponse(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if res.Header.Get("foo") != "bar" {
		t.Error("Invalid foo header")
	}
	if res.StatusCode != 200 {
		t.Errorf("Invalid status code: %d", res.StatusCode)
	}
	if !res.Ok {
		t.Error("Ok field must be true")
	}
	if res.ClientError || res.ServerError {
		t.Error("Invalid client/server status error")
	}
	if res.Context != ctx {
		t.Error("Invalid context")
	}
	if res.Header.Get("foo") != "bar" {
		t.Error("Invalid foo header")
	}
}

func TestResponseBuildStatusCodes(t *testing.T) {
	cases := []struct {
		code   int
		ok     bool
		client bool
		server bool
	}{
		{0, false, false, false},
		{200, true, false, false},
		{204, true, false, false},
		{300, true, false, false},
		{400, false, true, false},
		{404, false, true, false},
		{499, false, true, false},
		{500, false, false, true},
		{599, false, false, true},
	}

	for _, test := range cases {
		ctx := NewContext()
		ctx.Response.StatusCode = test.code
		res, _ := buildResponse(ctx)
		if res.StatusCode != test.code {
			t.Errorf("Invalid status code: %d", res.StatusCode)
		}
		if res.Ok != test.ok {
			t.Errorf("Invalid Ok field: %v", res.Ok)
		}
		if res.ClientError != test.client {
			t.Error("Invalid ClientError field")
		}
		if res.ServerError != test.server {
			t.Error("Invalid ServerError field")
		}
	}
}

func TestResponseReadError(t *testing.T) {
	ctx := NewContext()
	ctx.Error = errors.New("foo error")
	res, _ := buildResponse(ctx)
	num, err := res.Read([]byte{})
	st.Reject(t, err, nil)
	st.Expect(t, err.Error(), "foo error")
	st.Expect(t, num, -1)
}

func TestResponseCloseError(t *testing.T) {
	ctx := NewContext()
	ctx.Error = errors.New("foo error")
	res, _ := buildResponse(ctx)
	err := res.Close()
	st.Reject(t, err, nil)
	st.Expect(t, err.Error(), "foo error")
}

func TestResponseSaveToFile(t *testing.T) {
	ctx := NewContext()
	utils.WriteBodyString(ctx.Response, "hello world")
	res, _ := buildResponse(ctx)
	err := res.SaveToFile("body.tmp")
	st.Expect(t, err, nil)
	defer os.Remove("body.tmp")

	body, err := ioutil.ReadFile("body.tmp")
	st.Expect(t, err, nil)
	st.Expect(t, string(body), "hello world")
}

func TestResponseSaveToFileError(t *testing.T) {
	ctx := NewContext()
	ctx.Error = errors.New("foo error")
	res, _ := buildResponse(ctx)
	err := res.SaveToFile("body.tmp")
	st.Reject(t, err, nil)
}

func TestResponseJSON(t *testing.T) {
	type jsonData struct {
		Foo string `json:"foo"`
	}
	json := &jsonData{}
	ctx := NewContext()
	utils.WriteBodyString(ctx.Response, `{"foo":"bar"}`)
	res, _ := buildResponse(ctx)
	err := res.JSON(json)
	st.Expect(t, err, nil)
	st.Expect(t, json.Foo, "bar")
}

func TestResponseJSONError(t *testing.T) {
	type jsonData struct {
		Foo string `json:"foo"`
	}
	json := &jsonData{}
	ctx := NewContext()
	ctx.Error = errors.New("foo error")
	res, _ := buildResponse(ctx)
	err := res.JSON(json)
	st.Reject(t, err, nil)
	st.Expect(t, json.Foo, "")
}

func TestResponseXML(t *testing.T) {
	type xml struct {
		Foo string `xml:"foo"`
	}
	xmlData := &xml{}
	ctx := NewContext()
	utils.WriteBodyString(ctx.Response, `<xml><foo>bar</foo></xml>`)
	res, _ := buildResponse(ctx)
	err := res.XML(xmlData, nil)
	st.Expect(t, err, nil)
	st.Expect(t, xmlData.Foo, "bar")
}

func TestResponseXMLError(t *testing.T) {
	type xml struct {
		Foo string `xml:"foo"`
	}
	xmlData := &xml{}
	ctx := NewContext()
	ctx.Error = errors.New("foo error")
	utils.WriteBodyString(ctx.Response, `<xml><foo>bar</foo></xml>`)
	res, _ := buildResponse(ctx)
	err := res.XML(xmlData, nil)
	st.Reject(t, err, nil)
	st.Expect(t, xmlData.Foo, "")
}

func TestResponseString(t *testing.T) {
	ctx := NewContext()
	utils.WriteBodyString(ctx.Response, "foo bar")
	res, err := buildResponse(ctx)
	body := res.String()
	st.Expect(t, err, nil)
	st.Expect(t, body, "foo bar")
}

func TestResponseStringError(t *testing.T) {
	ctx := NewContext()
	ctx.Error = errors.New("foo error")
	utils.WriteBodyString(ctx.Response, "foo bar")
	res, err := buildResponse(ctx)
	body := res.String()
	st.Reject(t, err, nil)
	st.Expect(t, body, "")
}

func TestResponseBytes(t *testing.T) {
	ctx := NewContext()
	utils.WriteBodyString(ctx.Response, "foo bar")
	res, err := buildResponse(ctx)
	body := res.Bytes()
	st.Expect(t, err, nil)
	st.Expect(t, string(body), "foo bar")
}

func TestResponseBytesError(t *testing.T) {
	ctx := NewContext()
	ctx.Error = errors.New("foo error")
	utils.WriteBodyString(ctx.Response, "foo bar")
	res, err := buildResponse(ctx)
	body := res.Bytes()
	st.Reject(t, err, nil)
	st.Expect(t, string(body), "")
}

func TestResponseReaderBuffer(t *testing.T) {
	ctx := NewContext()
	utils.WriteBodyString(ctx.Response, "foo bar")
	res, err := buildResponse(ctx)
	res.RawResponse.ContentLength = 7

	body := res.Bytes()
	st.Expect(t, err, nil)
	st.Expect(t, string(body), "foo bar")
	st.Expect(t, res.buffer.String(), "foo bar")

	res.ClearInternalBuffer()
	st.Expect(t, string(res.Bytes()), "")
}

func TestResponseReaderEmtpyBuffer(t *testing.T) {
	ctx := NewContext()
	res, err := buildResponse(ctx)
	res.RawResponse.ContentLength = 0

	body := res.Bytes()
	st.Expect(t, err, nil)
	st.Expect(t, string(body), "")
	st.Expect(t, res.buffer.String(), "")

	res.ClearInternalBuffer()
	st.Expect(t, string(res.Bytes()), "")
}

func TestResponseReaderBufferError(t *testing.T) {
	ctx := NewContext()
	ctx.Error = errors.New("foo error")
	res, err := buildResponse(ctx)
	body := res.Bytes()
	st.Reject(t, err, nil)
	st.Expect(t, string(body), "")
	st.Expect(t, res.buffer.Len(), 0)
	res.ClearInternalBuffer()
	st.Expect(t, res.buffer.Len(), 0)
}
