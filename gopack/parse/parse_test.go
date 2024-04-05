package parse

import (
	"testing"
)

//package parse provides utilities for getting dependency names from files.
//it does this by reading each file in the entryfiles variable line by line
//and checking if a line has an import statement.

// supports import and require statements
// TODO: make sure to test for require module system
// find out how to test resolvedependency and parse
// make sure to include the fs package in parse
func TestCheckDependency(t *testing.T) {
	var tests = []struct {
		input string
		want  string
		ok    bool
	}{
		{"import defaultExport from \"module-name\";", "module-name", true},
		{"import * as name from \"module-name\";", "module-name", true},
		{"import { export1 } from \"module-name\";", "module-name", true},
		{"import { export1 as alias1 } from \"module-name\";", "module-name", true},
		{"import { default as alias } from \"module-name\";", "module-name", true},
		{"import { export1, export2 } from \"module-name\";", "module-name", true},
		{"import { export1, export2 as alias2, /* … */ } from \"module-name\";", "module-name", true},
		{"import { \"string name\" as alias } from \"module-name\";", "module-name", true},
		{"import defaultExport, { export1, /* … */ } from \"module-name\";", "module-name", true},
		{"import defaultExport, * as name from \"module-name\";", "module-name", true},
		{"import \"module-name\";", "module-name", true},
	}

	for _, sample := range tests {
		if got, _ := checkDependency(sample.input); got != sample.want {
			t.Errorf("checkDependency(%s) = %s", sample.input, got)
		}
	}
}

func TestResolveLocalDependency(t *testing.T) {

}
