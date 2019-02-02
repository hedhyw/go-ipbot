package main

import (
	"sync"

	"github.com/hedhyw/ip-bot/internal/config"
	"github.com/hedhyw/ip-bot/internal/tg"
	log "github.com/sirupsen/logrus"
)

func main() {
	tgBot := tg.NewBot()
	go func() {
		err := tgBot.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()

	s := &sync.Mutex{}
	s.Lock()
	s.Lock()
}

func init() {
	lvl, err := log.ParseLevel(config.App().LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(lvl)
}
