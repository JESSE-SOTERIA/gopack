package graph

import (
	//"fmt"
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
	graph *Graph
	Name  string
	Id    int64
}

// make sure you initialize each graph with an appropriate currentnode, maxnode, iterator value
type Graph struct {
	Edges         []Edge
	Vertices      map[int64]Node
	AdjacencyList map[Node][]Node
	IdGen         uniqueIdGenerator
	NodeIterator  Nodes
}

type Edge struct {
	graph  *Graph
	FromId int64
	ToId   int64
}

// this type and the iterator is based on the fact that node Id's will be assigned by a simple incrementing function say += 1 for the node id
// after running the topological sort, ill get back a slice of sorted vertices. which ill use the index of as the id later which would mean the sort
type Nodes struct {
	parentGraph *Graph
	currentNode int
	maxNode     int
}

type uniqueIdGenerator struct {
	graph   *Graph
	counter int64
}

func (u *uniqueIdGenerator) GetId() int64 {
	u.counter++
	return u.counter
}

// maybe have a function makeNodeIterator that returns a node iterator like the function does below?
func (g *Graph) Nodes() Nodes {
	return Nodes{
		parentGraph: g,
		currentNode: 0,
		maxNode:     len(g.Vertices),
	}
}

// complete, test
func InitializeGraph(list []string) Graph {
	var (
		newGraph Graph
		newIter  Nodes
	)

	//loop over the list
	for i := 0; i < len(list); i++ {
		newNode := Node{
			graph: &newGraph,
			Name:  list[i],
			Id:    newGraph.IdGen.GetId(),
		}
		newGraph.AddNode(newNode)
	}
	newGraph.NodeIterator, newIter.parentGraph = newIter, &newGraph

	newIter.maxNode = len(newIter.parentGraph.Vertices)
	return newGraph
}
