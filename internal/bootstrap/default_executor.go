package bootstrap

import "memo/internal/command"

func NewDefaultCommandExecutor(addExec *command.AddExecutor) *command.DefaultCommandExecutor {
	return command.NewDefaultCommandExecutor(addExec)
}
