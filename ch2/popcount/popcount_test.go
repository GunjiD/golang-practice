package popcount

import "testing"

func BenchmarkPopCount(b *testing.B) {
	PopCount(1)
}
