package bootstrap

import (
	"memo/internal/command"
	"memo/internal/telegram"
)

func NewCommandExecutors(
	addExecutor *command.AddExecutor,
	randExec *command.RandExecutor,
	startExec *command.StartExecutor,
	searchExec *command.SearchExecutor,
	defaultExec *command.DefaultCommandExecutor,
	deleteExec *command.DeleteExecutor,
	replier telegram.Replier,
) command.Executors {
	helpExecutor := command.NewHelpExecutor([]command.Executor{
		addExecutor,
		deleteExec,
		randExec,
		startExec,
		searchExec,
	}, replier)

	return command.Executors{
		addExecutor,
		deleteExec,
		randExec,
		startExec,
		searchExec,
		helpExecutor,
		defaultExec,
	}
}
