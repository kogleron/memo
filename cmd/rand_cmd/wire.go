//go:build wireinject
// +build wireinject

package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/wire"

	"memo/internal/apps"
	"memo/internal/bootstrap"
	"memo/internal/configs"
	"memo/internal/memo"
	"memo/internal/telegram"
	"memo/internal/user"
)

func initRandCommand() (*apps.RandCommand, error) {
	wire.Build(
		configs.GetAppConfig,
		configs.GetTelegramConfig,
		bootstrap.NewTgBot,
		wire.Bind(new(telegram.BotAPI), new(*tgbotapi.BotAPI)),
		configs.GetDBConfig,
		bootstrap.NewGORMDb,
		memo.NewGORMRepository,
		wire.Bind(new(memo.Repository), new(*memo.GORMRepository)),
		user.NewGORMRepository,
		wire.Bind(new(user.Repository), new(*user.GORMRepository)),
		bootstrap.NewRandCommand,
	)

	return &apps.RandCommand{}, nil
}
