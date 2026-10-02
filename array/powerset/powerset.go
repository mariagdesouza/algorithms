package main

import (
	"fmt"
	"math"
)

/*
Find all sets in set
*/
func main() {

	a := []int{1, 4, 6, 8}

	fmt.Println(findPowerSet(a))
}
func findPowerSet(a []int) [][]int {
	powerSetSize := int64(math.Pow(2, float64(len(a))))

	var result [][]int

	for i := int64(0); i < powerSetSize; i++ {
		var subset []int
		for j, elem := range a {
			if i&(1<<uint(j)) > 0 { // if jth bit is set use jth element
				subset = append(subset, elem)
			}
		}
		//fmt.Println(subset)
		result = append(result, subset)
	}

	return result
}
