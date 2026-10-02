package main

import (
	"fmt"
	"sort"
)

/*

Given an array of ints

find m and n such that if you  sort elements m to n  the entire array would be sorted

*/

func main() {
	a := []int{1, 2, 4, 7, 10, 12, 8, 11, 5, 6, 16, 18, 19}
	fmt.Println(a)
	start, end := findSubsortIndex(a)

	fmt.Println(subsort(a, start, end))

	b := []int{14, 6, 12, 8, 5, 6, 16, 18, 19}
	fmt.Println(b)
	fmt.Println(findSubsortIndex(b))
}

func subsort(a []int, start, end int) []int {
	left := a[:start]

	middle := a[start:end]

	if len(middle) > 0 {
		sort.Ints(middle)
	}

	result := append(left, middle...)

	if end < len(a)-1 {
		right := a[end+1:]
		result = append(result, right...)
	}
	return result
}

func findSubsortIndex(a []int) (int, int) {

	if len(a) < 2 {
		return 0, 0
	}
	start := 0
	end := len(a) - 1

	for start < len(a)-1 && a[start] < a[start+1] {
		start++
	}

	for end > start && a[end] > a[end-1] {
		end--
	}

	return start, end
}
