package main

import "fmt"

// given two sorted arrays, A and B, A has a large enough buffer to hold B
// Write a method to merge B into A in sorted array

func main() {
	a := [16]int{1, 2, 3, 4, 5, 6, 7, 8}
	b := [8]int{2, 4, 6, 8, 10, 12, 14, 16}

	merge(&a, &b)
	fmt.Println(a)
}
func sortedmerge(a *[16]int, b *[8]int) {
	indexB := len(b) - 1
	indexA := (len(a) - len(b)) - 1
	indexMerged := len(a) - 1
	fmt.Println(indexA, indexB, indexMerged)

	for indexB >= 0 && indexMerged >= 0 {
		if indexA >= 0 && a[indexA] > b[indexB] {

			a[indexMerged] = a[indexA]
			//fmt.Println("a", indexA, indexB, indexMerged, a[indexMerged])
			indexA--
		} else {

			a[indexMerged] = b[indexB]
			//fmt.Println("b", indexA, indexB, indexMerged, a[indexMerged])

			indexB--
		}

		indexMerged--
	}
}
