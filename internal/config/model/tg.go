package model

// TGConfig contains telegram-bot config
type TGConfig struct {
	BotToken   string `env:"IPBOT_TG_TOKEN" envDefault:""`
	BotDebug   bool   `env:"IPBOT_TG_DEBUG" envDefault:"false"`
	BotTimeout int    `env:"IPBOT_TG_TIMEOUT" envDefault:"256"`
}
