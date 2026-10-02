package main

import "fmt"

type Vertex struct {
	Item      int
	Neighbors []*Vertex
}

func NewVertex(val int) *Vertex {
	return &Vertex{Item: val, Neighbors: nil}
}

func (n *Vertex) AddEdge(vertex ...*Vertex) {
	n.Neighbors = append(n.Neighbors, vertex...)
}

// cycle in a undirected graph
func isCyclic(v *Vertex, visited map[int]struct{}, parent *Vertex) bool {

	//mark as visited - print node
	visited[v.Item] = struct{}{}

	for _, e := range v.Neighbors {

		if _, ok := visited[e.Item]; !ok {
			return isCyclic(e, visited, v)
		} else if parent != nil && e != parent {
			return true
		}
	}

	return false
}

func main() {
	visited := make(map[int]struct{})
	a := []int{0, 1, 2, 3, 4}
	var g []*Vertex
	for _, v := range a {
		g = append(g, NewVertex(v))
	}

	g[0].AddEdge(g[1])
	g[1].AddEdge(g[2], g[4])
	g[2].AddEdge(g[3], g[4])
	g[4].AddEdge(g[2])

	var foundCycle bool
	for _, v := range g {
		//if _, ok := visited[v.Item]; !ok {
		//fmt.Println(v.Item)
		if isCyclic(v, visited, nil) {
			foundCycle = true
			break
		}
		//}
	}
	fmt.Println(foundCycle)
}
