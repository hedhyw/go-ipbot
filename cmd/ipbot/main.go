package main

import (
	"github.com/hedhyw/ip-bot/internal/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info(config.App())
}

func init() {
	lvl, err := log.ParseLevel(config.App().LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(lvl)
}
