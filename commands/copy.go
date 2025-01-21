package commands

type CopyCommand LineStep

func (c CopyCommand) Validate() (bool, *LintError) {
	return true, nil
}
