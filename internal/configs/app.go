package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	RandQty int `envconfig:"RAND_QTY" required:"true"`
}

func GetAppConfig() AppConfig {
	conf := AppConfig{}

	err := envconfig.Process("APP", &conf)
	if err != nil {
		panic(err)
	}

	return conf
}
