package ingest

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/AnthonyHewins/ton/gen/go/marketdata/v0"
	"github.com/AnthonyHewins/ton/pkg/ton"
	"github.com/AnthonyHewins/tradovate"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/proto"
)

type MarketData struct {
	prefix  string
	logger  *slog.Logger
	timeout time.Duration
	js      jetstream.JetStream

	quotePubErrs, domPubErrs, histogramPubErrs prometheus.Counter
}

func (m *MarketData) Publish(d *tradovate.MarketData) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	switch {
	case len(d.DOMs) > 0:
		if err := m.pubDOM(ctx, d.DOMs); err != nil {
			m.domPubErrs.Inc()
		}
	case len(d.Histograms) > 0:
		if err := m.pubHistogram(ctx, d.Histograms); err != nil {
			m.histogramPubErrs.Inc()
		}
	case len(d.Quotes) > 0:
		if err := m.pubQuotes(ctx, d.Quotes); err != nil {
			m.quotePubErrs.Inc()
		}
	}
}

func (m *MarketData) pubDOM(ctx context.Context, x []*tradovate.DOM) error {

	return nil
}

func (m *MarketData) pubHistogram(ctx context.Context, x []*tradovate.Histogram) error {
	h := make([]*marketdata.Histogram, len(x))
	for i, v := range x {
		h[i] = ton.HistogramProtoV0(v)
	}

	buf, err := proto.Marshal(&marketdata.Histograms{Histograms: h})
	if err != nil {
		m.logger.ErrorContext(ctx, "failed marshal of histograms", "err", err, "raw", x)
		return err
	}

	_, err = m.js.Publish(ctx, fmt.Sprintf("%s.histogram.%s", m.prefix), buf)
	if err != nil {
		m.logger.ErrorContext(ctx, "failed publishing histogram", "err", err)
	}
	return err
}

func (m *MarketData) pubQuotes(ctx context.Context, x []*tradovate.Quote) error {

	return nil
}
