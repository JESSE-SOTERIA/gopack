package graph

// the following methods implement the Edge interface
func (e *Edge) ReversedEdge() Edge {
	var reversedEdge Edge
	reversedEdge.FromId, reversedEdge.ToId = e.ToId, e.FromId
	return reversedEdge
}

// node is a newNode with a similar id : should be handled properly
//should be the original node from the graph
func (e *Edge) From() Node {
	var newNode Node
	newNode.Id = e.FromId
	return newNode
}

// node is a newNode with a similar Id : should be handled appropriately
func (e *Edge) To() Node {
	node, _ := e.graph.Vertices[e.ToId]
	return node
}
