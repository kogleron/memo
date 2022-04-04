package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"memo/internal/apps"
	"memo/internal/configs"
	"memo/internal/memo"
	"memo/internal/user"
)

func main() {
	err := godotenv.Load(".env.local", ".env")
	if err != nil {
		panic(err)
	}

	appConf := configs.GetAppConfig()

	telegramConf := configs.GetTelegramConfig()

	tgBot, err := tgbotapi.NewBotAPI(telegramConf.BotToken)
	if err != nil {
		panic(err)
	}

	tgBot.Debug = telegramConf.DebugMode
	log.Printf("Authorized on account %s\n", tgBot.Self.UserName)

	conf := configs.GetDBConfig()

	db, err := gorm.Open(sqlite.Open(conf.Database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	memoRepo := memo.NewRepository(db)
	userRepo := user.NewRepository(db)

	randCommand := apps.NewRandCommand(uint(appConf.RandQty), memoRepo, userRepo, tgBot)
	randCommand.Run()
}
