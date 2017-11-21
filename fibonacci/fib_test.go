package main

import "testing"

func BenchmarkFib10(b *testing.B) {
	fib(10)
}

func BenchmarkFib40(b *testing.B) {
	fib(40)
}
