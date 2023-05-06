package bootstrap

import (
	"memo/internal/command"
)

type commandExecutors []command.Executor

func NewCommandExecutors(
	randExec *command.RandExecutor,
	startExec *command.StartExecutor,
	searchExec *command.SearchExecutor,
	defaultExec *command.DefaultCommandExecutor,
	deleteExec *command.DeleteExecutor,
) commandExecutors {
	return commandExecutors{
		deleteExec,
		randExec,
		startExec,
		searchExec,
		defaultExec,
	}
}
