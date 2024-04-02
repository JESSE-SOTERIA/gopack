package graph

//package to perform topological sort

// Detects a cycle in a directed graph and returns whether it exists or not
// and the slice of nodes showing the cycle of dependencies
// If there are no cycles, then the slice should be empty and the boolean should signal the continuation of the topological ordering of the graph in question
func DetectCycle(graph Graph) (bool, []Node) {
	//keeps track of visited nodes
	visited := make(map[Node]bool)
	//detects cycles
	recStack := make(map[Node]bool)
	//contains the cycle if any
	var cycle []Node

	for _, vertex := range graph.Vertices {
		if detectCycleDFS(graph, vertex, visited, recStack, &cycle) {
			return true, cycle
		}
	}

	//if there aren't any cycles detected
	return false, cycle
}

func detectCycleDFS(graph Graph, vertex Node, visited, recStack map[Node]bool, cycle *[]Node) bool {
	if !visited[vertex] {
		visited[vertex] = true
		recStack[vertex] = true

		for _, neighbor := range graph.AdjacencyList[vertex] {
			if !visited[neighbor] && detectCycleDFS(graph, neighbor, visited, recStack, cycle) {
				return true
			} else if recStack[neighbor] {
				// Cycle detected
				*cycle = append(*cycle, neighbor)
				for v := vertex; v != neighbor; v = (*cycle)[len(*cycle)-2] {
					*cycle = append(*cycle, v)
				}
				*cycle = append(*cycle, neighbor) // Close the cycle
				return true
			}
		}
	}
	recStack[vertex] = false // Remove vertex from recursion stack
	return false
}

//topological sort of the graph returns a slice of Nodes in their topological ordering
//maybe invert the order map to be a map of int to Node???
//for better ergonomics
func TopoSortDFS(graph Graph, start Node) map[Node]int {
	visited := make(map[Node]bool)
	order := make(map[Node]int)
	curLabel := len(graph.Vertices)

	var dfs func(node Node)
	dfs = func(node Node) {
		visited[node] = true
		for _, neighbor := range graph.AdjacencyList[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
		order[node] = curLabel
		curLabel--
	}

	dfs(start)
	return order
}

func TransitiveReduce() {

}
