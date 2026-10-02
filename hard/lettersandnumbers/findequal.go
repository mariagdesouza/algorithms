package main

import (
	"fmt"
	"unicode"
)

/*

Letters ad Numbers

Given an array filled with letters and numbers
find the longest subarray with an equal number of letters and numbers

   aaaaa11a1111aa11a
#a 12345556666678889
#1 00000122345666788



a1aaa1
*/

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func findLongestArray(a []rune) []rune {

	// compute deltaArray
	var deltaA, deltaN int
	delta := make([]int, len(a))
	maxDeltaIdx := -1
	for i, c := range a { //O(n)
		if unicode.IsNumber(c) {
			deltaN++
		} else {
			deltaA++
		}
		delta[i] = abs(deltaN - deltaA)
		fmt.Println(delta[i], deltaN, deltaA)
		if delta[i] == 0 && i > maxDeltaIdx {
			maxDeltaIdx = i
		}
	}
	fmt.Println(delta)
	fmt.Println(maxDeltaIdx)
	if maxDeltaIdx >= 0 && maxDeltaIdx < len(a) {
		return a[:maxDeltaIdx]
	}

	return nil
}

func main() {

	a := []rune{'a', 'a', 'a', 'a', '1', '1', 'a', '1', '1', '1', 'a', 'a', '1', '1', 'a'}

	fmt.Println(string(findLongestArray(a)))

}
