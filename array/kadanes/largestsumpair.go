package main

import (
	"fmt"
)

func main() {
	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(a, "kadaneMaxsum", kadaneMaxsum(a))
	resSum, resArr := kadaneMaxSubArray(a)
	fmt.Println(a, "kadaneMaxSubArray", resSum, resArr)

	b := []int{1, 3, 2, 3, 4, 8, 7, 9}
	fmt.Println(kadaneMaxSubArray(b))

	c := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	resSum, resArr = kadaneMaxSubArray(c)
	fmt.Println(c, "kadaneMaxsum", kadaneMaxsum(c))
	fmt.Println(c, "kadaneMaxSubArray", resSum, resArr)

	fmt.Println(kadaneMaxsum([]int{-1, -2, -3, -4}))

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func kadaneMaxsum(a []int) int {
	max_current := a[0]
	max_global_sum := a[0]

	for i := 1; i < len(a); i++ {
		max_current = max(max_current+a[i], a[i])
		max_global_sum = max(max_current, max_global_sum)
	}
	return max_global_sum
}

func kadaneMaxSubArray(a []int) (int, []int) {
	maxCurrent, max_global_sum := a[0], a[0]
	startOld, start, end := 0, 0, 0
	for i := 1; i < len(a); i++ {
		//	prev := maxCurrent
		maxCurrent += a[i]
		if maxCurrent > max_global_sum {
			max_global_sum = maxCurrent
			startOld = start
			end = i
		}
		if max_global_sum < 0 {
			maxCurrent = 0
			start = i + 1
		}

	}

	return max_global_sum, a[startOld : end+1]
}
