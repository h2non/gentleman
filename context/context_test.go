// Includes code from Gorilla context test:
// https://github.com/gorilla/context/blob/master/context_test.go
// © 2012 The Gorilla Authors
package context

import (
	gocontext "context"
	"testing"
	"time"

	"github.com/nbio/st"
)

type keyType int

const (
	key1 keyType = iota
	key2
)

func TestContext(t *testing.T) {
	ctx := New()
	store := ctx.getStore()

	// Get()
	st.Expect(t, ctx.Get(key1), nil)

	// Set()
	ctx.Set(key1, "1")
	st.Expect(t, ctx.Get(key1), "1")
	st.Expect(t, len(store), 1)
	st.Expect(t, store[key1], "1")

	ctx.Set(key2, "2")
	st.Expect(t, ctx.Get(key2), "2")
	st.Expect(t, len(store), 2)

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
	st.Expect(t, len(store), 4)

	ctx.Delete(key2)
	st.Expect(t, ctx.Get(key2), nil)
	st.Expect(t, len(store), 3)

	// Clear()
	ctx.Set(key1, true)
	values = ctx.GetAll()
	ctx.Clear()
	st.Expect(t, len(store), 0)
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

func TestContextGetAll(t *testing.T) {
	parent := New()
	ctx := New()
	ctx.UseParent(parent)

	parent.Set("foo", "bar")
	ctx.Set("bar", "foo")
	st.Expect(t, ctx.Get("foo"), "bar")
	st.Expect(t, ctx.Get("bar"), "foo")

	store := ctx.GetAll()
	st.Expect(t, len(store), 2)
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

func TestContextGetters(t *testing.T) {
	parent := New()
	ctx := New()
	ctx.UseParent(parent)

	parent.Set("foo", "bar")
	ctx.Set("bar", "foo")
	st.Expect(t, ctx.GetString("foo"), "bar")
	st.Expect(t, ctx.GetString("bar"), "foo")
	ctx.Clear()

	parent.Set("foo", 1)
	ctx.Set("bar", 2)
	foo, ok := ctx.GetInt("foo")
	st.Expect(t, ok, true)
	st.Expect(t, foo, 1)
	bar, ok := ctx.GetInt("bar")
	st.Expect(t, ok, true)
	st.Expect(t, bar, 2)

	store := ctx.GetAll()
	st.Expect(t, len(store), 2)
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

func TestContextCopyRequestContext(t *testing.T) {
	tests := []struct {
		name    string
		prepare func() (gocontext.Context, gocontext.CancelFunc)
	}{
		{"active", func() (gocontext.Context, gocontext.CancelFunc) {
			return gocontext.WithCancel(gocontext.Background())
		}},
		{"canceled", func() (gocontext.Context, gocontext.CancelFunc) {
			ctx, cancel := gocontext.WithCancel(gocontext.Background())
			cancel()
			return ctx, cancel
		}},
		{"deadline", func() (gocontext.Context, gocontext.CancelFunc) {
			return gocontext.WithDeadline(gocontext.Background(), time.Now().Add(time.Hour))
		}},
		{"expired", func() (gocontext.Context, gocontext.CancelFunc) {
			return gocontext.WithDeadline(gocontext.Background(), time.Now().Add(-time.Hour))
		}},
	}

	for _, operation := range []string{"clone", "copy"} {
		for _, test := range tests {
			t.Run(operation+"/"+test.name, func(t *testing.T) {
				parent, cancel := test.prepare()
				defer cancel()
				parent = gocontext.WithValue(parent, key1, "caller")
				source := New()
				source.Set(key2, "source")

				var destination *Context
				if operation == "clone" {
					source.SetCancelContext(parent)
					destination = source.Clone()
				} else {
					other, stop := gocontext.WithCancel(gocontext.Background())
					stop()
					source.SetCancelContext(gocontext.WithValue(other, key1, "other"))
					destination = New().SetCancelContext(parent)
					source.CopyTo(destination)
				}

				if destination.Value(key1) != "caller" {
					t.Errorf("caller value = %v", destination.Value(key1))
				}
				if destination.Done() != parent.Done() {
					t.Error("cancellation channel was replaced")
				}
				if destination.Err() != parent.Err() {
					t.Errorf("context error = %v, want %v", destination.Err(), parent.Err())
				}
				deadline, ok := destination.Deadline()
				wantDeadline, wantOK := parent.Deadline()
				if ok != wantOK || !deadline.Equal(wantDeadline) {
					t.Errorf("deadline = %v, %v, want %v, %v", deadline, ok, wantDeadline, wantOK)
				}
				st.Expect(t, destination.Get(key2), "source")
				destination.Set(key2, "destination")
				st.Expect(t, source.Get(key2), "source")

				cancel()
				if destination.Err() != parent.Err() {
					t.Errorf("error after cancellation = %v, want %v", destination.Err(), parent.Err())
				}
			})
		}
	}
}

func TestContextCopy(t *testing.T) {
	ctx := New()
	ctx.Set("bar", "foo")
	st.Expect(t, ctx.Get("bar"), "foo")

	newCtx := New()
	ctx.CopyTo(newCtx)
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
	ctx.CopyTo(newCtx)
	st.Expect(t, newCtx.Get("bar"), "foo")

	// Ensure inmutability
	newCtx.Set("bar", "bar")
	st.Expect(t, ctx.Get("bar"), "foo")
	st.Expect(t, newCtx.Get("bar"), "bar")
}
