package apps

import (
	"log"

	"memo/internal/api/telegram"
)

func NewPollingBotApp(
	api telegram.PollingBot,
) *PollingBotApp {
	return &PollingBotApp{
		api: api,
	}
}

type PollingBotApp struct {
	api telegram.PollingBot
}

func (a *PollingBotApp) Run() {
	a.api.Run()
	log.Println("Shutdowning...")
}
