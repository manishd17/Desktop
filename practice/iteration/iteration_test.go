package iteration

import "testing"

func BenchmarkRepeat(b *testing.B) {
	b.N = 1000000
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
