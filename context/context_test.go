// Includes code from Gorilla context test:
// https://github.com/gorilla/context/blob/master/context_test.go
// © 2012 The Gorilla Authors
package context

import (
	"testing"

	"github.com/nbio/st"
)

type keyType int

const (
	key1 keyType = iota
	key2
)

func TestContext(t *testing.T) {
	ctx := New()
	crc := getContextReadCloser(ctx.Request)

	// Get()
	st.Expect(t, ctx.Get(key1), nil)

	// Set()
	ctx.Set(key1, "1")
	st.Expect(t, ctx.Get(key1), "1")
	st.Expect(t, len(crc.Context()), 1)

	ctx.Set(key2, "2")
	st.Expect(t, ctx.Get(key2), "2")
	st.Expect(t, len(crc.Context()), 2)

	// GetOk()
	value, ok := ctx.GetOk(key1)
	st.Expect(t, value, "1")
	st.Expect(t, ok, true)

	value, ok = ctx.GetOk("not exists")
	st.Expect(t, value, nil)
	st.Expect(t, ok, false)

	ctx.Set("nil value", nil)
	value, ok = ctx.GetOk("nil value")
	st.Expect(t, value, nil)
	st.Expect(t, ok, true)

	// GetString()
	ctx.Set("int value", 13)
	ctx.Set("string value", "hello")
	str := ctx.GetString("int value")
	st.Expect(t, str, "")
	str = ctx.GetString("string value")
	st.Expect(t, str, "hello")

	// GetAll()
	values := ctx.GetAll()
	st.Expect(t, len(values), 5)

	// Delete()
	ctx.Delete(key1)
	st.Expect(t, ctx.Get(key1), nil)
	st.Expect(t, len(crc.Context()), 4)

	ctx.Delete(key2)
	st.Expect(t, ctx.Get(key2), nil)
	st.Expect(t, len(crc.Context()), 3)

	// Clear()
	ctx.Set(key1, true)
	values = ctx.GetAll()
	ctx.Clear()
	st.Expect(t, len(crc.Context()), 0)
	val, _ := values["int value"].(int)
	st.Expect(t, val, 13) // Clear shouldn't delete values grabbed before
}

func TestContextInheritance(t *testing.T) {
	parent := New()
	ctx := New()
	ctx.UseParent(parent)

	parent.Set("foo", "bar")
	ctx.Set("bar", "foo")
	st.Expect(t, ctx.Get("foo"), "bar")
	st.Expect(t, ctx.Get("bar"), "foo")

	ctx.Set("foo", "foo")
	st.Expect(t, ctx.Get("foo"), "foo")
}

func TestContextRoot(t *testing.T) {
	root := New()
	parent := New()
	parent.UseParent(root)
	ctx := New()
	ctx.UseParent(parent)
	if ctx.Root() != root {
		t.Error("Invalid root context")
	}
}

func TestContextClone(t *testing.T) {
	ctx := New()
	ctx.Set("bar", "foo")
	st.Expect(t, ctx.Get("bar"), "foo")

	newCtx := ctx.Clone()
	st.Expect(t, newCtx.Get("bar"), "foo")

	// Overwrite value
	newCtx.Set("bar", "bar")
	st.Expect(t, ctx.Get("bar"), "foo")
	st.Expect(t, newCtx.Get("bar"), "bar")
}

func TestContextCopy(t *testing.T) {
	ctx := New()
	ctx.Set("bar", "foo")
	st.Expect(t, ctx.Get("bar"), "foo")

	newCtx := New()
	ctx.Copy(newCtx.Request)
	st.Expect(t, newCtx.Get("bar"), "foo")

	// Ensure inmutability
	newCtx.Set("bar", "bar")
	st.Expect(t, ctx.Get("bar"), "foo")
	st.Expect(t, newCtx.Get("bar"), "bar")
}

func TestContextSetRequest(t *testing.T) {
	ctx := New()
	ctx.Set("bar", "foo")
	st.Expect(t, ctx.Get("bar"), "foo")

	newCtx := New()
	ctx.Copy(newCtx.Request)
	st.Expect(t, newCtx.Get("bar"), "foo")

	// Ensure inmutability
	newCtx.Set("bar", "bar")
	st.Expect(t, ctx.Get("bar"), "foo")
	st.Expect(t, newCtx.Get("bar"), "bar")
}
