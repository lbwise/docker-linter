package commands

type CmdCommand LineStep

func (c CmdCommand) Validate() (bool, *LintError) {
	return true, nil
}
