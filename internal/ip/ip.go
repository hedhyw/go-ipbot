package ip

import (
	"net/http"

	"github.com/hedhyw/go-ipbot/internal/config"
	"github.com/hedhyw/go-ipbot/internal/interfaces"
)

type ifConfig struct {
}

// NewIFConfig for getting remote ip address
func NewIFConfig() interfaces.IFConfig {
	return &ifConfig{}
}

// GetIP remote address with ifconfig
func (ifconf *ifConfig) GetIP() (string, error) {
	req, err := http.NewRequest(http.MethodGet, config.IF().URL, nil)
	if err != nil {
		return "", err
	}

	clnt := http.Client{}
	resp, err := clnt.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	buf := make([]byte, 3*4+3) // 144.144.144.144
	_, err = resp.Body.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf), err
}
