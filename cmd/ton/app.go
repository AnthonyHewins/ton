package main

import (
	"context"
	"time"

	"github.com/AnthonyHewins/ton/internal/conf"
	"github.com/AnthonyHewins/ton/internal/ingest"
	"github.com/caarlos0/env/v11"
	"github.com/nats-io/nats.go/jetstream"
)

type app struct {
	*conf.Server
	accountSpec string
	accountID   uint
	controller  *ingest.Controller
}

func newApp(ctx context.Context) (*app, error) {
	var c config
	if err := env.Parse(&c); err != nil {
		return nil, err
	}

	b, err := c.BootstrapConf.New(ctx, appName)
	if err != nil {
		return nil, err
	}

	a := app{
		Server:      (*conf.Server)(b),
		accountSpec: c.AccountSpec,
		accountID:   c.AccountID,
	}
	defer func() {
		if err != nil {
			a.shutdown()
		}
	}()

	js, err := jetstream.New(a.NC)
	if err != nil {
		a.Logger.ErrorContext(ctx, "failed connecting to jetstream", "err", err)
		return nil, err
	}

	kv, err := js.KeyValue(ctx, c.Bucket)
	if err != nil {
		a.Logger.ErrorContext(ctx, "failed connecting to nalpaca KV bucket", "err", err, "bucket", c.Bucket)
		return nil, err
	}

	a.controller, err = ingest.New(ctx, &ingest.Opts{
		App:       appName,
		Prefix:    c.Prefix,
		Logger:    a.Logger,
		TP:        a.TP,
		JetStream: js,
		KeyValue:  kv,
		SocketURL: c.TradovateWebsocketURL,
		RestURL:   c.TradovateRestURL,
		Creds:     c.Creds(),
		Timeout:   c.Timeout,
	})
	if err != nil {
		return nil, err
	}

	if err != nil {
		a.Logger.ErrorContext(ctx, "failed tradovate websocket creation", "err", err)
		return nil, err
	}

	if err = a.controller.Initialize(ctx); err != nil {
		return nil, err
	}

	a.Logger.InfoContext(ctx, "finished bootstrapping")
	return &a, nil
}

func (a *app) shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	a.controller.Close()
	a.Server.Shutdown(ctx)
}
