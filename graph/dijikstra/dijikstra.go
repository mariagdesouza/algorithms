package main

import (
	"fmt"
	"math"
)

/*

O(|E|+VlogV)

Calculate the shortest path between one node and every other node in the graph
   B
  / |\
 /  | \
A  /   \
\ /
 C------D
Selected node
*/

type Vertex struct {
	value     string
	neighbors []*Vertex
}

type Graph struct {
	Visited map[string]struct{}
	V       *Vertex
}

func main() {

	/* Let us create the example graph discussed above */
	graph := [][]int{
		{0, 4, 0, 0, 0, 0, 0, 8, 0},
		{4, 0, 8, 0, 0, 0, 0, 11, 0},
		{0, 8, 0, 7, 0, 4, 0, 0, 2},
		{0, 0, 7, 0, 9, 14, 0, 0, 0},
		{0, 0, 0, 9, 0, 10, 0, 0, 0},
		{0, 0, 4, 14, 10, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 1, 6},
		{8, 11, 0, 0, 0, 0, 1, 0, 7},
		{0, 0, 2, 0, 0, 0, 6, 7, 0},
	}

	dijkstra(graph, 0)
}

const (
	V = 9
)

func dijkstra(graph [][]int, v int) {

	distance := make([]int, V)
	sptSet := make([]bool, V)

	for i := 0; i < V; i++ { // set all distance vertexes to infinity
		distance[i] = math.MaxInt32
	}

	// set 0th to 0 - that wil be the starting point
	distance[0] = 0
	for i := 0; i < V-1; i++ {
		u := minDistance(distance, sptSet) // find the next node to visit
		sptSet[u] = true                   // set as visited

		for v := 0; v < V; v++ {
			// Update dist[v] only if is not in sptSet, there is an edge from
			// u to v, and total weight of path from src to  v through u is
			// smaller than current value of dist[v]
			if !sptSet[v] && distance[u] != math.MaxInt32 && distance[u]+graph[u][v] < distance[v] {
				distance[v] = distance[u] + graph[u][v]
			}
		}
	}

	for i := 0; i < V; i++ {
		fmt.Printf("%d tt %d\n", i, distance[i])
	}
}

func minDistance(d []int, s []bool) int {

	min := math.MaxInt64
	for i, v := range d {
		if v < min && s[i] == false {
			min = v
		}
	}
	return min
}
