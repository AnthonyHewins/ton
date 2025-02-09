package main

import (
	"context"
	"time"

	"github.com/AnthonyHewins/ton/internal/conf"
	"github.com/AnthonyHewins/ton/internal/ingest"
	"github.com/AnthonyHewins/tradovate"
	"github.com/caarlos0/env/v11"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/prometheus/client_golang/prometheus"
)

func newCounter(system, name, desc string) prometheus.Counter {
	return prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: appName,
		Subsystem: system,
		Name:      name,
		Help:      desc,
	})
}

type app struct {
	*conf.Server
	kv         jetstream.KeyValue
	ws         *tradovate.WS
	controller *ingest.Controller
}

type consumer struct {
	ctx      jetstream.ConsumeContext
	ingestor jetstream.Consumer
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

	a := app{Server: (*conf.Server)(b)}
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

	controller := ingest.New(appName, c.Prefix, a.Logger, a.TP, js, kv, c.Timeout)
	a.ws, err = b.Socket(ctx,
		&c.Tradovate,
		tradovate.WithChartHandler(controller.Chart.Publish),
		// tradovate.WithEntityHandler(a.controller.Entity.PublishEntity),
		// tradovate.WithMarketDataHandler(a.controller.MarketData.Publish),
		tradovate.WithErrHandler(func(err error) {
			if err != nil {
				a.Logger.ErrorContext(ctx, "websocket error", "err", err)
			}
		}),
	)

	if err != nil {
		a.Logger.ErrorContext(ctx, "failed tradovate websocket creation", "err", err)
		return nil, err
	}

	if err = a.controller.Initialize(ctx, a.ws); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a *app) shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if err := a.ws.Close(ctx); err != nil {
		a.Logger.ErrorContext(ctx, "failed graceful shutdown of tradovate websocket", "err", err)
	}

	a.Server.Shutdown(ctx)
}
