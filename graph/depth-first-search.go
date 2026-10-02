package graph

import "fmt"

// DepthFirstSearch : Start at root abd explore each branch completely before moving to the next branch
//DFS is recursive
func DepthFirstSearch(v *Vertex) {

	if v.visited { // return if visited
		return
	}
	//print node
	v.visited = true
	fmt.Println(v.value)
	for _, vert := range v.neighbors { //visit all neighbors recursively
		DepthFirstSearch(vert)
	}

}

func DepthFirstSearchVisited(v *Vertex, visited map[string]struct{}) {

	if v == nil {
		return
	}

	if _, ok := visited[v.value]; ok { // return if visited
		return
	}
	//mark as visited - print node
	visited[v.value] = struct{}{}
	fmt.Println(v.value)

	for _, vert := range v.neighbors { //visit all neighbors recursively
		DepthFirstSearchVisited(vert, visited)
	}

}
