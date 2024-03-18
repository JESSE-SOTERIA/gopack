package parse

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Parse(files ...string) []string {
	var parsedFiles []string
	for _, file := range files {
		buf, err := os.Open(file)

		if err != nil {
			log.Fatal(err)
		}
		defer buf.Close()
		//reads each file (buf) line by line
		scanner := bufio.NewScanner(buf)

	}
	return parsedFiles
}

func checkDependency(line string) string {

}
