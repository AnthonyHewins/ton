package ingest

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/AnthonyHewins/tradovate"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/proto"
)

type translator[Proto proto.Message, Model any] interface {
	fromEntity(*tradovate.EntityMsg) (Model, error)
	proto(Model) Proto
	id(Model) string

	putErrMetric() prometheus.Counter
	delErrMetric() prometheus.Counter
}

type entityKV[Proto proto.Message, Model any] struct {
	logger *slog.Logger
	c      translator[Proto, Model]
	kv     jetstream.KeyValue
}

func (e *entityKV[Proto, Model]) publish(ctx context.Context, data *tradovate.EntityMsg) error {
	x, err := e.c.fromEntity(data)
	if err != nil {
		var want Model
		var wantProto Proto
		e.logger.ErrorContext(ctx,
			"failed casting entity msg to wanted type",
			"event", data.Event,
			"type", data.Type,
			"raw", data.Data,
			"wantProto", fmt.Sprintf("%T", wantProto),
			"want", fmt.Sprintf("%T", want),
		)
		return err
	}

	switch data.Event {
	case tradovate.EventTypeCreated, tradovate.EventTypeUpdated:
		if err = e.put(ctx, x); err != nil {
			e.c.putErrMetric().Inc()
		}
	case tradovate.EventTypeDeleted:
		if err = e.del(ctx, x); err != nil {
			e.c.delErrMetric().Inc()
		}
	default:
		e.logger.ErrorContext(ctx, "unknown event received for orderpub", "event", data.Event)
		return fmt.Errorf("unknown event: %s", data.Event)
	}

	return err
}

func (e *entityKV[X, Y]) put(ctx context.Context, y Y) error {
	buf, err := proto.Marshal(e.c.proto(y))
	if err != nil {
		e.logger.ErrorContext(ctx, "failed marshal of model", "model", y, "err", err)
		return err
	}

	key := e.c.id(y)
	if _, err = e.kv.Put(ctx, key, buf); err != nil {
		e.logger.ErrorContext(ctx, "failed PUT", "model", y, "err", err, "key", key)
		return err
	}

	return nil
}

func (e *entityKV[Proto, Model]) del(ctx context.Context, y Model) error {
	key := e.c.id(y)
	if err := e.kv.Delete(ctx, key); err != nil {
		e.logger.ErrorContext(ctx, "failed deleting model", "model", y, "err", err, "key", key)
		return err
	}

	return nil
}
