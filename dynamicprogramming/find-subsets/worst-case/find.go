/*
  find subsets that add to a certain number
*/
package main

import "fmt"

func main() {
	var findSum int
	fmt.Scan(&findSum)
	a := []int{3, 34, 4, 12, 5, 2, 9}
	fmt.Println(countSubsets(a, findSum, len(a)-1))
}

func countSubsets(a []int, findSum int, i int) int {
	//if sum is 0 return 1
	if findSum == 0 {
		return 1
	}
	if findSum < 0 || i < 0 {
		return 0
	}

	if findSum < a[i] { //cant use i if its greater tahn sum - look for i-1
		return countSubsets(a, findSum, i-1)
	} else { // find those that add upto a[i] and look for i-1
		return countSubsets(a, findSum-a[i], i-1) + countSubsets(a, findSum, i-1)
	}

}
