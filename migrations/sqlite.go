package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"memo/internal/configs"
	"memo/internal/memo"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	conf := configs.GetDbConfig()

	db, err := gorm.Open(sqlite.Open(conf.Database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&memo.Memo{})
	if err != nil {
		panic(err)
	}
}
