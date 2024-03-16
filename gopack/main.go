package main

import (
	"fmt"
	"github.com/JESSE-SOTERIA/gopack/cmd/cmd"
	"os"
)

func main() {
	fmt.Println("hello from main!")
	RootCommand := cmd.NewRootCommand()

	if err := RootCommand.Execute(); err != nil {
		fmt.Println("an error occured!")
		os.Exit(1)
	}
}
