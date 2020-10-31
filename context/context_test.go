package ctx

import (
	"context"
	"testing"
)

type forCtxKey struct{}

var ctxKey = &forCtxKey{}

type val struct {
	s string
	i int
}

func TestCtx(t *testing.T) {
	ctx := context.Background()
	in := &val{s: "sss", i: 100}
	ctx = context.WithValue(ctx, ctxKey, in)

	out, ok := ctx.Value(ctxKey).(*val)
	if !ok {
		t.Fail()
	}
	if in != out {
		t.Fail()
	}
}

func TestCtx_Cancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	in := &val{s: "sss", i: 100}
	ctx = context.WithValue(ctx, ctxKey, in)
	cancel()
	out, ok := ctx.Value(ctxKey).(*val)
	if !ok {
		t.Fail()
	}
	if in != out {
		t.Fail()
	}
}
