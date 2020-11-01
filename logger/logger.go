package main

import (
	"context"

	"go.uber.org/zap"
)

type forCtxKey struct{}

var ctxKey = &forCtxKey{}

type hoge struct {
	f []zap.Field
}

type obj struct {
	l *zap.Logger
}

func main() {

	ctx := context.Background()
	l := zap.NewExample()
	obj := &obj{l: l}
	fields := map[string]string{"key1": "val1", "key2": "val2"}
	ff := []zap.Field{}
	for k, v := range fields {
		ff = append(ff, zap.String(k, v))
	}
	l.Debug("#main", ff...)

	ctx = context.WithValue(ctx, ctxKey, &hoge{f: ff})
	obj.run(ctx)
}

func (o *obj) run(ctx context.Context) {
	val := ctx.Value(ctxKey).(*hoge)
	o.l.Debug("#run", val.f...)

	s, err := o.service(ctx)
	if err != nil {
		panic(err)
	}
	val.f = append(val.f, zap.String("#service", s))
	o.l.Debug("#run", val.f...)
}

func (o *obj) service(ctx context.Context) (string, error) {
	val := ctx.Value(ctxKey).(*hoge)
	o.l.Debug("#service", val.f...)

	return "service", nil
}
