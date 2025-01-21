package commands

type EntryPointCommand LineStep

func (c EntryPointCommand) Validate() (bool, *LintError) {
	return true, nil
}
