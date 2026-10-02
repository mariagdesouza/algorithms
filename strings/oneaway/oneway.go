package main

import (
	"fmt"
	"math"
)

// edits - insert a char; remove a char or replace a char
// write a function to check if it is one or zero edits away

// Dynamic programming - If you have solved a problem with the given input, then save the result for future reference,
// to avoid solving the same problem again.. shortly
//'Remember your Past' :) . - memoization

func main() {
	fmt.Println("string1:")
	var str1 string
	fmt.Scan(&str1)

	fmt.Println("string2:")
	var str2 string
	fmt.Scan(&str2)

	fmt.Println(isOneAway(str1, str2))

}

// pale ple
// pales pale

func isOneAway(str1, str2 string) bool {

	// 1. If difference between lengths is more that one
	m := len(str1)
	n := len(str2)
	if math.Abs(float64(m-n)) > 1 {
		return false
	}

	var i, j, diffs int
	for i < m && j < n {
		if str1[i] == str2[j] {
			i++
			j++
		} else {
			diffs++
			if m > n {
				i++
			} else if m < n {
				j++
			}
		}
	}
	if i < m || j < n {
		diffs++
	}
	if diffs > 1 {
		return false
	}
	return true
}
