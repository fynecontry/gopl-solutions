package popcountv1_test

import (
	"testing"

	"popcount"
	"popcountv1"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountv1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcountv1.PopCount(0x1234567890ABCDEF)
	}
}

// $ go test -bench=. ch2/popcountv1/popcountv1_test.go
// goos: linux
// goarch: amd64
// cpu: Intel(R) Core(TM) i5-4200U CPU @ 1.60GHz
// BenchmarkPopCount-4             1000000000               0.4225 ns/op
// BenchmarkPopCountv1-4           129425227                8.592 ns/op
// PASS
// ok      command-line-arguments  2.264s
