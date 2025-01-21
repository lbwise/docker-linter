package commands

type RunCommand LineStep

func (c RunCommand) Validate() (bool, *LintError) {
	return true, nil
}
