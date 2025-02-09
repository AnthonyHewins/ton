package ingest

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/AnthonyHewins/ton/gen/go/marketdata/v0"
	"github.com/AnthonyHewins/tradovate"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ChartPublisher struct {
	prefix  string
	js      jetstream.JetStream
	timeout time.Duration
	logger  *slog.Logger

	barChartErrs  prometheus.Counter
	tickChartErrs prometheus.Counter
}

func (c *ChartPublisher) Publish(chart *tradovate.Chart) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	if len(chart.Bars) > 0 {
		if err := c.pushBars(ctx, chart); err != nil {
			c.barChartErrs.Inc()
		}
		return
	}

	if err := c.pushTicks(ctx, chart); err != nil {
		c.tickChartErrs.Inc()
	}
}

func (c *ChartPublisher) pushBars(ctx context.Context, chart *tradovate.Chart) error {
	bars := make([]*marketdata.Bar, len(chart.Bars))
	for i, v := range chart.Bars {
		bars[i] = &marketdata.Bar{
			Timestamp:   timestamppb.New(v.Timestamp),
			Open:        v.Open,
			High:        v.High,
			Low:         v.Low,
			Close:       v.Close,
			UpVolume:    v.UpVolume,
			DownVolume:  v.DownVolume,
			UpTicks:     v.UpTicks,
			DownTicks:   v.DownTicks,
			BidVolume:   v.BidVolume,
			OfferVolume: v.OfferVolume,
		}
	}

	buf, err := proto.Marshal(&marketdata.BarChart{
		Id:        int64(chart.ID),
		TradeDate: timestamppb.New(chart.BaseTimestamp),
		Bars:      bars,
	})

	if err != nil {
		c.logger.ErrorContext(ctx, "failed marshal of bar chart", "err", err, "chart", chart)
		return err
	}

	if _, err = c.js.Publish(ctx, fmt.Sprintf("%s.chart.%d", c.prefix, chart.ID), buf); err != nil {
		c.logger.ErrorContext(ctx, "failed publishing bar chart", "err", err, "chart", chart)
		return err
	}

	c.logger.DebugContext(ctx, "published bar chart", "id", chart.ID)
	return nil
}

func (c *ChartPublisher) pushTicks(ctx context.Context, chart *tradovate.Chart) error {
	ticks := make([]*marketdata.Tick, len(chart.Ticks))
	for i, v := range chart.Ticks {
		ticks[i] = &marketdata.Tick{
			Id:               int64(v.ID),
			RelativeTime:     int64(v.RelativeTime),
			RelativePrice:    int64(v.RelativePrice),
			Volume:           int64(v.Volume),
			RelativeBidPrice: v.RelativeBidPrice,
			RelativeAskPrice: v.RelativeAskPrice,
			BidSize:          v.BidSize,
			AskSize:          v.AskSize,
		}
	}

	buf, err := proto.Marshal(&marketdata.TickChart{
		Id:            int64(chart.ID),
		TradeDate:     timestamppb.New(chart.Td),
		EndOfHistory:  chart.EndOfHistory,
		Source:        chart.Source,
		BasePrice:     int64(chart.BasePrice),
		BaseTimestamp: timestamppb.New(chart.BaseTimestamp),
		TickSize:      chart.TickSize,
		Ticks:         ticks,
	})

	if err != nil {
		c.logger.ErrorContext(ctx, "failed marshal of tick chart", "err", err, "chart", chart)
		return err
	}

	if _, err := c.js.Publish(ctx, fmt.Sprintf("%s.tick.%d", c.prefix, chart.ID), buf); err != nil {
		c.logger.ErrorContext(ctx, "failed publishing tick chart", "err", err, "chart", chart)
		return err
	}

	c.logger.DebugContext(ctx, "published tick chart", "id", chart.ID)
	return nil
}
