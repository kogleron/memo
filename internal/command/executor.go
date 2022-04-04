package command

type Executor interface {
	Supports(cmd Command) bool
	Run(cmd Command) error
}
