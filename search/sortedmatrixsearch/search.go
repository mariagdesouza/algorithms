package main

import "fmt"

/*
Sorted Matrix Search
Given an MxN matrix where each row is and colum is sorted in ascending order,
write a method to find an element

5  10 15
12 23 26
14 35 37

23

[0][c-1] > num


*/

const (
	M, N = 4, 4
)

func main() {

	a := [][]int{
		{5, 10, 15, 31},
		{12, 23, 26, 33},
		{14, 35, 37, 44},
		{16, 36, 38, 45},
	}
	fmt.Println(find(a, 35))
	fmt.Println(find(a, 33))
}

func find(a [][]int, num int) bool {
	row, col := 0, N-1

	for row < M && col >= 0 {
		fmt.Println(a[row][col])
		if a[row][col] < num {
			row++
		} else if a[row][col] > num {
			col--
		} else {
			return true
		}
	}
	return false
}
