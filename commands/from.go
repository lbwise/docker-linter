package commands

type FromCommand LineStep

func (c FromCommand) Validate() (bool, *LintError) {
	return true, nil
}
