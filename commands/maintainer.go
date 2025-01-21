package commands

type MaintainerCommand LineStep

func (c MaintainerCommand) Validate() (bool, *LintError) {
	return true, nil
}
