package main

import "fmt"

const (
	row = 3
	col = 3
)

//Count paths from source to destination in a matrix in k step
/*

0 0 0
0 0 0
0 0 0

0 1 1
1 2 3
1 3 6
*/
func main() {

	//Init the matrix
	matrix := make([][]int, row)
	for j := 0; j < row; j++ {
		matrix[j] = make([]int, col)
	}

	paths := findPaths(matrix, &Point{0, 0}, &Point{row - 1, col - 1})

	fmt.Println(paths)
	fmt.Println(paths[row-1][col-1])
}

type Point struct {
	X int
	Y int
}

func findPaths(matrix [][]int, s, d *Point) [][]int {

	//Init the matrix
	paths := make([][]int, row)
	for j := 0; j < row; j++ {
		paths[j] = make([]int, col)
	}
	for i := s.X; i <= d.X; i++ {
		for j := s.Y; j <= d.Y; j++ {
			if i == s.X || j == s.Y { //edge squares set to 1
				paths[i][j] = 1
			} else {
				paths[i][j] = paths[i-1][j] + paths[i][j-1]
			}
		}
	}
	return paths
}
