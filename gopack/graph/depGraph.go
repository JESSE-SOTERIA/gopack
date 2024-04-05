package graph

// THERE IS A DEFAULT NODE VALUE THAT MIGHT BE RETURNED BY METHODS THAT RETURN NODES, NEEDS TO BE HANDLED APPROPRIATELY BY CALLERS

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
