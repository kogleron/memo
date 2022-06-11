// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"memo/internal/apps"
	"memo/internal/bootstrap"
	"memo/internal/command"
	"memo/internal/configs"
	"memo/internal/memo"
	"memo/internal/user"
)

// Injectors from wire.go:

func initPollingBot() (*apps.PollingBot, error) {
	telegramConfig := configs.GetTelegramConfig()
	botAPI, err := bootstrap.NewTgBot(telegramConfig)
	if err != nil {
		return nil, err
	}
	parser := command.NewParser()
	appConfig := configs.GetAppConfig()
	dbConfig := configs.GetDBConfig()
	db, err := bootstrap.NewGORMDb(dbConfig)
	if err != nil {
		return nil, err
	}
	gormRepository, err := memo.NewGORMRepository(db)
	if err != nil {
		return nil, err
	}
	userGORMRepository := user.NewGORMRepository(db)
	randExecutor := bootstrap.NewRandExecutor(appConfig, gormRepository, botAPI, userGORMRepository)
	startExecutor := command.NewStartExecutor(userGORMRepository, botAPI)
	searchExecutor := bootstrap.NewSearchExecutor(botAPI, gormRepository, userGORMRepository, appConfig)
	addExecutor := command.NewAddExecutor(gormRepository, botAPI, userGORMRepository)
	defaultCommandExecutor := bootstrap.NewDefaultCommandExecutor(addExecutor)
	commandExecutors := bootstrap.NewCommandExecutors(randExecutor, startExecutor, searchExecutor, defaultCommandExecutor)
	pollingBot := bootstrap.NewPollingBot(botAPI, telegramConfig, parser, commandExecutors)
	return pollingBot, nil
}
