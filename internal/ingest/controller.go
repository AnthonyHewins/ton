package ingest

import (
	"context"
	"log/slog"
	"time"

	"github.com/AnthonyHewins/ton/gen/go/ordersvc/v0"
	"github.com/AnthonyHewins/ton/gen/go/positionpb/v0"
	"github.com/AnthonyHewins/tradovate"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

type Controller struct {
	Entity     EntityPub
	Chart      ChartPublisher
	MarketData MarketData
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
		MarketData: MarketData{
			prefix:           subjectPrefix,
			logger:           logger,
			timeout:          timeout,
			js:               js,
			quotePubErrs:     metric("quote", "publish_err", "Number of quote publish errors"),
			domPubErrs:       metric("dom", "publish_err", "Number of dom publish errors"),
			histogramPubErrs: metric("histogram", "publish_err", "Number of histogram publish errors"),
		},
		Entity: EntityPub{
			prefix:        subjectPrefix,
			logger:        logger,
			timeout:       timeout,
			js:            js,
			entityPubErrs: metric("entity", "publish_err", "Number of publish errors"),
			orders: entityKV[*ordersvc.Order, *tradovate.Order]{
				logger: logger,
				kv:     kv,
				c: orderTranslator{
					putErrs: metric("order", "put_errs", "Count of PUT errors in nats KV"),
					delErrs: metric("order", "del_errs", "Count of DEL errors in nats KV"),
				},
			},
			positions: entityKV[*positionpb.Position, *tradovate.Position]{
				logger: logger,
				kv:     kv,
				c: positionTranslator{
					putErrs: metric("position", "put_errs", "Count of PUT errors in nats KV"),
					delErrs: metric("position", "del_errs", "Count of DEL errors in nats KV"),
				},
			},
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
	g.Go(func() error { return c.initPositions(ctx, ws) })

	return g.Wait()
}

func (c *Controller) initPositions(ctx context.Context, ws *tradovate.WS) error {
	p := c.Entity.positions

	positions, err := ws.ListPositions(ctx)
	if err != nil {
		p.logger.ErrorContext(ctx, "failed listing positions", "err", err)
		return err
	}

	for _, v := range positions {
		if err = p.put(ctx, v); err != nil {
			return err
		}
	}

	p.logger.InfoContext(ctx, "initialized position keyvalue", "len(positions)", len(positions))
	return nil
}

func (c *Controller) initOrders(ctx context.Context, ws *tradovate.WS) error {
	o := c.Entity.orders

	orders, err := ws.ListOrders(ctx)
	if err != nil {
		o.logger.ErrorContext(ctx, "failed listing orders", "err", err)
		return err
	}

	for _, v := range orders {
		if err = o.put(ctx, v); err != nil {
			return err
		}
	}

	o.logger.InfoContext(ctx, "initialized order keyvalue", "len(orders)", len(orders))
	return nil
}
