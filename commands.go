package main

type Command interface {
	Validate() (bool, []LintError)
}

type LintError struct {
	Line  int
	Error error
}

type AddCommand LineStep

func (a *AddCommand) Validate() {
	return a.Keyword
}

type ArgCommand struct {
	Keyword   string
	Arguments []string
}

type CmdCommand struct {
	Keyword   string
	Arguments []string
}

type CopyCommand struct {
	Keyword   string
	Arguments []string
}

type EntrypointCommand struct {
	Keyword   string
	Arguments []string
}

type EnvCommand struct {
	Keyword   string
	Arguments []string
}

type ExposeCommand struct {
	Keyword   string
	Arguments []string
}

type FromCommand struct {
	Keyword   string
	Arguments []string
}

type HealthcheckCommand struct {
	Keyword   string
	Arguments []string
}

type LabelCommand struct {
	Keyword   string
	Arguments []string
}

type MaintainerCommand struct {
	Keyword   string
	Arguments []string
}

type OnbuildCommand struct {
	Keyword   string
	Arguments []string
}

type RunCommand struct {
	Keyword   string
	Arguments []string
}

type ShellCommand struct {
	Keyword   string
	Arguments []string
}

type StopsignalCommand struct {
	Keyword   string
	Arguments []string
}

type UserCommand struct {
	Keyword   string
	Arguments []string
}

type VolumeCommand struct {
	Keyword   string
	Arguments []string
}

type WorkdirCommand struct {
	Keyword   string
	Arguments []string
}
