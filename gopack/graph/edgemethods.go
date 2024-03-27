package graph

import (
	"fmt"
)

// the following methods implement the Edge interface
func (e *Edge) ReversedEdge() Edge {
	var reversedEdge Edge
	reversedEdge.FromId = e.ToId
	reversedEdge.ToId = e.FromId

	return reversedEdge
}

// node is a newNode with a similar id : should be handled properly
func (e *Edge) From() Node {
	var newNode Node
	newNode.Id = e.FromId
	return newNode
}

// node is a newNode with a similar Id : should be handled appropriately
func (e *Edge) To() Node {
	var newNode Node
	newNode.Id = e.ToId
	return newNode
}
