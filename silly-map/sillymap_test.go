package sillymap

import "testing"

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Map()
	}
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Array()
	}
}
