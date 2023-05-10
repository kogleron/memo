package command

type Executor interface {
	GetName() string
	GetDescription() string
	Supports(cmd Command) bool
	Run(cmd Command) error
}
