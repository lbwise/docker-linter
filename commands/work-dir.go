package commands

type WorkDirCommand LineStep

func (c WorkDirCommand) Validate() (bool, *LintError) {
	return true, nil
}
