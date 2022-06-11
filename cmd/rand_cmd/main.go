package main

import (
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.local", ".env")
	if err != nil {
		panic(err)
	}

	randCommand, err := initRandCommand()
	if err != nil {
		panic(err)
	}

	randCommand.Run()
}
