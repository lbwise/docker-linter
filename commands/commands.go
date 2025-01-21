package commands

import "errors"

type LineStep struct {
	Line      int
	Keyword   string
	Arguments []string
}

type Command interface {
	Validate() (bool, *LintError)
}

type LintError struct {
	Line  int
	Error error
}

var (
	NotEnoughArgumentsError = errors.New("not enough arguments provided for command")
)
