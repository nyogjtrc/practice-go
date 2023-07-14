package main

import (
	"encoding/json"
	"testing"
)

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3})

	}
}

func BenchmarkDecode2Struct(b *testing.B) {
	jsonString := []byte("{\"title\":\"Document\",\"author\":\"me\"}")
	for i := 0; i < b.N; i++ {
		b := Book{}
		json.Unmarshal(jsonString, &b)
	}
}

func BenchmarkDecode2Map(b *testing.B) {
	jsonString := []byte("{\"title\":\"Document\",\"author\":\"me\"}")
	for i := 0; i < b.N; i++ {
		m := make(map[string]string)
		json.Unmarshal(jsonString, &m)
	}
}
