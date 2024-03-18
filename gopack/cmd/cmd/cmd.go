package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"log"
	//"os"
	"path/filepath"
	"strings"
)

var (
	EntryFiles []string
	OutputPath string
)

var RootCommand = &cobra.Command{
	Use:   "bundle",
	Short: "start bundling your files",
	Long:  "gopack bundles your javascript projects for better performance of you application",
	Run: func(cmd *cobra.Command, args []string) {
		//check if the entry files and outputpath strings are of the required format
		//append output path to the current working directory if it's not an absolute path
		if validateFlags(EntryFiles, OutputPath) {
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

func validateFlags(entry []string, out string) bool {
	//check the lengthof the entry slice
	if len(entry) < 1 {
		fmt.Println("please enter an entry file and and output path to use the bundler")
		return false
	}

	for i := 0; i < len(entry); i++ {
		if isValidArg(entry[i]) {
			if fileExists(entry[i]) {
				fmt.Printf("%s is a valid file\n", entry[i])
			} else {
				fmt.Printf("%s does not exist in the current directory\n", entry[i])
				//uncomment the next line for intended functionality
				//os.Exit(1)
			}
		} else {
			fmt.Printf("[%s] is not a valid file\n", entry[i])
		}
	}

	if filepath.IsAbs(out) {
		fmt.Println("bundling started successfully")
		return true
	} else {
		absOutput, err := filepath.Abs(out)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("converted to valid file path and bundling successfully!")
		OutputPath = absOutput
		return true
	}
}

// checks whether the file exists within the current directory
func fileExists(fileName string) bool {
	var exists bool
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Name() == fileName {
			exists = true
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
		//end the task
	}
	return exists
}
