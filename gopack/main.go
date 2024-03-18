package main

import (
	"fmt"
	"github.com/JESSE-SOTERIA/gopack/cmd/cmd"
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

	//run bundle --help to see what the app does
	if err := cmd.RootCommand.Execute(); err != nil {
		fmt.Println("an error occured!")
		os.Exit(1)
	}
	//cmd.Entry files stores the entry file names of the project
	//cmd.outputpath holds the path of the bundled output

	//loop through the entry points and parse the files
	//write the parsing logic in its own package

}
