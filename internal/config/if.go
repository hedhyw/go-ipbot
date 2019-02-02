package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/caarlos0/env"
	"github.com/hedhyw/ip-bot/internal/config/model"
)

var ifConf model.IFConfig

// IF is IFConfig getter
func IF() *model.IFConfig {
	return &ifConf
}

func init() {
	if err := env.Parse(&ifConf); err != nil {
		log.Fatal(err)
	}
}
