//go:build wireinject
// +build wireinject

package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/wire"

	"memo/internal/apps"
	"memo/internal/bootstrap"
	"memo/internal/command"
	"memo/internal/configs"
	"memo/internal/memo"
	"memo/internal/telegram"
	"memo/internal/user"
)

func initPollingBot() (*apps.PollingBot, error) { //nolint
	wire.Build(
		configs.GetAppConfig,
		configs.GetTelegramConfig,
		command.NewParser,
		bootstrap.NewTgBot,
		wire.Bind(new(telegram.BotAPI), new(*tgbotapi.BotAPI)),
		configs.GetDBConfig,
		bootstrap.NewGORMDb,
		memo.NewGORMRepository,
		wire.Bind(new(memo.Repository), new(*memo.GORMRepository)),
		user.NewGORMRepository,
		wire.Bind(new(user.Repository), new(*user.GORMRepository)),
		bootstrap.NewRandExecutor,
		command.NewStartExecutor,
		command.NewAddExecutor,
		command.NewDeleteExecutor,
		bootstrap.NewSearchExecutor,
		bootstrap.NewDefaultCommandExecutor,
		bootstrap.NewCommandExecutors,
		bootstrap.NewPollingBot,
	)

	return &apps.PollingBot{}, nil
}
