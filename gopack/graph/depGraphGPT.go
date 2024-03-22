package graph

import (
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

// DirectedGraph implements the graph.Directed and graph.NodeAdder interfaces.
type DirectedGraph struct {
	*simple.DirectedGraph
}

// NewDirectedGraph creates a new DirectedGraph.
func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{simple.NewDirectedGraph()}
}

// AddNode adds a node to the graph.
func (g *DirectedGraph) AddNode(n graph.Node) {
	g.DirectedGraph.AddNode(n)
}

// AddEdge adds an edge to the graph.
func (g *DirectedGraph) AddEdge(from, to graph.Node) {
	g.DirectedGraph.SetEdge(g.NewEdge(from, to))
}

func main() {
	// Create a new directed graph
	g := NewDirectedGraph()

	// Add nodes to the graph
	node1 := g.NewNode()
	node2 := g.NewNode()
	node3 := g.NewNode()

	// Add the nodes to the graph
	g.AddNode(node1)
	g.AddNode(node2)
	g.AddNode(node3)

	// Add edges to the graph
	g.AddEdge(node1, node2)
	g.AddEdge(node2, node3)

	// Perform topological sort
	order, err := topo.Sort(g)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the topological order
	fmt.Println("Topological Order:")
	for _, node := range order {
		fmt.Printf("%d ", node.ID())
	}
	fmt.Println()
}
