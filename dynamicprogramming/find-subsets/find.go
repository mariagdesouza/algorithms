/*

Breaking up a problem into simpler subproblems, solving each of those subproblems just once, and storing their solutions
The important part of dynamic programming is storing the intermediate result set

Return all subsets of a set


*/
package main

import "fmt"

/* Find number of subsets that add up to a sum */

func main() {
	var findSum int
	fmt.Scan(&findSum)
	a := []int{3, 34, 4, 12, 5, 2, 9, 6}

	memo := make(map[int]int)

	//var result int
	fmt.Println(findSubset(a, findSum, len(a)-1))
	fmt.Println(findSubsetsDP(a, findSum, len(a)-1, memo))
}

// WIth memoization O(n)
func findSubsetsDP(a []int, sum, i int, memo map[int]int) int {

	if s, ok := memo[sum]; ok {
		return s
	}
	if sum == 0 { //empty set
		return 1
	}

	if sum < 0 || i < 0 {
		return 0
	}
	var result int
	if sum < a[i] {
		result = findSubsetsDP(a, sum, i-1, memo)
	} else {
		result = findSubsetsDP(a, sum-a[i], i-1, memo) + findSubsetsDP(a, sum, i-1, memo)
	}
	memo[sum] = result
	return result
}

// recursion O(2^n)
func findSubset(a []int, sum, n int) int {

	if sum == 0 {
		return 1
	}

	if sum < 0 || n < 0 {
		return 0
	}

	var result int
	if sum < a[n] { // sum = 9 element 12 then it cannot be part of a subset adding up to 9
		result += findSubset(a, sum, n-1)
	} else {
		result += findSubset(a, sum-a[n], n-1) + findSubset(a, sum, n-1)
	}

	return result
}
