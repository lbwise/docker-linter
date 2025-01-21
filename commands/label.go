package commands

type LabelCommand LineStep

func (c LabelCommand) Validate() (bool, *LintError) {
	return true, nil
}
