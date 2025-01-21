package commands

type UserCommand LineStep

func (c UserCommand) Validate() (bool, *LintError) {
	return true, nil
}
