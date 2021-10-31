package popcount

import "testing"

func BenchmarkPopCount(b *testing.B) {
	PopCount(18446744073709551615)
}

func BenchmarkPopCount23(b *testing.B) {
	PopCount(18446744073709551615)
}
