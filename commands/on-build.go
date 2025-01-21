package commands

type OnBuildCommand LineStep

func (c OnBuildCommand) Validate() (bool, *LintError) {
	return true, nil
}
