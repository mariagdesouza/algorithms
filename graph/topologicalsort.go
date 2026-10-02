package graph

import "fmt"

//Topological Sort is the most important operation on directed acyclic graphs or DAGs.
// Topological Sort is useful in scheduling tasks where precedence ordering matters, i.e, one task needs to be done before starting another.
// Topological sort can sequence tasks while respecting all sequence constraints without any conflict.
/*
func TopologicalSort(v *Vertex, result []string) []string {
	v.visited = true
	for _, e := range v.neighbors {
		if !e.visited {
			result = TopologicalSort(e, result)
		}
	}
	result = append(result, v.value)
	return result
}
*/

func TopologicalSort(v *Vertex) {
	v.visited = true
	for _, e := range v.neighbors {
		if !e.visited {
			TopologicalSort(e)
		}
	}
	fmt.Println(v.value)
}

func TopologicalSortVisited(v *Vertex, visited map[string]struct{}, results []string) []string {

	visited[v.value] = struct{}{} // mark as visited

	for _, e := range v.neighbors {
		if _, ok := visited[e.value]; !ok {
			results = TopologicalSortVisited(e, visited, results)
		}
	}

	results = append(results, v.value)
	return results
}

/*
func TopologicalSortStack(v *Vertex) {

	var resultStack []string

	visited := make(map[string]struct{})

	for len(resultStack) {
		for _, e :=
	}

}*/
