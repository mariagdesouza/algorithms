package main

/*
2 players

X O

3 x 3

win = all X's row
	  all X's column
	  all X's diagnola

best move start in the middle

2 Xs one O dont win


Assuming a 3x3 board, there are 3raise to 9 combinations, 20,000 boards

each space : E, X, O

O "" O
X X  ""
O X  O

X X X X
O O E E
X O X O
E E X O

*/

import (
	"fmt"
)

//states
const (
	E = iota
	O
	X
)

const (
	N = 4
)

//Point : point on board
type Point struct {
	Row, Col, Player int
}

func main() {

	fmt.Println(abs(-5))
	board := [][]int{
		{O, E, O, X},
		{X, X, E, O},
		{O, X, O, E},
		{E, X, O, E},
	}
	lastMove := Point{Row: 0, Col: 1, Player: X}

	addMove(board, &lastMove)

	fmt.Println(hasWon(board, &lastMove))

}

//0,1, 0,2
//1,0, 1,3
//2,0  2,3

//0,3 1,2

func addMove(board [][]int, lastMove *Point) {
	board[lastMove.Row][lastMove.Col] = lastMove.Player
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func hasWonRow(board [][]int, lastMove *Point) bool {
	for i := 0; i < N; i++ {
		if board[lastMove.Row][i] != lastMove.Player {
			return false
		}
	}
	return true
}

func hasWonCol(board [][]int, lastMove *Point) bool {
	for i := 0; i < N; i++ {
		if board[i][lastMove.Col] != lastMove.Player {
			return false
		}
	}
	return true
}

func hasWon(board [][]int, lastMove *Point) bool {

	//
	//hasWonRow
	if hasWonRow(board, lastMove) {
		return true
	}

	//hasWonCol
	if hasWonCol(board, lastMove) {
		return true
	}

	// check if we need to check diagonal

	//0,0 1,1 2,2
	if lastMove.Row == lastMove.Col && hasWonDiagonal(board, 1, lastMove.Player) == true {
		return true
	}

	//0,2 1,1 2,0
	if lastMove.Row == N-lastMove.Col-1 && hasWonDiagonal(board, -1, lastMove.Player) == true {
		return true
	}

	return false
}

func hasWonDiagonal(board [][]int, direction int, player int) bool {

	col := 0
	if direction == -1 {
		col = 2
	}
	for i := 0; i < 3; i++ {
		if board[i][col] != player {
			return false
		}
		col = col + direction
	}
	return true
}
