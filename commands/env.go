package commands

type EnvCommand LineStep

func (c EnvCommand) Validate() (bool, *LintError) {
	return true, nil
}
