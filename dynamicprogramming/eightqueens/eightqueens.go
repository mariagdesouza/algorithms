package main

import (
	"fmt"
	"math"
)

/*

eight queens have lined up on 8x8 chess board
None share the same row, column or diagonal

row = 0 no other queen can have row 0

ways to arrange at (7,0) + ways at (7,1) + ways at (7,2) ... (7,7)


*/

const (
	N = 8
)

func main() {
	//fmt.Println(board)

	columns := make([]int, N)

	for col := 0; col < N; col++ {
		columns[col] = -1
	}

	placeQueens(0, columns)

	fmt.Println(columns)

}

// start with row = 0 and recursivelly call till row == N
func placeQueens(row int, columns []int) bool {
	//fmt.Println(columns, row)
	if row == N {
		return true
	}
	for col := 0; col < N; col++ { // for each row check each column
		if checkValid(columns, row, col) == true {
			columns[row] = col
			if placeQueens(row+1, columns) { // recursively call for next row
				return true
			}
		}
	}
	return false
}

//check if row, col is a valid spot
func checkValid(columns []int, row, col int) bool {

	for r1 := 0; r1 < row; r1++ { // go through all rows prior if it invalidates this row
		c1 := columns[r1]
		if c1 == col { // if rows have a queen in the same column - invalid
			return false
		}

		// if distance between columns is equal to to distance between rows, the daigonal is the same - invalid

		if math.Abs(float64(row-r1)) == math.Abs(float64(col-c1)) {
			return false
		}

	}
	return true
}
