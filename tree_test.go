package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestRootCommand(t *testing.T) {

	// test command
	cmd := exec.Command("go", "run", ".", "--no-color", "-l=3")

	// expected output
	lines := []string{
		`├── .github`,
		`|  └── workflows`,
		`|     ├── ci.yml`,
		`|     └── release.yml`,
		`├── LICENSE`,
		`├── README.md`,
		`├── assets`,
		`|  ├── demo.cast`,
		`|  ├── demo.gif`,
		`|  └── demo.gif.sh`,
		`├── go.mod`,
		`├── go.sum`,
		`├── list.go`,
		`├── tree.go`,
		`└── tree_test.go`,
	}

	// get output
	if output, err := cmd.Output(); err != nil {
		fmt.Println("Error: ", err)
	} else {
		outputString := fmt.Sprintf("%s", output)

		for _, line := range lines {
			if !strings.Contains(outputString, line) {
				t.Fail()
			}
		}
	}
}
