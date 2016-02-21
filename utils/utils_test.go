package utils

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestReplyWithStatus(t *testing.T) {
	res := &http.Response{}
	ReplyWithStatus(res, 200)
	if res.StatusCode != 200 || res.Status != "200 OK" {
		t.Fatalf("Invalid response status: %s", res.Status)
	}

	res = &http.Response{}
	ReplyWithStatus(res, 400)
	if res.StatusCode != 400 || res.Status != "400 Bad Request" {
		t.Fatalf("Invalid response status: %d", res.StatusCode)
	}
}

func TestWriteBodyString(t *testing.T) {
	res := &http.Response{}
	body := "hello world"
	WriteBodyString(res, body)

	if res.ContentLength != int64(len(body)) {
		t.Fatalf("Invalid content length: %d", res.ContentLength)
	}

	contents, _ := ioutil.ReadAll(res.Body)
	if string(contents) != body {
		t.Fatal("Invalid body data")
	}
}
