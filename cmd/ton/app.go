package main

import (
	"context"
	"time"

	"github.com/AnthonyHewins/ton/internal/conf"
	"github.com/AnthonyHewins/ton/internal/controller"
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
	controller *controller.Controller
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

	a.kv, err = js.KeyValue(ctx, c.Bucket)
	if err != nil {
		a.Logger.ErrorContext(ctx, "failed connecting to nalpaca KV bucket", "err", err, "bucket", c.Bucket)
		return nil, err
	}

	a.controller = controller.New(appName, a.Logger, a.TP, js, c.Timeout)
	a.ws, err = b.Socket(ctx,
		&c.Tradovate,
		tradovate.WithChartHandler(a.controller.Chart.Publish),
		tradovate.WithEntityHandler(func(em *tradovate.EntityMsg) {}),
		tradovate.WithMarketDataHandler(func(md *tradovate.MarketData) {}),
		tradovate.WithErrHandler(func(err error) {
			if err != nil {
				a.Logger.ErrorContext(ctx, "websocket error", "err", err)
			}
		}),
	)

	return &a, nil
}

// func (a *app) connect(ctx context.Context, js jetstream.JetStream, stream, consumer string) (jetstream.Consumer, error) {
// x, err := js.Consumer(ctx, stream, consumer)
// if err != nil {
// a.Logger.ErrorContext(ctx,
// "failed connecting to consumer",
// "err", err,
// "stream", stream,
// "consumer", consumer,
// )
// }
//
// return x, err
// }

func (a *app) shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	a.Server.Shutdown(ctx)
}
