//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"memo/internal/apps"
	"memo/internal/bootstrap"
	"memo/internal/command"
	"memo/internal/configs"
	"memo/internal/memo"
	"memo/internal/user"
)

func initPollingBot() (*apps.PollingBot, error) { //nolint
	wire.Build(
		configs.GetAppConfig,
		configs.GetTelegramConfig,
		command.NewParser,
		bootstrap.NewTgBot,
		configs.GetDBConfig,
		bootstrap.NewGORMDb,
		memo.NewRepository,
		user.NewRepository,
		bootstrap.NewRandExecutor,
		command.NewStartExecutor,
		command.NewAddExecutor,
		bootstrap.NewDefaultCommandExecutor,
		bootstrap.NewCommandExecutors,
		bootstrap.NewPollingBot,
	)

	return &apps.PollingBot{}, nil
}
