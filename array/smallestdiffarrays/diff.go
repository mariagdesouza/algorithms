package main

import (
	"fmt"
	"sort"
)

/*
Given two arrays, compute the pair of values with the smallest non-negative diference


*/
func main() {
	a := []int{1, 3, 15, 11, 2}
	b := []int{23, 127, 235, 19, 8}

	sort.Ints(a) //O(mlogm)
	sort.Ints(b) ////O(nlogn)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("Min:", findMinDiff(a, b))

}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func findMinDiff(a, b []int) int {

	if len(a) == 0 || len(b) == 0 {
		return 0
	}

	min := abs(a[0] - b[0])
	i, j := 0, 0
	for i < len(a) && j < len(b) { //O(n)
		diff := abs(a[i] - b[j])
		if diff < min {
			//fmt.Println(diff)
			min = diff
		}

		//Move smaller value
		if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}

	return min
}
