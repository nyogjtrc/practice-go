package main

import (
	"errors"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// lumberjack.Logger is already safe for concurrent use, so we don't need to
	// lock it.
	wFile := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "foo.log",
		MaxSize:    1, // megabytes
		MaxBackups: 7,
		MaxAge:     1, // days
	})

	// stdout
	wStdout := zapcore.AddSync(os.Stdout)

	encoderConfig := zap.NewProductionConfig()
	encoderConfig.Build()

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
		zapcore.NewMultiWriteSyncer(wFile, wStdout),
		zap.NewAtomicLevel(),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	for i := 0; i < 10; i++ {
		logger.Warn("this is test message.")
		logger.Error("error", zap.Error(errors.New("test error")))
	}
}
