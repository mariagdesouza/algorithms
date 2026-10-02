package main

import "fmt"

//Given an array, generate all the possible subarrays of the given array using recursion.
// Note subarray is consecutive elements not a powersets that is all elements any order
//O(n2)

func main() {

	a := []int{1, 4, 6, 8}

	findPowerSet(a, 0, 0)
}

func findPowerSet(a []int, s, e int) {

	if e == len(a) {
		return
	}

	if s > e {
		findPowerSet(a, 0, e+1)
	} else {
		fmt.Println(a[s : e+1])
		findPowerSet(a, s+1, e)
	}

}
