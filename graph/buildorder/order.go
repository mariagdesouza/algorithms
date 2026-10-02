package main

import "fmt"

/*

Topological Sort

sorts a directed acyclic graph

INput is a list of dependencies


Lookup
*/

type Node struct {
	Item  string
	Edges []*Node
}

func NewNode(value string) *Node {
	return &Node{Item: value, Edges: nil}
}

func (n *Node) AddEdge(edges ...*Node) {
	for _, e := range edges {
		n.Edges = append(n.Edges, e)
	}
}

func topologicalSort(v *Node, visited map[string]struct{}, result []string) []string {
	visited[v.Item] = struct{}{}
	for _, e := range v.Edges {
		if _, ok := visited[e.Item]; !ok {
			result = topologicalSort(e, visited, result)
		}
	}
	result = append(result, v.Item)
	return result
}

//ordered - slice of ordered node items
//visited - map of visited items
//unresolved - map to track what we are working on so we dont encounter a cyclic ?? map or can it just be a variable

func FindBuildOrder(n *Node, ordered []string, visited map[string]struct{}, unresolved map[string]struct{}) []string {

	unresolved[n.Item] = struct{}{}

	// First Visit all edges
	for _, e := range n.Edges {
		if _, ok := visited[e.Item]; !ok { // !visited
			if _, seen := unresolved[e.Item]; seen {
				fmt.Println("Circular dependency error")
				return nil
			}
			ordered = FindBuildOrder(e, ordered, visited, unresolved)
		}
	}
	visited[n.Item] = struct{}{}
	delete(unresolved, n.Item)
	ordered = append(ordered, n.Item)
	return ordered
}

/*
   d
/    \
a->b->c
   \ /
    e

order -> d,  e, b, c,

*/

func main() {
	var g []*Node
	a := []string{"A", "B", "C", "D", "E"}
	for _, c := range a {
		g = append(g, NewNode(c))
	}

	fmt.Println(len(g))
	g[0].AddEdge(g[1])       // a depends on b
	g[0].AddEdge(g[3])       // a-> d
	g[1].AddEdge(g[2], g[4]) //b -> c     //b -> e
	g[2].AddEdge(g[3], g[4]) //c -> d     // c-> e
	//g[4].AddEdge(g[1])       // circular dependency

	var result []string

	lookup := make(map[string]struct{})
	unresolved := make(map[string]struct{})
	result = FindBuildOrder(g[0], result, lookup, unresolved)
	fmt.Println("Build Order:")
	for _, s := range result {
		fmt.Print(" ", s)
	}
	fmt.Println()
	var result1 []string
	lookup1 := make(map[string]struct{})
	result1 = topologicalSort(g[0], lookup1, result1)

	fmt.Println("Build Order:")
	for _, s := range result {
		fmt.Print(" ", s)
	}
	fmt.Println()
}
