package main

import (
	"fmt"
	"sort"
)

/* unsorted array medians*/

func main() {
	a := []int{2, 10, 3, 6, 8, 5, 7, 13, 16}
	fmt.Println(medianOfMedians(a, 0, len(a)-1, len(a)/2))
}

// 2, 10, 3, 6, 8, 5
// 2, 3, 5, 6, 7 8, 10

func medianOfMedians(a []int, lo, hi, k int) int {

	if lo == hi {
		return a[lo]
	}
	n := len(a)
	if n <= 5 {
		return lazyMedian(a)
	}

	med := ((n - 1) + 5) / 5

	medians := make([]int, med)
	i := 0
	for ; i < med; i++ {
		start := (i * 5)
		end := (i * 5) + 5

		var s []int
		if end >= n {
			s = a[start:]
		} else {
			s = a[start:end]
		}

		medians[i] = lazyMedian(s)

	}
	medOfMed := medianOfMedians(medians, 0, i-1, i/2)

	p := partition(a, lo, hi, medOfMed)

	if k == p {
		return a[p]
	} else if k < p {
		return medianOfMedians(a, lo, p-1, k)
	}

	return medianOfMedians(a, p+1, hi, k-p)

}

func partition(a []int, lo, hi, p int) int {

	i := lo
	j := hi
	for {
		for ; a[i] < p; i++ {

		}
		for ; a[j] > p; j-- {

		}
		if i >= j {
			return j
		}
		a[i], a[j] = a[j], a[i]
	}
}

func lazyMedian(a []int) int { // O(nlogn) to sort
	n := len(a)
	if n < 1 {
		return 0
	}
	if n == 1 {
		return a[0]
	}
	sort.Ints(a)
	return a[n/2]
}
