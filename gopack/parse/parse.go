package parse

import (
	"bufio"
	"fmt"
	"github.com/JESSE-SOTERIA/gopack/cmd/cmd"
	"os"
	"path/filepath"
	"path/kr/fs"
	"strings"
)

//Parse function opens each file in the entry points slice and reads them line by line
//check dependency function returns the file name of the module  imported by an import statement if it exists in the line its called with
//resolvedependency func takes the import string and returns an absolute path to that file
//resolved dependencies are appended to a slice which is then returned by the parse func

//maybe make the returned slice a map that maps every dependency to a file(DEPENDENCY GRAPH)
//we can have a package (dependency graph), and all its methods in a seperate package and import it here
//add slices to that graph based on the computation of the parse func

// checks for dependency statements in the line provided and returns the string in quotes from that line (the dependency)
// the boolean will later be used in the logic for making the dependency graph to determine whether the added node is the last one ( leaf of a tree)
func checkDependency(line string) (string, bool) {

	if strings.Contains(line, "import") {
		start := strings.Index(line, `"`)
		if start == -1 {
			return "", false
		}
		end := strings.Index(line[start+1:], `"`)
		if end == -1 {
			return "", false
		}
		return line[start+1 : start+1+end], true
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
func resolveDependency(dependencyName string) (string, error) {

	var foundPath string
	err := filepath.WalkDir(cmd.RootDependency, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			wrapped := fmt.Errorf("failed to walk the file path: %w", err)
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

	if foundPath == "" {
		return "", fmt.Errorf("dependency %s not found", dependencyName)
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
		resolved, err := resolveDependency(dependencyName)
		if err != nil {
			return []string{}, err
		}
		parsedStrings = append(parsedStrings, resolved)
	}
	return parsedStrings, nil

}
