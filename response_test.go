package gentleman

import (
	"testing"
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
		{300, false, false, false},
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
			t.Errorf("Invalid status code: %s", res.StatusCode)
		}
		if res.Ok != test.ok {
			t.Error("Invalid Ok field")
		}
		if res.ClientError != test.client {
			t.Error("Invalid ClientError field")
		}
		if res.ServerError != test.server {
			t.Error("Invalid ServerError field")
		}
	}
}
