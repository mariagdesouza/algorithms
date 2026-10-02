package main

import "testing"

func BenchmarkisAnagram1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isAnagram("icecream", "miceacre")
	}
}
