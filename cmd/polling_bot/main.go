package main

import (
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.local", ".env")
	if err != nil {
		panic(err)
	}

	pollingBot, err := initPollingBot()
	if err != nil {
		panic(err)
	}

	pollingBot.Run()
}
