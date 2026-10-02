package main

import "fmt"

/*


 */

const (
	ROW = 4
	COL = 4
)

func main() {
	matrix := [][]int{
		{0, 2, 1, 0},
		{0, 1, 0, 1},
		{1, 1, 0, 1},
		{0, 1, 0, 1},
	}
	s := computePondSizes(matrix)
	fmt.Println(s)
}

func computePondSizes(m [][]int) []int {

	var sizes []int

	for r := 0; r < ROW; r++ {
		for c := 0; c < COL; c++ {

			if m[r][c] == 0 {
				s := computeSize(m, r, c)
				sizes = append(sizes, s)
			}
		}
	}

	return sizes
}

func computeSize(m [][]int, r, c int) int {

	if r < 0 || c < 0 || r >= ROW || c >= COL || m[r][c] != 0 {
		return 0
	}
	size := 1
	m[r][c] = -1                     //visited
	size += computeSize(m, r+1, c)   //vertically
	size += computeSize(m, r, c+1)   //horizontally
	size += computeSize(m, r+1, c+1) //diagonally
	return size
}
