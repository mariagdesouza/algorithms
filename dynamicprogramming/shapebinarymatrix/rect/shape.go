package main

import "fmt"

//finding biggest rectangle in 2d matrix of 0 and 1 is quite popular question
//size of the rectangle - array/DFS

const (
	R = 7
	C = 6
)

func main() {
	m := [][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 1, 1},
		{1, 0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1}}

	//fmt.Println(findMaxRect(m))
	countRectangles(m)
}

func countRectangles(m [][]int) {
	count := 0
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if m[i][j] == 0 {
				fmt.Print(i, j)
				//checkRectangle(i, j, m)
				checkRectangleRec(i, j, m, len(m[i]))
				count++
			}
		}
	}
	fmt.Println(count)
}

func checkRectangleRec(x, y int, m [][]int, c int) {
	if m[x][y] != 0 {
		fmt.Print(" ", x-1, c-1)
		fmt.Println()
		return
	}
	i := y
	for ; i < C; i++ {
		if m[x][i] != 0 {
			break
		}
		m[x][i] = 1
	}
	checkRectangleRec(x+1, y, m, i)
}

func checkRectangle(x, y int, m [][]int) {
	i := y
	for ; i < C; i++ {
		if m[x][i] != 0 {
			break
		}
		m[x][i] = 1
	}
	//fmt.Println("Stopping at ", x, i)
	j := x + 1
Row:
	for j < R {
		k := y
		for ; k < i; k++ {
			if m[j][k] != 0 {
				if k == i {
					break
				} else {
					break Row
				}
			}
		}
		if k == i {
			for k = y; k < i; k++ {
				m[j][k] = 1
			}
		} else {
			break
		}
		j++
	}
	fmt.Print(" ", j-1, i-1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findMaxRect(m [R][C]int) int {

	//For each cell (i, j), we store the largest value of K such that
	//K x K is a submatrix with all equal elements and position of (i, j)
	//being the bottom-right most element.
	result := maxHistogram(m[0])

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			// if A[i][j] is 1 then add A[i -1][j]
			if m[i][j] == 1 && i > 0 {
				m[i][j] += m[i-1][j]
			}
			r := maxHistogram(m[i])
			result = max(result, r)
		}
	}
	return result
}

// Create an empty stack. The stack holds indexes of
// hist[] array/ The bars stored in stack are always
// in increasing order of their heights.
func maxHistogram(row [C]int) int {
	maxArea := 0

	var stack []int
	// add to end - remove from end - top is end

	i := 0
	area := 0
	for i < C {
		if len(stack) == 0 || row[i] > stack[len(stack)-1] {
			stack = append(stack, row[i]) //push
			i++
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			area = top * i
			if len(stack) > 0 {
				area = top * (i - stack[len(stack)-1] - 1)
			}
			maxArea = max(area, maxArea)
		}

	}
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(stack) > 0 {
			area = top * (i - stack[len(stack)-1] - 1)
		}
		maxArea = max(area, maxArea)
	}
	return maxArea
}
