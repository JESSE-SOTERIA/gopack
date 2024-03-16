package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func NewRootCommand() cobra.Command {

	RootCommand := cobra.Command{
		Use:   "bundle",
		Short: "start bundling your files",
		Long:  "this command starts the bundling process of your project starting from the file you provided as an entry point and the bundled result is stored in the output path withuut modifying the original project so you can use either.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("root command is running!")
			//functionality for checking whether an argument is valid
			if isValidArg(args[0]) && isValidArg(args[1]) {
				fmt.Println("bundling in progress... please wait.....")
			}
		},
		Args: cobra.MinimumNArgs(2),
	}

	return RootCommand

}

func isValidArg(arg string) bool {
	dotIndex := strings.LastIndex(arg, ".")
	var extension string

	if dotIndex != -1 && dotIndex < len(arg) {
		extension = arg[dotIndex+1:]
	}

	switch extension {
	case "js":
		fmt.Println("File is a Javascript file")
		return true
	case "css":
		fmt.Println("File is a css file")
		return true
	case "png":
		fmt.Println("File is a png")
		return true
	case "jpg":
		fmt.Println("File is a jpg")
		return true
	default:
		fmt.Println("This is not a valid file for the bundler!")
		return false

	}

}
