package model

// AppConfig contains main application preferences
type AppConfig struct {
	LogLevel   string   `env:"IPBOT_LOGLEVEL" envDefault:"info"`
	SuperUsers []string `env:"IPBOT_USERS" envSeparator:":"`
}
