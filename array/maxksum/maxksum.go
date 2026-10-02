package main

import "fmt"

/*
Find maximum (or minimum) sum of a subarray of size k
Given an array of integers and a number k, find maximum sum of a subarray of size k.

*/

func main() {
	testTable := []struct {
		input    []int
		k        int
		expected int
	}{
		{[]int{100, 200, 300, 400}, 2, 700},
		{[]int{1, 4, 2, 10, 23, 3, 1, 0, 20}, 4, 39},
		{[]int{1, 4, 2, 10, 23, 3, 1, 0, 20}, 3, 36},
		{[]int{2, 3}, 3, 0},
	}

	for _, v := range testTable {
		fmt.Println(" Got:", maxKSum(v.input, v.k))
		fmt.Println(" Expected:", v.expected)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxKSum(a []int, k int) int {

	if len(a) == 0 || k > len(a) {
		return 0
	}
	maxsum := 0

	for i := 0; i < k; i++ { // calculate first k
		maxsum += a[i]
	}

	currentSum := maxsum
	for i := k; i < len(a); i++ {
		// k = 2; a[3]- a[0]
		currentSum += a[i] - a[i-k]      //  here we are removing first element from sum and adding a new element
		maxsum = max(currentSum, maxsum) // if greater than previous set of k; replace
	}

	return maxsum
}
