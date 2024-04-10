package parse

import (
	"bufio"
	//"errors"
	"fmt"
	"github.com/JESSE-SOTERIA/gopack/cmd/cmd"
	"io/fs"
	"os"
	"os/exec"
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
	if strings.Contains(line, "require") || strings.Contains(line, "import") {
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

	//verify if rootDependency is a valid filepath in the users system
	_, rootErr := os.Stat(cmd.RootDependency)
	if rootErr != nil {
		wrapped := fmt.Errorf("Dependency Root does not Exist! %:w", rootErr)
		return "", wrapped
	}
	var foundPath string
	//walk down the file tree from cmd.RootDependency and look for a file named like the parameter, dependencyName
	walkErr := filepath.WalkDir(cmd.RootDependency, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			wrapped := fmt.Errorf("failed to look for file in directory %s: %w", path, err)
			return wrapped
		}
		if d.Name() == dependencyName {
			foundPath = path
			//signal that the file has been found
			return fs.SkipAll
		}
		return nil
	})

	if walkErr != nil && walkErr != fs.SkipAll {
		return "", walkErr
	}
	//****unwrap failed to look for file error and handle it appropriately

	if foundPath == "" {
		//this should trigger downloading dependencies using the package manager of choice
		fmt.Printf(dependencyName, "%s is not a local dependency.", dependencyName)
		//return "", errors.New("failed: invalid dependency")
		installWithNpm()
		return "installed with npm", nil
	}

	//combine the path of the dependency to the path of the root dependency directory

	absolutePath, err := filepath.Rel(cmd.RootDependency, foundPath)
	fmt.Println(absolutePath, err)
	if err != nil {
		wrapped := fmt.Errorf("cant make the filepath relative: %w", err)
		return "", wrapped
	}

	return absolutePath, nil
}

// to be used when intergrating online module resolution
func installWithNpm() bool {
	cmd := exec.Command("npm", "install")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("installling online dependencies...")
	err := cmd.Run()
	if err != nil {
		fmt.Println("there has been an error installing remote dependencies:", err)
		os.Exit(1)
	}
	return true

}
func installWithYarn() bool {
	cmd := exec.Command("yarn", "install")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("installling online dependencies...")
	err := cmd.Run()
	if err != nil {
		fmt.Println("there was an error installing remote dependencies")
		os.Exit(1)
	}
	return true
}

// parse reads files line by line, checks for dependencies, and resolves them then returns a slice of resolved dependencies in the form of parsed strings.
func Parse(file string) ([]string, error) {
	var parsedStrings []string
	buf, err := os.Open(file)
	if err != nil {
		wrapped := fmt.Errorf("failed to open %s: %w", file, err)
		return []string{}, wrapped
	}

	defer buf.Close()
	scanner := bufio.NewScanner(buf)

	for scanner.Scan() {
		dependencyName, ok := checkDependency(scanner.Text())
		var installedNpm bool
		if ok {
			continue
		}
		resolved, err := resolveLocalDependency(dependencyName)
		if err != nil {
			return []string{}, err
		}
		if resolved == "installed with npm" {
			installedNpm = true
			resolved, _ = resolveLocalDependency(dependencyName)
		}
		parsedStrings = append(parsedStrings, resolved)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error while scanning files:%w", err)
	}
	return parsedStrings, nil

}
