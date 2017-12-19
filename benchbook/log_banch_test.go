package main

import (
	"log"
	"testing"

	"go.uber.org/zap"
)

func BenchmarkLog(b *testing.B) {
	temp := "temp variable"
	for i := 0; i < b.N; i++ {
		log.Println("log message: ", temp)
	}
}

func BenchmarkZap(b *testing.B) {
	temp := "temp variable"
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	for i := 0; i < b.N; i++ {
		logger.Info("log message: ", zap.String("temp", temp))
	}
}
