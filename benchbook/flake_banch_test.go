package main

import (
	"testing"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/sony/sonyflake"
)

func BenchmarkSonyflake(b *testing.B) {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Date(2018, 1, 1, 0, 0, 0, 0, &time.Location{}),
	})
	for i := 0; i < b.N; i++ {
		sf.NextID()
	}
}

func BenchmarkSnowflake(b *testing.B) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		b.Log(err)
	}
	for i := 0; i < b.N; i++ {
		node.Generate()
	}
}
