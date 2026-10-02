// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

var (
	board = []string{
		"ABCD",
		"EFJK",
		"GHLM",
		"INOP",
	}
)

type wordList []string

func (w wordList) IsWord(s string) bool {
	for _, word := range w {
		if word == s {
			return true
		}
	}
	return false
}

func NewLexicon() Lexicon {
	return wordList([]string{
		"BEGIN", "GIN", "BEG",
		"MOP", "MONK", "MOM", "POP",
	})
}

type Node struct {
	Item      rune
	Neighbors []*Node
}

func NewNode(i rune) *Node {
	return &Node{Item: i, Neighbors: nil}
}

var (
	deltas = [][]int{
		{0, 1},
		{1, 0},
		{1, 1},
		{-1, 1},
		{0, -1},
		{-1, 0},
		{-1, -1},
		{1, -1},
	}
)

func MakeBoard(board []string) []*Node {
	var output []*Node
	var tmp [][]*Node
	for i, line := range board {
		tmp = append(tmp, nil)
		for _, ch := range line {
			node := &Node{Item: ch}
			tmp[i] = append(tmp[i], node)
			output = append(output, node)
		}
	}
	for i, row := range tmp {
		for j, item := range row {
			for _, d := range deltas {
				oi := i + d[0]
				oj := j + d[1]
				if oi < 0 || oi >= len(tmp) || oj < 0 || oj >= len(row) {
					continue
				}
				other := tmp[oi][oj]
				item.Neighbors = append(item.Neighbors, other)
			}
		}
	}
	return output
}

func findWords(node *Node, lexicon Lexicon, s string, visited map[*Node]struct{}) {

	if _, ok := visited[node]; ok {
		return
	}

	visited[node] = struct{}{}
	s += string(node.Item)
	if lexicon.IsWord(s) {
		fmt.Println(s)
	}
	for _, v := range node.Neighbors {
		findWords(v, lexicon, s, visited)
	}
	delete(visited, node)

}

type Lexicon interface {
	IsWord(string) bool
}

// aBcd
// Efjk
// Ghlm
// INop
func main() {
	nodes := MakeBoard(board)
	lexicon := NewLexicon()

	for i := 0; i < len(nodes); i++ {
		visited := make(map[*Node]struct{})
		findWords(nodes[i], lexicon, "", visited)
	}
}
