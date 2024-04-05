package parse

import (
	"bufio"
	"fmt"
	"fs"
	"github.com/JESSE-SOTERIA/gopack/cmd/cmd"
	"os"
	"path/filepath"
	"strings"
)

// IMPORTANT: list of errors wrapped from this package:
// 1.failed to look for file:
// 2. failed to make the filepath relative:
// 3. failed to open file for dependency resolution:
// what if the file is an online asset?
func checkDependency(line string) (string, bool) {

	//checks the dependency of a file based on the import or require  string
	if strings.Contains(line, "import") {
		//checks if the line contains "
		//TODO: add support for single quote imports
		start := strings.Index(line, `"`)
		if start == -1 {
			return "", false
		}
		end := strings.Index(line[start+1:], `"`)
		//make sure there is no syntax error
		if end == -1 {
			return "", false
		}
		return line[start+1 : start+1+end], true
	}
	if strings.Contains(line, "require()") {
		//checks the contents of the requre
		//handles both single and double quotes
		start := strings.Index(line, `"`)
		startSingle := strings.Index(line, `'`)
		if start == -1 && startSingle == -1 {
			return "", false
		}
		if start != -1 {
			end := strings.Index(line[start+1:], `"`)
			if end == -1 {
				return "", false
			} else {
				return line[start+1 : start+1+end], true

			}
		} else {
			endSingle := strings.Index(line[startSingle+1:], `'`)
			if endSingle == -1 {
				return "", false
			} else {
				return line[startSingle+1 : startSingle+1+endSingle], true
			}
		}
	}
	return "", false
}

// implement own errror type for signaling the file has been found when resolving dependencies with the resolveDependency function
type FoundError struct{}

func (f *FoundError) Error() string {
	return "found file"
}

// takes a string (dependency file name) checks whether it exists in the root dependency directory and returns the absolute path to that file,
// or an error if that file does not exist
// only relosves files that exists within the users machine
// if it returns an error the calling function should unwrap the error and call the resolveOnlineDependency or abort funtion depending on the error
func resolveLocalDependency(dependencyName string) (string, error) {

	var foundPath string
	//walk down the file tree from cmd.RootDependency and look for a file named like the parameter, dependencyName
	err := filepath.WalkDir(cmd.RootDependency, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			wrapped := fmt.Errorf("failed to look for file: %w", err)
			return wrapped
		}
		if d.Name() == dependencyName {
			foundPath = path
			//signal that the file has been found
			return &FoundError{}
		}
		return nil
	})

	if err != nil && err.Error() != "found file" {
		return "", err
	}
	//****unwrap failed to look for file error and handle it appropriately

	if foundPath == "" {
		return "", fmt.Errorf("%s is not a local dependency", dependencyName)
	}

	absolutePath, err := filepath.Rel(cmd.RootDependency, foundPath)
	if err != nil {
		wrapped := fmt.Errorf("cant make the filepath relative: %w", err)
		return "", wrapped
	}

	return filepath.Join(cmd.RootDependency, absolutePath), nil
}

// parse reads files line by line, checks for dependencies, and resolves them then returns a slice of resolved dependencies in the form of parsed strings.
func Parse(file string) ([]string, error) {
	var parsedStrings []string
	buf, err := os.Open(file)
	if err != nil {
		wrapped := fmt.Errorf("failed to open %s: %w", err)
		return []string{}, wrapped
	}

	defer buf.Close()
	scanner := bufio.NewScanner(buf)

	for scanner.Scan() {
		dependencyName, ok := checkDependency(scanner.Text())
		if ok {
			continue
		}
		resolved, err := resolveLocalDependency(dependencyName)
		if err != nil {
			return []string{}, err
		}
		parsedStrings = append(parsedStrings, resolved)
	}
	return parsedStrings, nil

}
