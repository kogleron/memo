package bootstrap

import "memo/internal/api/telegram/command"

func NewDefaultCommandExecutor(addExec *command.AddExecutor) *command.DefaultCommandExecutor {
	return command.NewDefaultCommandExecutor(addExec)
}
