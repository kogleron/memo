package bootstrap

import (
	"memo/internal/command"
)

type commandExecutors []command.Executor

func NewCommandExecutors(randExec *command.RandExecutor, startExec *command.StartExecutor, defaultExec *command.DefaultCommandExecutor) commandExecutors {
	return commandExecutors{
		randExec,
		startExec,
		defaultExec,
	}
}
