package graph

import (
	"fmt"
	"testing"
)

func TestInitializeGraph(t *testing.T) {
	var samplefiles []string = []string{"aroma.js", "is.js", "the.css", "best.js"}
	for _, i := range samplefiles {
		graph := InitializeGraph(samplefiles)
		fmt.Println(graph.Vertices, i)
	}
}
