package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	EntryFiles []string
	OutputPath string
	//the root file from where the dependencies should be traced
	RootDependency string
)

//this package is to validate the entry files and output path

//call stack
//validateFlags
//isFileExists

var RootCommand = &cobra.Command{
	Use:   "bundle",
	Short: "start bundling your files",
	Long:  "gopack bundles your javascript projects for better performance of your web application you need to enter -e flag before each entry file for appropriate behaviour",
	Run: func(cmd *cobra.Command, args []string) {
		//check if the entry files and outputpath strings are of the required format
		//append output path to the current working directory if it's not an absolute path
		if validateFlags(EntryFiles, OutputPath) {
			fmt.Println("bundling was a success!!!")
		}
	},
}

// used this name because we are checking arguments to the  root command
func isValidArg(arg string) bool {
	dotIndex := strings.LastIndex(arg, ".")
	var extension string

	if dotIndex != -1 && dotIndex < len(arg) {
		extension = arg[dotIndex+1:]
	}

	//the file formats that gopach supports
	//more to be added with timett
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
		panic("program terminated because no entryfiles were provided!")
	}

	for i := 0; i < len(entry); i++ {
		if isValidArg(entry[i]) {
			if fileExists(entry[i], ".") {
				continue
			} else {
				fmt.Printf("%s does not exist in the current directory\n", entry[i])
				os.Exit(1)
			}
		} else {
			//should be an error
			fmt.Printf("[%s] is not a valid file\n", entry[i])
		}
	}

	if filepath.IsAbs(out) {
		//output will be within the cwd, with the extension that is outputpath
		//means that "/" is a valid filepath
		fmt.Println("bundling started successfully")
		return true
	} else {
		//check that the path is not a nonsensical text/ doesn't end with an extension eg .js
		// 1. check if the output ends with an extension
		valid := strings.LastIndex(out, ".")
		if valid != -1 {
			//replace with os formatted error
			log.Fatal(errors.New("output path cannot be a file"))
		}
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
// starting point is the directory level from where this checker will begin. the root of the file structure
func fileExists(fileName, startingPoint string) bool {
	var mu sync.Mutex
	var exists bool
	err := filepath.WalkDir(startingPoint, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		//check if the file exists in the current directory within the file path
		if d.Name() == fileName {
			mu.Lock()
			exists = true
			mu.Unlock()
		}

		return nil
	})

	if err != nil {
		//log because of the return value
		log.Fatal(err)
		//end the task
	}
	return exists
}
