package commands

type ExposeCommand LineStep

func (c ExposeCommand) Validate() (bool, *LintError) {
	return true, nil
}
