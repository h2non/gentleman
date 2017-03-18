package body

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v2/context"
)

func TestBodyJSONEncodeMap(t *testing.T) {
	ctx := context.New()
	fn := newHandler()

	json := map[string]string{"foo": "bar"}
	JSON(json).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	buf, err := ioutil.ReadAll(ctx.Request.Body)
	st.Expect(t, err, nil)
	st.Expect(t, ctx.Request.Method, "GET")
	st.Expect(t, ctx.Request.Header.Get("Content-Type"), "application/json")
	st.Expect(t, int(ctx.Request.ContentLength), 14)
	st.Expect(t, string(buf[0:len(buf)-1]), `{"foo":"bar"}`)
}

func TestBodyJSONEncodeString(t *testing.T) {
	ctx := context.New()
	fn := newHandler()

	json := `{"foo":"bar"}`
	JSON(json).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	buf, err := ioutil.ReadAll(ctx.Request.Body)
	st.Expect(t, err, nil)
	st.Expect(t, ctx.Request.Method, "GET")
	st.Expect(t, ctx.Request.Header.Get("Content-Type"), "application/json")
	st.Expect(t, int(ctx.Request.ContentLength), 13)
	st.Expect(t, string(buf), `{"foo":"bar"}`)
}

func TestBodyJSONEncodeBytes(t *testing.T) {
	ctx := context.New()
	fn := newHandler()

	json := []byte(`{"foo":"bar"}`)
	JSON(json).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	buf, err := ioutil.ReadAll(ctx.Request.Body)
	st.Expect(t, err, nil)
	st.Expect(t, ctx.Request.Method, "GET")
	st.Expect(t, ctx.Request.Header.Get("Content-Type"), "application/json")
	st.Expect(t, int(ctx.Request.ContentLength), 13)
	st.Expect(t, string(buf), `{"foo":"bar"}`)
}

func TestBodyXMLEncodeStruct(t *testing.T) {
	ctx := context.New()
	fn := newHandler()

	type xmlTest struct {
		Name string `xml:"name>first"`
	}
	xml := xmlTest{Name: "foo"}
	XML(xml).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)

	buf, err := ioutil.ReadAll(ctx.Request.Body)
	st.Expect(t, err, nil)
	st.Expect(t, ctx.Request.Method, "GET")
	st.Expect(t, ctx.Request.Header.Get("Content-Type"), "application/xml")
	st.Expect(t, int(ctx.Request.ContentLength), 50)
	st.Expect(t, string(buf), `<xmlTest><name><first>foo</first></name></xmlTest>`)
}

func TestBodyXMLEncodeString(t *testing.T) {
	ctx := context.New()
	fn := newHandler()

	xml := "<test>foo</test>"
	XML(xml).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)

	buf, err := ioutil.ReadAll(ctx.Request.Body)
	st.Expect(t, err, nil)
	st.Expect(t, ctx.Request.Method, "GET")
	st.Expect(t, ctx.Request.Header.Get("Content-Type"), "application/xml")
	st.Expect(t, int(ctx.Request.ContentLength), 16)
	st.Expect(t, string(buf), `<test>foo</test>`)
}

func TestBodyXMLEncodeBytes(t *testing.T) {
	ctx := context.New()
	ctx.Request.Method = ""
	fn := newHandler()

	xml := []byte("<test>foo</test>")
	XML(xml).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)

	buf, err := ioutil.ReadAll(ctx.Request.Body)
	st.Expect(t, err, nil)
	st.Expect(t, ctx.Request.Method, "POST")
	st.Expect(t, ctx.Request.Header.Get("Content-Type"), "application/xml")
	st.Expect(t, int(ctx.Request.ContentLength), 16)
	st.Expect(t, string(buf), `<test>foo</test>`)
}

func TestBodyReader(t *testing.T) {
	ctx := context.New()
	ctx.Request.Method = "POST"
	fn := newHandler()

	reader := bytes.NewReader([]byte("foo bar"))
	Reader(reader).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)

	buf, err := ioutil.ReadAll(ctx.Request.Body)
	st.Expect(t, err, nil)
	st.Expect(t, ctx.Request.Method, "POST")
	st.Expect(t, ctx.Request.Header.Get("Content-Type"), "")
	st.Expect(t, int(ctx.Request.ContentLength), 7)
	st.Expect(t, string(buf), "foo bar")
}

func TestBodyReaderContextDataSharing(t *testing.T) {
	ctx := context.New()
	ctx.Request.Method = "POST"
	fn := newHandler()

	// Set sample context data
	ctx.Set("foo", "bar")
	ctx.Set("bar", "baz")

	reader := bytes.NewReader([]byte("foo bar"))
	Reader(reader).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)

	buf, err := ioutil.ReadAll(ctx.Request.Body)
	st.Expect(t, err, nil)
	st.Expect(t, ctx.Request.Method, "POST")
	st.Expect(t, ctx.Request.Header.Get("Content-Type"), "")
	st.Expect(t, int(ctx.Request.ContentLength), 7)
	st.Expect(t, string(buf), "foo bar")

	// Test context data
	st.Expect(t, ctx.GetString("foo"), "bar")
	st.Expect(t, ctx.GetString("bar"), "baz")
}

type handler struct {
	fn     context.Handler
	called bool
}

func newHandler() *handler {
	h := &handler{}
	h.fn = context.NewHandler(func(c *context.Context) {
		h.called = true
	})
	return h
}
