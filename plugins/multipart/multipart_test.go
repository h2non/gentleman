package multipart

import (
	"bytes"
	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v2/context"
	"io/ioutil"
	"strings"
	"testing"
)

func TestFile(t *testing.T) {
	ctx := context.New()
	fn := newHandler()
	reader := bytes.NewReader([]byte("hello world"))

	File("foo", reader).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	st.Expect(t, ctx.Request.Method, "POST")
	st.Expect(t, match(body, "Content-Type: application/octet-stream"), true)
	st.Expect(t, match(body, `Content-Disposition: form-data; name="foo"; filename="foo"`), true)
	st.Expect(t, match(body, "hello world"), true)
}

func TestFiles(t *testing.T) {
	ctx := context.New()
	fn := newHandler()
	reader1 := bytes.NewReader([]byte("content1"))
	reader2 := bytes.NewReader([]byte("content2"))
	file1 := FormFile{"file1", reader1}
	file2 := FormFile{"file2", reader2}

	Files([]FormFile{file1, file2}).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	st.Expect(t, ctx.Request.Method, "POST")
	st.Expect(t, match(body, "Content-Type: application/octet-stream"), true)
	st.Expect(t, match(body, `Content-Disposition: form-data; name="file1"; filename="file1"`), true)
	st.Expect(t, match(body, `Content-Disposition: form-data; name="file2"; filename="file2"`), true)
	st.Expect(t, match(body, "content1"), true)
	st.Expect(t, match(body, "content2"), true)
}

func TestFields(t *testing.T) {
	ctx := context.New()
	fn := newHandler()
	fields := map[string]Values{"foo": {"data=bar"}, "bar": {"data=baz"}}

	Fields(fields).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	st.Expect(t, ctx.Request.Method, "POST")
	st.Expect(t, match(body, `Content-Disposition: form-data; name="foo"`), true)
	st.Expect(t, match(body, `Content-Disposition: form-data; name="bar"`), true)
	st.Expect(t, match(body, "data=bar"), true)
	st.Expect(t, match(body, "data=baz"), true)
}

func TestData(t *testing.T) {
	ctx := context.New()
	fn := newHandler()
	reader := bytes.NewReader([]byte("hello world"))
	fields := map[string]Values{"foo": {"data=bar"}, "bar": {"data=baz"}}
	data := FormData{
		Files: []FormFile{{Name: "foo", Reader: reader}},
		Data:  fields,
	}

	Data(data).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	st.Expect(t, ctx.Request.Method, "POST")
	st.Expect(t, match(body, `Content-Disposition: form-data; name="foo"`), true)
	st.Expect(t, match(body, `Content-Disposition: form-data; name="bar"`), true)
	st.Expect(t, match(body, "data=bar"), true)
	st.Expect(t, match(body, "data=baz"), true)
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

func match(body []byte, str string) bool {
	return strings.Contains(string(body), str)
}
