package stream

import (
	"fmt"
	"log/slog"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/prometheus/client_golang/prometheus"
)

type Consumer struct {
	logger    *slog.Logger
	ackFails  prometheus.Counter
	nakFails  prometheus.Counter
	termFails prometheus.Counter
}

func NewConsumer(app, subsystem string, logger *slog.Logger) *Consumer {
	return &Consumer{
		logger: logger.With("app", app, "subsystem", subsystem),
		ackFails: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: app,
			Subsystem: subsystem,
			Name:      "ack_fails",
			Help:      "Number of times a jetstream ACK failed in this jetstream consumer",
		}),
		nakFails: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: app,
			Subsystem: subsystem,
			Name:      "nak_fails",
			Help:      "Number of times a jetstream NAK failed in this jetstream consumer",
		}),
		termFails: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: app,
			Subsystem: subsystem,
			Name:      "term_fails",
			Help:      "Number of times a jetstream TERM failed in this jetstream consumer",
		}),
	}
}

func (c *Consumer) ack(m jetstream.Msg) error {
	err := m.Ack()
	if err != nil {
		c.logger.Error("failed ack of msg", "err", err, "subj", m.Subject())
		c.ackFails.Inc()
	}
	return err
}

func (c *Consumer) nak(m jetstream.Msg) error {
	err := m.Nak()
	if err != nil {
		c.logger.Error("failed nak of msg", "err", err, "subj", m.Subject())
		c.nakFails.Inc()
	}
	return err
}

func (c *Consumer) term(m jetstream.Msg, reason string) error {
	err := m.TermWithReason(reason)
	if err != nil {
		c.logger.Error("failed nak of msg", "err", err, "reason", reason, "subj", m.Subject())
		c.termFails.Inc()
	}
	return err
}

type MsgAction byte

const (
	Ack MsgAction = iota + 1
	Nak
	Term
)

type Action struct {
	MsgAction
	Reason string
}

func (c *Consumer) MsgHandler(m jetstream.Msg, x Action) {
	switch x.MsgAction {
	case Ack:
		c.ack(m)
	case Nak:
		c.nak(m)
	case Term:
		c.term(m, x.Reason)
	default:
		panic(fmt.Sprintf("invalid action: %d", x.MsgAction))
	}
}
