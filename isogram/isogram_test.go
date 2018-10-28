package isogram

import "testing"

func BenchmarkIsIsogramAppendRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			IsIsogramAppendRune(c.input)
		}
	}
}

func BenchmarkIsIsogramSortWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			IsIsogramSortWord(c.input)
		}
	}
}

func BenchmarkIsIsogramUsingMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			IsIsogramUsingMap(c.input)
		}
	}
}
