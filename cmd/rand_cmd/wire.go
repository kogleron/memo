//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"memo/internal/apps"
	"memo/internal/bootstrap"
	"memo/internal/configs"
	"memo/internal/memo"
	"memo/internal/user"
)

func initRandCommand() (*apps.RandCommand, error) {
	wire.Build(
		configs.GetAppConfig,
		configs.GetTelegramConfig,
		bootstrap.NewTgBot,
		configs.GetDBConfig,
		bootstrap.NewGORMDb,
		memo.NewRepository,
		user.NewRepository,
		bootstrap.NewRandCommand,
	)

	return &apps.RandCommand{}, nil
}
