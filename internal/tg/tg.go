package tg

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hedhyw/go-ipbot/internal/ip"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hedhyw/go-ipbot/internal/config"
	"github.com/hedhyw/go-ipbot/internal/interfaces"
	"github.com/labstack/gommon/log"
)

type bot struct {
	api    *tgbotapi.BotAPI
	ifconf interfaces.IFConfig
}

// NewBot telegram
func NewBot() interfaces.Starter {
	bot := bot{}
	bot.ifconf = ip.NewIFConfig()
	return &bot
}

// Start main bot loop
func (bot *bot) Start() error {

	// HTTP Client
	var clnt *http.Client
	if config.App().ProxyURL == "" {
		clnt = &http.Client{}
	} else {
		u, err := url.Parse(config.App().ProxyURL)
		if err != nil {
			return err
		}
		clnt = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(u),
			},
		}
	}

	// API initialization
	api, err := tgbotapi.NewBotAPIWithClient(
		config.TG().BotToken,
		clnt,
	)
	if err != nil {
		return err
	}
	bot.api = api

	u := tgbotapi.NewUpdate(0) // 0 is offset
	u.Timeout = config.TG().BotTimeout

	// Updates loop
	updates, err := api.GetUpdatesChan(u)
	if err != nil {
		return err
	}
	for update := range updates {
		if err := bot.handleUpdate(update); err != nil {
			log.Error(err)
		}
	}

	return errors.New("Update channel is closed")
}

func (bot *bot) stateText() string {
	ip, err := bot.ifconf.GetIP()
	if err != nil {
		return bot.messageError(err)
	}
	return fmt.Sprintf("IP Address `%s`", ip)
}

func (bot *bot) messageError(err error) string {
	return fmt.Sprintf("Error `%v`", err)
}

func (bot *bot) handleUpdate(upd tgbotapi.Update) error {
	// Ignore non-message update
	if upd.Message == nil {
		return nil
	}

	// Check access
	userName := upd.Message.From.UserName
	userHasAccess := false
	for _, superUser := range config.App().SuperUsers {
		if superUser == userName {
			userHasAccess = true
			break
		}
	}
	if !userHasAccess {
		return fmt.Errorf("Access denied for @%s", userName)
	}

	// Send message
	msg := tgbotapi.NewMessage(
		upd.Message.Chat.ID,
		bot.stateText(),
	)
	msg.ParseMode = tgbotapi.ModeMarkdown
	msg.DisableWebPagePreview = true
	_, err := bot.api.Send(msg)

	return err
}
