// nolint
package main

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
	ref.
	https://github.com/uber-go/zap/issues/588
	https://github.com/uber-go/zap/issues/874
*/

var hook = func(e zapcore.Entry, d zapcore.SamplingDecision) {
	// If SamplingDecision is 1, ignore this entry.
	fmt.Printf("entry: %+v, dicision: %d\n", e, d)
}

func main() {
	conf := zap.NewDevelopmentConfig()
	conf.Sampling = &zap.SamplingConfig{
		Initial:    5,
		Thereafter: 10,
		Hook:       hook,
	}
	// Build()はconf.Samplingがnilでなければ、内部でSamplerを初期化してくれる。その際のDurationは1s。
	logger, err := conf.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	start := time.Now()
	for i := 1; i <= 100; i++ {
		logger.Info("Info message",
			zap.String("url", "someurl"),
			zap.Int("attempt", i),
			zap.Duration("backoff", time.Second),
		)
	}
	elapsed := time.Since(start)
	fmt.Printf("took %s\n", elapsed)
}
