package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/caarlos0/env"
	"github.com/hedhyw/ip-bot/internal/config/model"
)

var appConf model.AppConfig

// App is AppConfig getter
func App() *model.AppConfig {
	return &appConf
}

func init() {
	if err := env.Parse(&appConf); err != nil {
		log.Fatal(err)
	}
}
