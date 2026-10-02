package main

import (
	"fmt"
	"sort"
)

/*

Median - is the
- middle element if it is odd
- (sum of middle two)/2 if it is even

Method
 1.  sort and find middle O(nlogn)

if data is partially sorted use randomised quickselect

if data is unsorted use normal quickselect (starting from first index as pivot element

choice is three methods
i) medians of medians
ii) randomised quickselect
iii) normal quickselect

NOTE:

Mean = sum of elements / num of elements
*/

func main() {
	a := []int{4, 1, 6, 3, 9, 3, 7, 1, 2, 1, 1, 1, 1, 1}

	fmt.Println("median of", a, "is:")
	fmt.Println(lazymedian(a))

	//	fmt.Println(medianOfmedians2(a, 0, len(a)-1, len(a)/2))
	fmt.Println(medianOfmedians(a, 0, len(a)-1, len(a)/2))

	//fmt.Println(medianOfmedians([]int{2, 10, 3, 6, 8, 5, 7, 13}, 8))

	b := []int{2, 10, 3, 6, 8, 5, 7, 13}
	fmt.Println("median of", b, "is:")
	//fmt.Println(medianOfmedians2(b, 0, len(b)-1, len(b)/2))
	fmt.Println(medianOfmedians(b, 0, len(b)-1, len(b)/2))

	c := []int{7, 10, 4, 3, 20, 15}
	fmt.Println("median of", c, "is:")
	fmt.Println(medianOfmedians(c, 0, len(c)-1, len(c)/2))
}

// This is deterministic linear-time selection algorithm
// 1. divide a list into sublists and then determines the median in each of the sublists
// 2. takes those medians and puts them into a list and finds the median of that list.
// 3. Use that median value as a pivot and compares other elements of the list against the pivot
// 4. If an element is less than  pivot, it is placed  left of the pivot,
//     if the element is greater than the pivot, it is placed to the right.
// 5. The algorithm recurses on the list, honing in on the value it is looking for.
func medianOfmedians(a []int, lo, hi, k int) int {

	if lo > hi {
		return -1
	}

	if lo == hi {
		return a[lo]
	}
	n := len(a)
	if n <= 5 {
		return lazymedian(a)
	}
	// provision for a median array
	med := ((n - 1) + 5) / 5
	medians := make([]int, med)
	i := 0
	for ; i < med; i++ { //O(1)
		end := (i * 5) + 5 // 5, 10, 15, 20 ...
		var s []int
		if end >= n { //5 or less left - take till end
			s = a[(i * 5):]
		} else {
			s = a[(i * 5):end] // take 5-10, 10-15
		}
		medians[i] = lazymedian(s)
	}

	var medOfMed int
	if i == 1 {
		medOfMed = medians[0]
	} else {
		medOfMed = medianOfmedians(medians, 0, i-1, i/2)
	}
	// Partition the array around p  - O(n)
	p := partition(a, lo, hi, medOfMed)

	if k == p {
		return a[p]
	} else if k < p {
		return medianOfmedians(a, lo, p-1, k)
	}
	return medianOfmedians(a, p+1, hi, (k - (p + 1) - 1))
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

func lazymedian(a []int) int {

	n := len(a)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return a[0]
	}

	sort.Ints(a)
	return a[(n / 2)]
}
