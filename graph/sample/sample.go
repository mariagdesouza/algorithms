package main

import (
	"fmt"

	"github.com/mariadesouza/algorithms/graph"
)

func main() {

	a := []string{"A", "B", "C", "D", "E"}
	var g []*graph.Vertex
	for _, v := range a {
		g = append(g, graph.NewVertex(v))
	}

	g[0].AddEdge(g[1])
	g[1].AddEdge(g[2], g[4])
	g[2].AddEdge(g[3], g[4])

	fmt.Println("DepthFirstSearch")

	m := make(map[string]struct{})
	graph.DepthFirstSearchVisited(g[0], m)

	//graph.ResetVisitedAll(g...)
	fmt.Println("BreadthFirstSearch")
	graph.BreadthFirstSearch(g[0])

	graph.ResetVisitedAll(g...)
	fmt.Println("TopologicalSearch")
	graph.TopologicalSort(g[0])

	graph.ResetVisitedAll(g...)
	visited := make(map[string]struct{})
	var results []string
	results = graph.TopologicalSortVisited(g[0], visited, results)
	fmt.Println(results)

}
