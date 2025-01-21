package commands

type ArgCommand LineStep

func (c ArgCommand) Validate() (bool, *LintError) {
	return true, nil
}
