package main

import (
	"fmt"
	"github.com/JESSE-SOTERIA/gopack/cmd/cmd"
	"github.com/JESSE-SOTERIA/gopack/graph"
	"github.com/JESSE-SOTERIA/gopack/parse"
	"log"
	"os"
)

var message string

//make a dependency graph data structure
//recursivelytraverse the dependency tree to identify all the imported modules
//Add transformations to the modules transpiling ES6+ code to ES5 using Babel
//implement loaders
//optional: optimize code by tree shaking
//concatenate the transformed code into one or more js files depending on the input
//ensure that the correct order of modules is maintained to satisfy dependencies
//optional: apply optimizations such as scope hoistin, minification
//write the generated bundles to the specified output directory
//optional: generatae the source maps to hep debugging by mapping the code in the bundle to the oroginal source code

func main() {
	//find a way to terminate if onw of the conditions in the root command is not met
	//example: if the last arg is not a valid file path
	//if one of the entry points is not a valid file

	cmd.RootCommand.Flags().StringSliceVarP(&cmd.EntryFiles, "entry", "e", []string{}, "list of entry files")
	cmd.RootCommand.Flags().StringVarP(&cmd.OutputPath, "output", "o", "", "output path for the bundle")
	cmd.RootCommand.Flags().StringVarP(&cmd.RootDependency, "roodDep", "r", "", "root directory from where all your local dependencies can be reached")

	//run bundle --help to see what the app does
	if err := cmd.RootCommand.Execute(); err != nil {
		fmt.Println("an error occured!")
		os.Exit(1)
	}

	//this graph only has vertices and nodeIterator initialized
	projectGraph := graph.InitializeGraph(cmd.EntryFiles)

	//adds the Edges and the adjacencyList to the graph.
	//graph is now complete
	for _, val := range projectGraph.Vertices {
		list, err := parse.Parse(val.Name)
		if err != nil {
			//need to unwrap the error and handle it effectively
			log.Fatal(err)
		}

		var neighbourNodes []graph.Node
		var newEdges []graph.Edge
		for _, name := range list {
			newNode := projectGraph.NewNode(name)
			newNode.Id = projectGraph.IdGen.GetId()
			//make edge between each node and each neighbour
			newEdges = append(newEdges, projectGraph.NewEdge(val, newNode))
			neighbourNodes = append(neighbourNodes, newNode)
		}
		projectGraph.AdjacencyList[val] = neighbourNodes
		projectGraph.Edges = newEdges

	}

	//call topo, then transitive reduction
	exists, cycle := graph.DetectCycle(projectGraph)
	if exists {
		log.Fatal("Detected a cycle of dependencies", cycle)
	}
	entry := projectGraph.NewNode(cmd.EntryFiles[0])
	//order wil lbe used to generate the bundle and output
	order := graph.TopoSortDFS(projectGraph, entry)
	fmt.Println(order, "successfully built topo!!!")

}
