package main

import "testing"

func BenchmarkcheckPermutation1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		checkPermutation("icecream", "miceacre")
	}
}
