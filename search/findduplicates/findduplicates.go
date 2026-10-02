package main

import (
	"fmt"
	"math"
)

/*
Find Duplicates

you have an array with numbers 1 to N
max N   = 32000

The array may have duplicates

N is unknown

with 4KB of memory  - how to print all duplicates in the array

*/

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	result := n * factorial(n-1)
	return result
}

func findDuplicate(a []int, n int) int {

	for _, d := range a {
		abs := int(math.Abs(float64(d)))
		if a[abs] > 0 {
			a[abs] = a[abs] * -1
		} else {
			return abs
		}
	}
	return -1
}

func main() {
	a := []int{4, 3, 4, 5, 1}
	fmt.Println(findDuplicate(a, 5))

}
