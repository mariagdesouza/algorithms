package main

import "fmt"

/*
Find length of longest sub array
*/

func main() {
	b := []int{1, 3, 2, 3, 4, 8, 7, 9}
	fmt.Println(findLengthLongestSubarray(b))
}

func findLengthLongestSubarray(a []int) int {
	var maxlen int
	length := 1
	for i := 1; i < len(a); i++ {
		if a[i] > a[i-1] {
			length++
		} else { //reset
			if maxlen < length { //if the length at this point is greater set that
				maxlen = length
			}
			length = 1 // reset value
		}
	}
	if maxlen < length { //if the length at this point is greater set that
		maxlen = length
	}

	return maxlen
}

func max(x, y int) int {

	if x > y {
		return x
	}
	return y
}

func finMaxNonadjacent(a []int) int {
	incl := a[0]
	excl := 0

	for i := 1; i < len(a); i++ {
		excl_new := max(incl, excl)

		incl = excl + a[i] // current max
		excl = excl_new
	}
	return max(excl, incl)
}
