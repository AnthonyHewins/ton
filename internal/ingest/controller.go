package ingest

import (
	"context"
	"log/slog"
	"net/http"
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
	Orders Orders
	Entity EntityPub
}

type Opts struct {
	App, Prefix string
	Logger      *slog.Logger
	TP          trace.TracerProvider
	jetstream.JetStream
	jetstream.KeyValue

	Token *tradovate.Token

	SocketURL string
	RestURL   string
	*tradovate.Creds

	Timeout time.Duration
}

func New(ctx context.Context, opts *Opts) (*Controller, error) {
	metric := func(subsystem, name, help string) prometheus.Counter {
		return prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: opts.App,
			Subsystem: subsystem,
			Name:      name,
			Help:      help,
		})
	}

	e := EntityPub{
		prefix:        opts.Prefix,
		logger:        opts.Logger,
		timeout:       opts.Timeout,
		js:            opts.JetStream,
		entityPubErrs: metric("entity", "publish_err", "Number of publish errors"),
		orders: entityKV[*ordersvc.Order, *tradovate.Order]{
			logger: opts.Logger,
			kv:     opts.KeyValue,
			c: orderTranslator{
				putErrs: metric("orderkv", "put_errs", "Count of PUT errors in nats KV"),
				delErrs: metric("orderkv", "del_errs", "Count of DEL errors in nats KV"),
			},
		},
		positions: entityKV[*positionpb.Position, *tradovate.Position]{
			logger: opts.Logger,
			kv:     opts.KeyValue,
			c: positionTranslator{
				putErrs: metric("position", "put_errs", "Count of PUT errors in nats KV"),
				delErrs: metric("position", "del_errs", "Count of DEL errors in nats KV"),
			},
		},
	}

	r := tradovate.NewREST(opts.RestURL, &http.Client{Timeout: opts.Timeout}, opts.Creds)

	if opts.Token != nil && !opts.Token.Expired() {
		r.SetToken(opts.Token)
	}

	ws, err := tradovate.NewSocket(
		ctx,
		opts.SocketURL,
		nil,
		r,
		tradovate.WithTimeout(opts.Timeout),
		tradovate.WithEntityHandler(e.PublishEntity),
	)

	if err != nil {
		opts.Logger.ErrorContext(ctx,
			"failed connecting to tradovate WS",
			"err", err,
			"creds", opts.Creds,
		)
		return nil, err
	}

	return &Controller{
		Orders: Orders{
			logger:      opts.Logger,
			ws:          ws,
			orderErr:    metric("orders", "place_order_err", "Count of errors placing regular orders"),
			ocoErr:      metric("orders", "place_oco_order_err", "Count of errors placing oco orders"),
			osoErr:      metric("orders", "place_oso_order_err", "Count of errors placing oso orders"),
			placedOrder: metric("orders", "place_order", "Count of placed regular orders"),
			placedOCO:   metric("orders", "place_oco_order", "Count of  placed oco orders"),
			placedOSO:   metric("orders", "place_oso_order", "Count of placed oso orders"),
		},
		Entity: e,
	}, nil
}

func (c *Controller) Close() error {
	if c.Orders.ws == nil {
		return nil
	}

	if err := c.Orders.ws.Close(); err != nil {
		c.Orders.logger.Error("failed closing socket")
		return err
	}

	c.Orders.logger.Debug("closed socket")
	return nil
}

func (c *Controller) Initialize(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error { return c.initOrders(ctx, c.Orders.ws) })
	g.Go(func() error { return c.initPositions(ctx, c.Orders.ws) })

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
