package conf

import (
	"time"

	"github.com/AnthonyHewins/tradovate"
	"github.com/google/uuid"
)

type Tradovate struct {
	TradovateRestURL      string `env:"TRADOVATE_REST_URL" envDefault:"https://demo.tradovateapi.com/v1"`
	TradovateWebsocketURL string `env:"TRADOVATE_WS_URL" envDefault:"wss://demo.tradovateapi.com/v1/websocket"`

	Timeout time.Duration `env:"TRADOVATE_TIMEOUT" envDefault:"3s"`

	Name       string    `env:"TRADOVATE_NAME"`
	Password   string    `env:"TRADOVATE_PASSWORD"`
	AppID      string    `env:"TRADOVATE_APPID"`
	AppVersion string    `env:"TRADOVATE_APPVERSION"`
	CID        string    `env:"TRADOVATE_CLIENT_ID"`
	DeviceID   uuid.UUID `env:"TRADOVATE_DEVICEID"`
	Sec        uuid.UUID `env:"TRADOVATE_SECRET"`
}

func (t *Tradovate) Creds() *tradovate.Creds {
	return &tradovate.Creds{
		Name:       t.Name,
		Password:   t.Password,
		AppID:      t.AppID,
		AppVersion: t.AppVersion,
		ClientID:   t.CID,
		DeviceID:   t.DeviceID,
		Secret:     t.Sec,
	}
}
