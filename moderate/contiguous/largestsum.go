package main

import "fmt"

/*
Variation of Kadane;s algorithm
- Find largets sum of contiguous elements


*/
func main() {
	a := []int{2, -8, 3, -2, 4, -10}

	fmt.Println(findmaxsum(a))
}

func findmaxsum(a []int) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}

	maxsum := a[0]
	sum := 0
	//max := a[0]
	for i := 1; i < len(a); i++ {
		sum += a[i]
		if sum > maxsum {
			maxsum = sum
		} else if sum < 0 {
			sum = 0
			//max = a[i]
		}
	}

	return maxsum
}
