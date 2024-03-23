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
	Vertices      map[int64]Node
	AdjacencyList map[Node][]Node
	NodeIterator  Nodes
}

type Edge struct {
	from Node
	to   Node
}

// this type and the iterator is based on the fact that node Id's will be assigned by a simple incrementing function say += 1 for the node id
// after running the topological sort, ill get back a slice of sorted vertices. which ill use the index of as the id later which would mean the sort
type Nodes struct {
	parentGraph *Graph
	currentNode int
	maxNode     int
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

// is this a viable way to instantiate a graph??? or part of it???
// maybe have a function makeNodeIterator that returns a node iterator like the function does below?
func (g *Graph) Nodes() Nodes {
	return Nodes{
		parentGraph: g,
		currentNode: 0,
		maxNode:     len(g.Vertices),
	}
}

// the node doesnt exist or doesnt have any neighbours if the value of maxNode of the nodes returned is 0
func (g *Graph) From(id int64) Nodes {
	node, exists := g.Vertices[id]
	if exists {
		//neighbours is a list of nodes that border the id node that from is called with
		neighbours, _ := g.AdjacencyList[node]
		//return an iterator that iterates through a node slice
		return Nodes{
			parentGraph: g,
			currentNode: 0,
			maxNode:     len(neighbours),
		}

	}
	//maxNode value of 0 means that the node is non existent
	return Nodes{
		parentGraph: g,
		currentNode: 0,
		maxNode:     0,
	}

}

func (g *Graph) HasEdgeBetween(xid, yid int64) bool {

}

func (g *Graph) Edge(uid, yid int64) Edge {

}

func (g *Graph) HasEdgeFromTo(uid, vid int64) bool {

}

func (g *Graph) To(id int64) Nodes {

}

//THE FOLLOWING IMPLEMENT THE ITERATOR INTERFACE
//the iterator should be a generic type that iterates over a certain item
//code the methods such that it could iterate throug graph.vertices and []node graph.adjacencylist

func (n *Nodes) Next() bool {
	if n.currentNode == n.maxNode {
		return false
	}
	n.currentNode++
	return true
}

// determines which nodes the iterator iterates
func (n *Nodes) Len() int {
	//maxNode should be set depending on the item being iterated on
	return n.maxNode - n.currentNode
}

func (n *Nodes) Reset() {
	n.currentNode = 0
}

func (n *Nodes) Node() Node {
	currentNode := int64(n.currentNode)
	var defaultNode = Node{
		Name: "default",
		Id:   -1,
	}
	node, exists := n.parentGraph.Vertices[currentNode]

	if exists {
		return node
	}
	return defaultNode
}
func (e *Edge) ReversedEdge() Edge {

}
func (e *Edge) From() Node {

}

func (e *Edge) To() Node {

}

func InitializeGraph() Graph {

}
