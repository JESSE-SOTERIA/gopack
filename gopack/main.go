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

//parse the entry points
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
	//find a way to terminate if one of the conditions in the root command is not met
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

	//loop over the projectgraph and call Parse for each entry
	for _, val := range projectGraph.Vertices {
		list, err := parse.Parse(val.Name)
		if err != nil {
			//need to unwrap the error and handle it effectively
			log.Fatal(err)
		}

	}
	//call newNode on all list Elements to get a list of nodes
	//use val as key and the node lista as a value to make the map that is projectGraph.adjaceccylist
	//add edges to the graph depending on the list of dependencies on the adjacencyList
	//call topo, then transitive reduction
}
