package main

import (
	"fmt"
	"github.com/lbwise/docker-linter/commands"
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

func main() {
	file, err := os.ReadFile("Dockerfile")
	if err != nil {
		panic(err)
	}

	// Parse and assemble Dockerfile commands
	stepStrings := strings.Split(string(file), "\n")
	var fileSteps []*commands.LineStep
	for lineNo, stepString := range stepStrings {
		if isEmptyOrComment(stepString) {
			continue
		}

		// Create steps with the keyword and arguments from the line
		stp, err := createNewStep(stepString, lineNo)
		if err != nil {
			fmt.Println(fmt.Errorf("error creating command: %v", err))
		}

		// Create the command from the keyword
		// TODO: Fix the dereferencing here
		cmd, err := fetchCommandFromStep(*stp)
		if err != nil {
			fmt.Println(fmt.Errorf("error fetching command: %v", err))
		}

		// Validate the semantics
		valid, err := cmd.Validate()
		if !valid {
			fmt.Println(err)
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

// Validates syntax of
func createNewStep(cmdString string, lineNo int) (*commands.LineStep, *commands.LintError) {
	vals := strings.Split(cmdString, " ")
	if len(vals) < 2 {
		return nil, nil
	}

	cmd := &commands.LineStep{
		Keyword:   vals[0],
		Arguments: vals[1:],
		Line:      lineNo,
	}
	return cmd, nil
}

func fetchCommandFromStep(stp commands.LineStep) (commands.Command, *commands.LintError) {
	switch stp.Keyword {
	case "ADD":
		return commands.AddCommand(stp), nil
	case "ARG":
		return commands.ArgCommand(stp), nil
	default:
		return nil, &commands.LintError{Line: 0, Error: fmt.Errorf("unknown command: %s", stp.Keyword)}
	}
}
