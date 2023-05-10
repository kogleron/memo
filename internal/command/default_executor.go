package command

type DefaultCommandExecutor struct {
	realExecutor Executor
}

func (e DefaultCommandExecutor) GetDescription() string {
	return ""
}

func (DefaultCommandExecutor) GetName() string {
	return ""
}

func (e *DefaultCommandExecutor) Supports(_ Command) bool {
	return true
}

func (e *DefaultCommandExecutor) Run(cmd Command) error {
	return e.realExecutor.Run(cmd)
}

func NewDefaultCommandExecutor(realExecutor Executor) *DefaultCommandExecutor {
	return &DefaultCommandExecutor{
		realExecutor: realExecutor,
	}
}
