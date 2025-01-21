package commands

type VolumeCommand LineStep

func (c VolumeCommand) Validate() (bool, *LintError) {
	return true, nil
}
