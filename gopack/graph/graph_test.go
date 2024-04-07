package graph

import (
	"fmt"
	"testing"
)

type entryFiles struct {
	file1 string
	file2 string
	file3 string
	file4 string
}

func TestInitializeGraph(t *testing.T) {
	samplefiles := []string{"one.js", "two.js", "three.js", "four.js", "five.js"}
	graph := InitializeGraph(samplefiles)
	fmt.Println(graph.Vertices)
}
