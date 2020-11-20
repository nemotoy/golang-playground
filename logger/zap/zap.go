// nolint
package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	logger := fromConf()
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

func fromConf() *zap.Logger {
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
	return logger
}

func fromCore() *zap.Logger {
	// The bundled Config struct only supports the most common configuration
	// options. More complex needs, like splitting logs between multiple files
	// or writing to non-file outputs, require use of the zapcore package.
	//
	// In this example, imagine we're both sending our logs to Kafka and writing
	// them to the console. We'd like to encode the console output and the Kafka
	// topics differently, and we'd also like special treatment for
	// high-priority logs.

	// First, define our level-handling logic.
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	// Assume that we have clients for two Kafka topics. The clients implement
	// zapcore.WriteSyncer and are safe for concurrent use. (If they only
	// implement io.Writer, we can use zapcore.AddSync to add a no-op Sync
	// method. If they're not safe for concurrent use, we can add a protecting
	// mutex with zapcore.Lock.)
	topicDebugging := zapcore.AddSync(ioutil.Discard)
	topicErrors := zapcore.AddSync(ioutil.Discard)

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	// Optimize the Kafka output for machine consumption and the console output
	// for human operators.
	kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core := zapcore.NewTee(
		zapcore.NewCore(kafkaEncoder, topicErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(kafkaEncoder, topicDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	opt := []zap.Option{
		zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return zapcore.NewSamplerWithOptions(core, time.Second, 5, 10, zapcore.SamplerHook(hook))
		}),
	}

	// From a zapcore.Core, it's easy to construct a Logger.
	return zap.New(core, opt...)
}
