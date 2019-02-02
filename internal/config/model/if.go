package model

// IFConfig contains ifconfig.me url
type IFConfig struct {
	URL string `env:"IPBOT_IFCONFIG_URL" envDefault:"ifconfig.me"`
}
