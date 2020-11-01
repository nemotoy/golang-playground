package ctx

import (
	"context"
	"testing"
	"time"
)

type forCtxKey struct{}

var ctxKey = &forCtxKey{}

type val struct {
	s string
	i int
}

// go1.15/src/context/context.go:532
// WithValue()のレシーバは、 valueCtxであり Context インターフェースが埋め込まれているため、Contextメソッドも利用できる実装である。
// Contextインターフェースへの操作が実施されても、key/valueへの副作用は発生しない。(valueCtx, timerCtx, cancelCtx)
// https://blog.golang.org/context
// https://blog.gopheracademy.com/advent-2016/context-logging/
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

func TestCtx_DeadLine(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
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
