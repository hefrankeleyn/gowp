package q2

import (
	"testing"
)

func BenchmarkGetPrimes(b *testing.B) {
	// runtime.GOMAXPROCS(5)
	for i := 0; i < b.N; i++ {
		GetPrimes(1000)
	}
}
