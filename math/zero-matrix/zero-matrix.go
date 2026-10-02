package main

import "fmt"

const (
	ROW = 3
	COL = 4
)

// if an element of a MxN matrix is zero, set the row and column to zero
func main() {

	matrix := [][]int{
		{1, 2, 3, 4},
		{0, 0, 6, 7},
		{8, 9, 10, 11},
	}
	for _, row := range matrix {
		fmt.Println(row)
	}

	flaggedRow := make([]bool, ROW)
	flaggedCol := make([]bool, COL)

	//O(M * N)
	for i := 0; i < ROW; i++ {
		for j := 0; j < COL; j++ {
			if matrix[i][j] == 0 {
				flaggedRow[i] = true
				flaggedCol[j] = true
				continue
			}
		}
	}

	//O(N)
	for i := 0; i < ROW; i++ {
		if flaggedRow[i] {
			setRow(matrix, i)
		}
	}
	//O(M)
	for j := 0; j < COL; j++ {
		if flaggedCol[j] {
			setCol(matrix, j)
		}
	}

	fmt.Println("OUTPUT")
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func setRow(matrix [][]int, row int) {
	for i := 0; i < COL; i++ {
		matrix[row][i] = 0
	}
}

func setCol(matrix [][]int, col int) {
	for i := 0; i < ROW; i++ {
		matrix[i][col] = 0
	}
}
