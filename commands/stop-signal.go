package commands

type StopSignalCommand LineStep

func (c StopSignalCommand) Validate() (bool, *LintError) {
	return true, nil
}
