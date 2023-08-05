package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"memo/configs"
	"memo/internal/domain"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	conf := configs.GetDBConfig()

	db, err := gorm.Open(sqlite.Open(conf.Database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&domain.Memo{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		panic(err)
	}
}
