//go:build wireinject
// +build wireinject

package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/wire"

	"memo/configs"
	"memo/internal/apps"
	"memo/internal/bootstrap"
	"memo/internal/domain"
	"memo/internal/infra"
	"memo/internal/pkg/telegram"
)

func initRandCommand() (*apps.RandCommand, error) {
	wire.Build(
		configs.GetAppConfig,
		configs.GetTelegramConfig,
		bootstrap.NewTgBot,
		wire.Bind(new(telegram.BotAPI), new(*tgbotapi.BotAPI)),
		configs.GetDBConfig,
		bootstrap.NewGORMDb,
		infra.NewMemoGORMRepository,
		wire.Bind(new(domain.MemoRepository), new(*infra.MemoGORMRepository)),
		infra.NewUserGORMRepository,
		wire.Bind(new(domain.UserRepository), new(*infra.UserGORMRepository)),
		bootstrap.NewRandCommand,
	)

	return &apps.RandCommand{}, nil
}
