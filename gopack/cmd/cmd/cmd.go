package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
	"strings"
)

var( EntryFiles []string
     OutputPath string
)


     var RootCommand = &cobra.Command{
		Use:   "bundle",
		Short: "start bundling your files",
		Long:  "gopack bundles your javascript projects for better performance of you application",
		Run: func(cmd *cobra.Command, args []string) {
			if validateFlags(EntryFiles, OutputPath){
				fmt.Println("bundling was a success!!!")
		}
		},
	}



func isValidArg(arg string) bool {
	dotIndex := strings.LastIndex(arg, ".")
	var extension string

	if dotIndex != -1 && dotIndex < len(arg) {
		extension = arg[dotIndex+1:]
	}

	switch extension {
	case "js":
		return true
	case "css":
		return true
	case "png":
		return true
	case "jpg":
		return true
	default:
		return false

	}

}

func validateFlags(entry []string, out string) bool{
	//check the lengthof the entry slice
	if len(entry) < 1{
		fmt.Println("please enter an entry file and and output path to use the bundler")
		return false
		}

		for i := 0; i < len(entry); i++ {
		    if isValidArg(entry[i]) {
				fmt.Printf("%s is a valid argument\n", entry[i])
			} else {
				fmt.Printf("[%s] is not a valid entry file\n", entry[i])
			 }
		}

			if filepath.IsAbs(out) {
				fmt.Println("bundling started successfully")
				return true
			} else {
				filepath.Abs(out)
				fmt.Println("converted to valid file path and bundling successfully!")
				return true
			}
		}
