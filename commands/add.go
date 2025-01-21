package commands

type AddCommand LineStep

func (a AddCommand) Validate() (bool, *LintError) {
	if len(a.Arguments) < 2 {
		return false, &LintError{Line: a.Line, Error: NotEnoughArgumentsError}
	}
	return true, nil
}
