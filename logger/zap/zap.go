package main

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var hooks = func(e zapcore.Entry, d zapcore.SamplingDecision) {
	// If SamplingDecision is 1, ignore this entry.
	fmt.Printf("entry: %+v, dicision: %d\n", e, d)
}

func main() {
	conf := zap.NewDevelopmentConfig()
	conf.Sampling = &zap.SamplingConfig{
		Initial:    5,
		Thereafter: 10,
		Hook:       hooks,
	}
	// TODO: core required?
	// Build()はconf.Samplingがnilでなければ、内部でSamplerを初期化してくれる。その際のDurationは1s。
	l, err := conf.Build()
	if err != nil {
		panic(err)
	}
	defer l.Sync()

	for i := 1; i <= 100; i++ {
		l.Info(fmt.Sprintf("message(n=%d)", i))
	}
}
