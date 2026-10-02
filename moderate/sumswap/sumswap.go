package main

import "fmt"

func main() {
	a := []int{4, 1, 2, 1, 1}
	b := []int{3, 6, 3, 2}

	a1, b1 := sumswap(a, b)
	fmt.Println(a1, b1)
}

/*
	sumA = 9
	sumB = 15

sumA - a +b = sumB - a+ b
sumA -sumB = 2a - 2b


a-b = (sumA - sumB)/2

*/

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func sumswap(a, b []int) (int, int) {
	sumA, sumB := 0, 0
	for _, a1 := range a {
		sumA += a1
	}
	m := make(map[int]int)
	for i, b1 := range b {
		sumB += b1
		m[b1] = i
	}

	target := abs(sumA-sumB) / 2
	fmt.Println(sumA, sumB, target)

	for _, aa := range a {
		two := aa - target
		fmt.Println(two)
		if _, ok := m[two]; ok {
			return aa, two
		}
	}

	return -1, -1
}
