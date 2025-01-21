package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	KEYWORD = iota
	STRING
)

var (
	NotEnoughArgumentsError = fmt.Errorf("not enough arguments provided in step")
)

type LineStep struct {
	Keyword   string
	Arguments []string
}

func main() {
	file, err := os.ReadFile("Dockerfile")
	if err != nil {
		panic(err)
	}

	// Parse and assemble Dockerfile commands
	stepStrings := strings.Split(string(file), "\n")
	var fileSteps []*LineStep
	for _, stepString := range stepStrings {
		if isEmptyOrComment(stepString) {
			continue
		}

		cmd, err := createNewStep(stepString)
		// TODO: Add error directory for linting
		if err != nil {
			fmt.Println(fmt.Errorf("error creating command: %v", err))
		} else {
			fileSteps = append(fileSteps, cmd)
		}
	}

	// Print out the commands
	fmt.Printf("%d Commands found\n", len(fileSteps))
	for _, step := range fileSteps {
		fmt.Println(step.Keyword, step.Arguments)
	}
}

func isEmptyOrComment(s string) bool {
	return strings.Trim(s, "\n\t\r") == "" || strings.HasPrefix(s, "#")
}

func createNewStep(cmdString string) (*LineStep, error) {
	vals := strings.Split(cmdString, " ")
	if len(vals) < 2 {
		return nil, NotEnoughArgumentsError
	}

	cmd := &LineStep{
		Keyword:   vals[0],
		Arguments: vals[1:],
	}
	return cmd, nil
}
