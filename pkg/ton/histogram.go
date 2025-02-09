package ton

import (
	"github.com/AnthonyHewins/ton/gen/go/marketdata/v0"
	"github.com/AnthonyHewins/tradovate"
	"github.com/nats-io/nats.go/jetstream"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const HistogramConsumerV0Subject = "ton-histogram-consumer-v0"

func JetstreamHistograms(e jetstream.Msg) ([]*tradovate.Histogram, error) {
	var x marketdata.Histograms
	if err := proto.Unmarshal(e.Data(), &x); err != nil {
		return nil, err
	}

	h := make([]*tradovate.Histogram, len(x.Histograms))
	for i, v := range x.Histograms {
		h[i] = ProtoV0Histogram(v)
	}

	return h, nil
}

func ProtoV0Histogram(x *marketdata.Histogram) *tradovate.Histogram {
	return &tradovate.Histogram{
		ContractID: int(x.ContractId),
		Timestamp:  x.Timestamp.AsTime(),
		TradeDate:  x.TradeDate.AsTime(),
		Base:       x.Base,
		Items:      x.Items,
		Refresh:    x.Refresh,
	}
}

func HistogramProtoV0(x *tradovate.Histogram) *marketdata.Histogram {
	return &marketdata.Histogram{
		ContractId: int64(x.ContractID),
		Timestamp:  timestamppb.New(x.Timestamp),
		TradeDate:  timestamppb.New(x.TradeDate),
		Base:       x.Base,
		Items:      x.Items,
		Refresh:    x.Refresh,
	}
}
