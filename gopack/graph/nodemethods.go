package graph

import (
	"fmt"
)

// determines which nodes the iterator iterates
func (n *Nodes) Len() int {
	//maxNode should be set depending on the item being iterated on
	return n.maxNode - n.currentNode
}

func (n *Node) ID() int64 {
	return n.Id
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

func (n *Nodes) Next() bool {
	if n.currentNode == n.maxNode {
		return false
	}
	n.currentNode++
	return true
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
