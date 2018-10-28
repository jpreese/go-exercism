package suchrandom

import "testing"

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunMap()
	}
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunSlice()
	}
}
