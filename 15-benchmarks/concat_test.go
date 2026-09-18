package benchmarks

import "testing"

func BenchmarkConcatPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatPlus(1000)
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatBuilder(1000)
	}
}
