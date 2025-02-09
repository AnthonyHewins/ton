package ingest

import (
	"fmt"

	"github.com/AnthonyHewins/ton/gen/go/ordersvc/v0"
	"github.com/AnthonyHewins/ton/gen/go/positionpb/v0"
	"github.com/AnthonyHewins/ton/pkg/ton"
	"github.com/AnthonyHewins/tradovate"
	"github.com/prometheus/client_golang/prometheus"
)

type orderTranslator struct {
	putErrs, delErrs prometheus.Counter
}

func (o orderTranslator) fromEntity(m *tradovate.EntityMsg) (*tradovate.Order, error) {
	return m.Order()
}

func (o orderTranslator) proto(x *tradovate.Order) *ordersvc.Order {
	return ton.OrderProtoV0(x)
}

func (o orderTranslator) id(x *tradovate.Order) string     { return fmt.Sprintf("order.%d", x.ID) }
func (o orderTranslator) putErrMetric() prometheus.Counter { return o.putErrs }
func (o orderTranslator) delErrMetric() prometheus.Counter { return o.putErrs }

type positionTranslator struct {
	putErrs, delErrs prometheus.Counter
}

func (p positionTranslator) fromEntity(x *tradovate.EntityMsg) (*tradovate.Position, error) {
	return x.Position()
}

func (p positionTranslator) proto(x *tradovate.Position) *positionpb.Position {
	return ton.PositionProtoV0(x)
}

func (p positionTranslator) id(x *tradovate.Position) string  { return fmt.Sprintf("position.%s", x.ID) }
func (o positionTranslator) putErrMetric() prometheus.Counter { return o.putErrs }
func (o positionTranslator) delErrMetric() prometheus.Counter { return o.putErrs }
