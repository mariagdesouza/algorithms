package graph

import "fmt"

// BreadthFirstSearch : start at root and explore each neighbor before going to the children
// USED TO find shortest path between any two nodes
// BFS is NOT recursive
//It uses a queue
//O(E+V) => k is no of nodes d is depth
func BreadthFirstSearch(v *Vertex) {

	var q []*Vertex // queue of nodes
	visited := make(map[string]struct{})

	visited[v.value] = struct{}{} // mark visited
	q = append(q, v)              //enqueue

	for len(q) > 0 {

		vert := q[0] //front of queue
		fmt.Println(vert.value)
		q = q[1:] //dequeue

		for _, i := range vert.neighbors {
			if _, ok := visited[i.value]; !ok { //not visited
				visited[i.value] = struct{}{} // mark as visited
				q = append(q, i)              //enqueue
			}
		}
	}
}
