package main

import (
	"testing"
)

/**


All benchmarks:
go test -bench=.

Specific benchmark:
go test -bench=Fib3

Different minimum benchmark time:
go test -bench=. -benchtime=5s (default - 1s)

 */

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(1, b)
}

func BenchmarkFib2(b *testing.B) {
	benchmarkFib(2, b)
}

func BenchmarkFib3(b *testing.B) {
	benchmarkFib(3, b)
}

func BenchmarkFib10(b *testing.B) {
	benchmarkFib(10, b)
}

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}
