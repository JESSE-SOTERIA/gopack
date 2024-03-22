package graph

import (
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/iterator"
	"gonum.org/v1/gonum/simple"
	"gonum.org/v1/gonum/topo"
)

// implement graph, node, and edge, which are the building blocks of graphs
// implement all the methods to these types that will be needed for the desired functionality
// call the methods from the libraries.
// commit code.
// THERE IS A DEFAULT NODE VALUE THAT MIGHT BE RETURNED BY METHODS THAT RETURN NODES, NEEDS TO BE HANDLED APPROPRIATELY BY CALLERS
//implement the nodes method in the graph interface

type Node struct {
	Name string
	Id   int64
}

// make sure you initialize each graph with an appropriate currentnode, maxnode, iterator value
type Graph struct {
	Vertices map[int64]Node
	//node.name is the file names
	AdjacencyList map[Node][]Node
	//fields below this comment used to implement the iterator interface in the graph package
	//should be initialized to 0, valid numbers start from one, and zero denotes an empty adjacencylist
	currentNode int
	// should be initialized to the length of the adjacency list
	maxNode int
}

type Iterator struct {
	parentgraph *Graph
	currentNode int64
	maxNode     int
}

type Nodes struct {
	*Iterator
}

func (n *Node) ID() int64 {
	return n.Id
}

// make sure to check if the Id returned is (-1) which means the node doesn't exist and handle the case accordingly
func (g *Graph) Node(id int64) Node {
	var defaultNode = Node{
		Name: "default",
		Id:   -1,
	}
	node, exists := g.Vertices[id]

	if exists {
		return node
	}
	return defaultNode
}

//func (g *Graph) Nodes() Nodes{
//initializes an iterator which points the graph to which this method belongs
//initializes a value of type nodes with the pre initialized iterator
//
//}

func (i *Iterator) Next() bool {
	if i.parentgraph.currentNode == i.parentgraph.maxNode {
		return false
	}
	i.parentgraph.currentNode += 1
	if i.parentgraph.currentNode == i.parentgraph.maxNode {
		return false
	}

	return true
}
func (i *Iterator) Len() int {
	return len(i.parentgraph.AdjacencyList) - i.parentgraph.currentNode
}

func (i *Iterator) Reset() {
	i.parentgraph.currentNode -= len(i.parentgraph.AdjacencyList)
	if i.parentgraph.currentNode <= 0 {
		i.parentgraph.currentNode = 0
	}
}

func (n *Nodes) Node() Node {
	currentNode := n.Iterator.currentNode
	var defaultNode = Node{
		Name: "default",
		Id:   -1,
	}
	node, exists := n.Iterator.parentgraph.Vertices[currentNode]
	if exists {
		return node
	}
	return defaultNode
}
