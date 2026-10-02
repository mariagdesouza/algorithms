/*
A child is running up a staircase with n steps and can hop either 1 step, 2 steps, or 3 steps
at a time. Implement a method to count how many possible ways the child can run up the stairs.

There are two methods to solve this problem
1. Recursive Method
2. Dynamic Programming

*/
package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)
	//fmt.Println(countStepsRecursion(n))
	fmt.Println(countStepDynamic(n))
}

//O(3raise to n)
func countStepsRecursion(n int64) int64 {
	if n == 1 || n == 0 {
		return 1
	}
	if n == 2 {
		return n
	}
	return countStepsRecursion(n-1) + countStepsRecursion(n-2) + countStepsRecursion(n-3)
}

func countStepDynamic(n int64) int64 {
	if n < 0 {
		return 0
	}

	if n <= 1 {
		return 1
	}

	res := make([]int64, n+1)
	res[0] = 1
	res[1] = 1
	res[2] = 2

	for i := int64(3); i <= n; i++ {
		res[i] = res[i-1] + res[i-2] + res[i-3]
	}
	return res[n]
}
