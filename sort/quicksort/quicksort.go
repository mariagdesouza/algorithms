package main

import "fmt"

/*

Quicksort
Divide and conquer algorithm

Average complexity:  O(nlogn)
Worst case O(n2) but very rare

1. Pick a pivot element
2. Partition based on pivot
3. recursively apply 1 and 2


*/

func main() {
	a := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	quicksort(a, 0, len(a)-1)
	fmt.Println(a)
}

func quicksort(a []int, lo, hi int) {
	if lo < hi {
		//Partition based on pivot
		p := partition(a, lo, hi)

		quicksort(a, lo, p)   // lo ... p
		quicksort(a, p+1, hi) // p+1 .. hi
	}
}
func partition(a []int, low, high int) int {
	pivot := a[(low+high)/2]
	//fmt.Println(pivot, low, high)
	i := low
	j := high

	for {
		for ; a[i] < pivot; i++ {
		}
		for ; a[j] > pivot; j-- {
		}
		if i >= j {
			return j
		}
		//lo < p <hi
		a[i], a[j] = a[j], a[i]
		//fmt.Println(i, j, a)
	}
}
