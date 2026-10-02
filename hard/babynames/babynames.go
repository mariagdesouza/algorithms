package main

import "fmt"

/*

Governament releasea 1000 most common baby names and their frequencies



Sometimes Same name have multiple spellings

E.g. john Jon

Given two lists:

Name Frequency
John   10
Jon   3
Davis   2
Johnny 11
Carlton 8
Carleton 2
Carrie 5
Kari 3


Name Alternate
Jonathan  John
Jon     Johnny
Johnny  John
Carleton Carlton
Kari    Carrie


Print a final list of names and their actual frequency


JOhn 34 Davis 2 Calrton 10


Jonathan -> john
Jon -> Johnny -> John

*/

func main() {
	frequencies := map[string]int{"John": 10, "Jon": 3, "Davis": 2, "Johnny": 11, "Carlton": 8, "Carleton": 2}
	names := map[string]string{"Jonathan": "John", "Jon": "Johnny", "Johnny": "John", "Carleton": "Carlton", "Kari": "Carrie"}
	g := ConstructGraph(names, frequencies)
	g.PrintFrequencies()
}

type Node struct {
	Name      string
	Frequency int
	Visited   bool
	Neighbors []*Node
}

func NewNode(name string, frequency int) *Node {
	return &Node{Name: name, Frequency: frequency, Visited: false, Neighbors: nil}
}

func (n *Node) AddNeighbor(e *Node) {
	n.Neighbors = append(n.Neighbors, e)
}

type Graph []*Node

func DepthFirstSearch(n *Node, result *int) {
	if n == nil || n.Visited == true {
		return
	}

	n.Visited = true
	*result += n.Frequency
	for _, e := range n.Neighbors {
		DepthFirstSearch(e, result)
	}
}

func (g *Graph) PrintFrequencies() {
	for _, n := range *g { // look in disconnected graphs
		var result int
		DepthFirstSearch(n, &result)
		fmt.Println(n.Name, ":", result)
	}
}

func (g *Graph) AddToGraph(k string, v string, frequencies map[string]int) {

	for _, n := range *g { // look in disconnected graphs
		var s *Node
		if n.Name == k {
			s = NewNode(v, frequencies[v])
		} else if n.Name == v {
			s = NewNode(k, frequencies[k])
		} else {
			for _, p := range n.Neighbors {
				if p.Name == k {
					s = NewNode(v, frequencies[v])
				} else if p.Name == v {
					s = NewNode(k, frequencies[k])
				}
				if s != nil {
					p.AddNeighbor(s)
					return
				}
			}
		}
		if s != nil {
			n.AddNeighbor(s)
			return
		}
	}
	//disconnected
	n1 := NewNode(v, frequencies[v])
	n2 := NewNode(k, frequencies[k])
	fmt.Println(n1.Name, "->", n2.Name)
	n1.AddNeighbor(n2)
	*g = append(*g, n1)
}

func ConstructGraph(names map[string]string, frequencies map[string]int) *Graph {
	var g Graph
	for k, v := range names { //jonathan, john
		g.AddToGraph(k, v, frequencies)
	}
	return &g
}
