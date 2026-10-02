package graph

// a tree is a type of graph but not all graphs are trees
//A tree us a connected graph without cycles
// Graphs can be directed or undirected.

//If there us a path between every pair of vertices, is called connected

type Vertex struct {
	value     string
	visited   bool
	neighbors []*Vertex
}

func NewVertex(val string) *Vertex {
	return &Vertex{value: val, visited: false, neighbors: nil}
}

func (n *Vertex) AddEdge(vertex ...*Vertex) {
	n.neighbors = append(n.neighbors, vertex...)
}

func (n *Vertex) ResetVisited() {
	n.visited = false
}

func ResetVisitedAll(vertex ...*Vertex) {
	for _, v := range vertex {
		v.visited = false
	}
}

