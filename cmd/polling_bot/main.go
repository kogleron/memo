package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"memo/internal/apps"
	"memo/internal/command"
	"memo/internal/configs"
	"memo/internal/memo"
)

func main() {
	err := godotenv.Load(".env.local", ".env")
	if err != nil {
		panic(err)
	}

	telegramConf := configs.GetTelegramConfig()
	tgBot, err := tgbotapi.NewBotAPI(telegramConf.BotToken)
	if err != nil {
		panic(err)
	}
	tgBot.Debug = telegramConf.DebugMode
	log.Printf("Authorized on account %s\n", tgBot.Self.UserName)

	conf := configs.GetDbConfig()
	db, err := gorm.Open(sqlite.Open(conf.Database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	memoRepo := memo.NewRepository(db)

	cmdParser := command.NewParser()
	cmdExecutors := []command.Executor{
		command.NewDefaultCommandExecutor(
			command.NewAddExecutor(memoRepo),
		),
	}

	pollingBot := apps.NewPollingBot(
		tgBot,
		telegramConf,
		cmdParser,
		cmdExecutors,
	)
	pollingBot.Run()
}
