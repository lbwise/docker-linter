package commands

type HealthCheckCommand LineStep

func (c HealthCheckCommand) Validate() (bool, *LintError) {
	return true, nil
}
