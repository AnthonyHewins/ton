package controller

import (
	"context"
	"log/slog"
	"time"

	"github.com/AnthonyHewins/tradovate"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

type Controller struct {
	logger *slog.Logger
	Chart  ChartPublisher
	Entity EntityPublisher
}

func New(app, subjectPrefix string, logger *slog.Logger, tp trace.TracerProvider, js jetstream.JetStream, kv jetstream.KeyValue, timeout time.Duration) *Controller {
	metric := func(subsystem, name, help string) prometheus.Counter {
		return prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: app,
			Subsystem: subsystem,
			Name:      name,
			Help:      help,
		})
	}

	return &Controller{
		logger: logger,
		Entity: EntityPublisher{
			prefix:     subjectPrefix,
			js:         js,
			orderKV:    kv,
			orderErr:   metric("order", "publish_err", "Number of errors publishing errors to KV"),
			publishErr: metric("entity", "publish_err", "Number of errors publishing entities"),
			timeout:    timeout,
			logger:     logger,
		},
		Chart: ChartPublisher{
			js:            js,
			timeout:       timeout,
			logger:        logger,
			barChartErrs:  metric("chart", "publish_bar_chart_errs", "Number of bar chart publish errors"),
			tickChartErrs: metric("chart", "publish_tick_chart_errs", "Number of bar chart publish errors"),
		},
	}
}

func (c *Controller) Initialize(ctx context.Context, ws *tradovate.WS) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error { return c.initOrders(ctx, ws) })

	return g.Wait()
}

func (c *Controller) initOrders(ctx context.Context, ws *tradovate.WS) error {
	orders, err := ws.ListOrders(ctx)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed listing orders", "err", err)
		return err
	}

	for _, v := range orders {
		if err = c.Entity.publishOrder(ctx, v); err != nil {
			return err
		}
	}

	c.logger.InfoContext(ctx, "initialized order keyvalue", "len(orders)", len(orders))
	return nil
}
