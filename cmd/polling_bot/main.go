package main

import (
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.local", ".env")
	if err != nil {
		panic(err)
	}

	app, err := initApp()
	if err != nil {
		panic(err)
	}

	app.Run()
}
