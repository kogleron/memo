package command

import (
	"sort"
	"strings"

	"memo/internal/telegram"
)

func NewHelpExecutor(
	executors Executors,
	replier telegram.Replier,
) *HelpExecutor {
	return &HelpExecutor{
		executors: executors,
		replier:   replier,
	}
}

type HelpExecutor struct {
	executors []Executor
	replier   telegram.Replier
}

func (*HelpExecutor) GetDescription() string {
	return "prints all commands"
}

func (*HelpExecutor) GetName() string {
	return "help"
}

func (e *HelpExecutor) Supports(cmd Command) bool {
	return cmd.Name == e.GetName()
}

func (e HelpExecutor) Run(cmd Command) error {
	executors := append([]Executor{&e}, e.executors...)
	sort.Slice(executors, func(i, j int) bool {
		return executors[i].GetName() < executors[j].GetName()
	})

	lines := make([]string, 0, len(executors))

	for _, executor := range executors {
		lines = append(lines, "/"+executor.GetName()+" - "+executor.GetDescription())
	}

	return e.replier.ReplyTo(cmd.Message, strings.Join(lines, "\n"))
}
