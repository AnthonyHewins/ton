package conf

import (
	"context"
	"net/http"
	"time"

	"github.com/AnthonyHewins/tradovate"
	"github.com/google/uuid"
)

type Tradovate struct {
	TradovateRestURL      string `env:"TRADOVATE_REST_URL" envDefault:"wss://demo.tradovateapi.com/v1/websocket"`
	TradovateWebsocketURL string `env:"TRADOVATE_WS_URL" envDefault:""`

	Timeout time.Duration `env:"TRADOVATE_TIMEOUT" envDefault:"3s"`

	Name       string    `env:"TRADOVATE_NAME"`
	Password   string    `env:"TRADOVATE_PASSWORD"`
	AppID      string    `env:"TRADOVATE_APPID"`
	AppVersion string    `env:"TRADOVATE_APPVERSION"`
	CID        string    `env:"TRADOVATE_CID"`
	DeviceID   uuid.UUID `env:"TRADOVATE_DEVICEID"`
	Sec        uuid.UUID `env:"TRADOVATE_SEC"`
}

func (t *Tradovate) Creds() *tradovate.Creds {
	return &tradovate.Creds{
		Name:       t.Name,
		Password:   t.Password,
		AppID:      t.AppID,
		AppVersion: t.AppVersion,
		CID:        t.CID,
		DeviceID:   t.DeviceID,
		Sec:        t.Sec,
	}
}

func (b *Bootstrapper) Socket(ctx context.Context, c *Tradovate, opts ...tradovate.WSOpt) (*tradovate.WS, error) {
	ws, err := tradovate.NewSocket(
		ctx,
		c.TradovateWebsocketURL,
		nil,
		tradovate.NewREST(
			c.TradovateRestURL,
			&http.Client{Timeout: c.Timeout},
			c.Creds(),
		),
		append(opts, tradovate.WithTimeout(c.Timeout))...,
	)

	if err != nil {
		b.Logger.ErrorContext(ctx, "failed connecting to tradovate WS", "err", err)
		return nil, err
	}

	return ws, nil
}
