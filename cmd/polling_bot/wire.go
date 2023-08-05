//go:build wireinject
// +build wireinject

package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/wire"

	"memo/configs"
	"memo/internal/api/telegram/command"
	"memo/internal/apps"
	"memo/internal/bootstrap"
	"memo/internal/domain"
	"memo/internal/infra"
	"memo/internal/pkg/telegram"
)

func initApp() (*apps.PollingBotApp, error) { //nolint
	wire.Build(
		bootstrap.NewTgAPIPollingBot,
		telegram.NewReplier,
		configs.GetAppConfig,
		configs.GetTelegramConfig,
		bootstrap.NewParser,
		bootstrap.NewTgBot,
		wire.Bind(new(telegram.BotAPI), new(*tgbotapi.BotAPI)),
		configs.GetDBConfig,
		bootstrap.NewGORMDb,
		infra.NewMemoGORMRepository,
		wire.Bind(new(domain.MemoRepository), new(*infra.MemoGORMRepository)),
		infra.NewUserGORMRepository,
		wire.Bind(new(domain.UserRepository), new(*infra.UserGORMRepository)),
		bootstrap.NewRandExecutor,
		command.NewStartExecutor,
		command.NewAddExecutor,
		command.NewDeleteExecutor,
		bootstrap.NewSearchExecutor,
		bootstrap.NewDefaultCommandExecutor,
		bootstrap.NewCommandExecutors,
		apps.NewPollingBotApp,
	)

	return nil, nil
}
