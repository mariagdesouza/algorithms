package main

import "fmt"

/*
a robot sitting on the upper left corner of grid with r rows and c columns.
The robot can only move in two directions, right and down, but certain cells are "off limit" such that the robot cannot step on them.
Design an algorithm to find a path for the robot from the top left to the bottom right.

m x n  grid
robot
start - 0,0


end  - m-1 x n-1



obstacleGrid {
	{ 0, 0, 0, 0},
	{ 1, 0, 1, 0},
	{ 0, 1, 0, 0}
}

if obstacleGrid[0][0] == 1 {
	return 0 //no path
}

if obstacleGrid[0][1] == 1 && obstacleGrid[1][0] == 1{
	return 0 //no path
}

for r = 0; r <m; r++{
	for c =0; c<n; c++{

	if obstacleGrid[r][c] == 1 {
		obstacleGrid[r][c] = 0 //obstackle - no path
		continue
	}

	if (r == 0 && c == 0){ // 1st wil be 1
		obstacleGrid[r][c] = 1
	}

	if obstacleGrid[r][c+1] == 1 && obstacleGrid[r+1][c] == 1{
		obstacleGrid[r][c] = 0 //continue
		continue //no path
	}

	// if find number of paths
	if (r == 0 ){
		obstacleGrid[r][c] = obstacleGrid[r][c-1]
	}else if (c == 0 ){
		obstacleGrid[r-1][c] = obstacleGrid[r-1][c]
	} else {
		obstacleGrid[r][c] == obstacleGrid[r-1][c] +  obstacleGrid[r][c-1]
	}

	}

	no of ways to get here = obstacleGrid[m-1][n-1]
}


obstacleGrid {
	{ 1, 1, 1, 1},
	{ 0, 1, 0, 1},
	{ 0, 0, 0, 1}
}





*/

func main() {

	obstacleGrid := [][]int{
		{0, 0, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	}
	m := 3
	n := 4
	getPath(obstacleGrid, m, n)
}

func getPath(obstacleGrid [][]int, m int, n int) bool {

	if (m == 0 && n == 0) || obstacleGrid[0][0] == 1 {
		return false //no path
	}

	/*
		if  obstacleGrid[0][1] == 1 && obstacleGrid[1][0] == 1 {
			return false //no path
		}*/

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {

			if obstacleGrid[r][c] == 1 {
				obstacleGrid[r][c] = 0 //obstacle - no path
				continue
			}

			if r == 0 && c == 0 { // 1st wil be 1
				obstacleGrid[r][c] = 1
				continue
			}

			if c+1 < n && r+1 < m && obstacleGrid[r][c+1] == 1 && obstacleGrid[r+1][c] == 1 {
				obstacleGrid[r][c] = 0 //continue
				continue               //no path
			}

			// if find number of paths
			if r == 0 {
				obstacleGrid[r][c] = obstacleGrid[r][c-1]
			} else if c == 0 {
				obstacleGrid[r][c] = obstacleGrid[r-1][c]
			} else {
				obstacleGrid[r][c] = obstacleGrid[r-1][c] + obstacleGrid[r][c-1]
			}

		}

		fmt.Println(obstacleGrid[m-1][n-1])
	}
	fmt.Println(obstacleGrid)
	return true

}
