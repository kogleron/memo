// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"memo/internal/apps"
	"memo/internal/bootstrap"
	"memo/internal/configs"
	"memo/internal/memo"
	"memo/internal/user"
)

// Injectors from wire.go:

func initRandCommand() (*apps.RandCommand, error) {
	appConfig := configs.GetAppConfig()
	dbConfig := configs.GetDBConfig()
	db, err := bootstrap.NewGORMDb(dbConfig)
	if err != nil {
		return nil, err
	}
	repository := memo.NewRepository(db)
	userRepository := user.NewRepository(db)
	telegramConfig := configs.GetTelegramConfig()
	botAPI, err := bootstrap.NewTgBot(telegramConfig)
	if err != nil {
		return nil, err
	}
	randCommand := bootstrap.NewRandCommand(appConfig, repository, userRepository, botAPI)
	return randCommand, nil
}
