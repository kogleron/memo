package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type TelegramConfig struct {
	BotToken           string `envconfig:"BOT_TOKEN" required:"true"`
	AllowedAccountsMap map[string]struct{}
	AllowedAccounts    []string `envconfig:"ALLOWED_ACCOUNTS" required:"true"`
	DebugMode          bool     `envconfig:"DEBUG_MODE" required:"false"`
}

func GetTelegramConfig() TelegramConfig {
	conf := TelegramConfig{}
	conf.AllowedAccountsMap = make(map[string]struct{})

	err := envconfig.Process("TELEGRAM", &conf)
	if err != nil {
		panic(err)
	}

	for _, account := range conf.AllowedAccounts {
		conf.AllowedAccountsMap[account] = struct{}{}
	}

	return conf
}
