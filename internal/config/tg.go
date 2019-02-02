package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/caarlos0/env"
	"github.com/hedhyw/ip-bot/internal/config/model"
)

var tgConf model.TGConfig

// TG is TGConfig getter
func TG() *model.TGConfig {
	return &tgConf
}

func init() {
	if err := env.Parse(&tgConf); err != nil {
		log.Fatal(err)
	}
}
