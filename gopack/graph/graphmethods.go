package graph

// refactor and test
func (g *Graph) HasEdgeBetween(xid, yid int64) bool {
	from, exists := g.Vertices[xid]
	if exists {
		//handle the case where if from node doesn't exist
		neighbours, _ := g.AdjacencyList[from]
		if len(neighbours) > 0 {
			for neighbour := 0; neighbour < len(neighbours); neighbour++ {
				if neighbours[neighbour].Id == yid {
					return true
				}
			}
			return false
		}
		return false
	}
	return false
}

func (g *Graph) Edge(uid, yid int64) Edge {
	node, exists := g.Vertices[uid]
	node2, exists2 := g.Vertices[yid]

	if exists && exists2 {
		return g.NewEdge(node, node2)
	}
	return Edge{
		FromId: -1,
		ToId:   -1,
	}
}

func (g *Graph) HasEdgeFromTo(uid, vid int64) bool {
	for i := 0; i < len(g.Edges); i++ {
		if g.Edges[i].FromId == uid && g.Edges[i].ToId == vid {
			return true
		}
	}
	return false

}

func (g *Graph) To(id int64) Nodes {
	for _, list := range g.AdjacencyList {
		for i := 0; i < len(list); i++ {
			if list[i].Id == id {
				return Nodes{
					parentGraph: g,
					currentNode: 0,
					maxNode:     0,
				} //it exists
			}
		}
	}
	//it doesn't
	return Nodes{
		parentGraph: nil,
		currentNode: -1,
		maxNode:     -1,
	}

}

// the NewEdge method adds an edge to the graph and then returns that edge for processing
func (g *Graph) NewEdge(from, to Node) Edge {
	var newEdge Edge

	newEdge.FromId = from.Id
	newEdge.ToId = to.Id
	g.Edges = append(g.Edges, newEdge)
	//UPDATE THE ADJACENCY LIST

	return newEdge
}

// adds an edge to the edges slice in the graph
func (g *Graph) SetEdge(e Edge) {
	g.Edges = append(g.Edges, e)

}

// the following implement a NodeAdder interface
// makes a new node with an arbitrary id
func (g *Graph) NewNode() Node {
	var newNode Node
	newNode.Name = ""
	//(-1) id means its a default node
	newNode.Id = -1
	return newNode
}

// vertices field of the graph type is a map that maps a node id to a single node in the map
func (g *Graph) AddNode(vertex Node) {
	g.Vertices[vertex.Id] = vertex
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
