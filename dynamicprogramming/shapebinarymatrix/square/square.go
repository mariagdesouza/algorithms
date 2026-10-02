package main

import "fmt"

const (
	R = 6
	C = 5
)

// Find Maximum size square  sub-matrix with all 1s
func main() {

	m := [R][C]int{
		{0, 1, 1, 0, 1},
		{1, 1, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0}}

	fmt.Println(findMaxSquare(m))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//Find all top-left corners of squares in binary 2D matrix.
func findAlltopLeftCorners() {

}

func findMaxSquare(m [R][C]int) int {

	maxArea := m[0][0]
	for i := 0; i < R; i++ { //O(m *n)
		for j := 0; j < C; j++ {
			if i > 0 && j > 0 && m[i][j] == 1 {
				m[i][j] = min(min(m[i-1][j], m[i][j-1]), m[i-1][j-1]) + 1
				if m[i][j] > maxArea {
					maxArea = m[i][j]
				}
			}
		}
	}

	return maxArea
}
