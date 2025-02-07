package ton

import (
	"context"
	"errors"
	"time"

	"github.com/AnthonyHewins/ton/gen/go/ordersvc/v0"
	"github.com/AnthonyHewins/tradovate"
	"github.com/nats-io/nats.go/jetstream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrKeywatchStopped = errors.New("order keywatch stopped")

	ErrActionMissing = status.Error(codes.InvalidArgument, "action must be buy or sell")
	ErrSymbolMissing = status.Error(codes.InvalidArgument, "missing symbol")
	ErrQtyMissing    = status.Error(codes.InvalidArgument, "missing qty")
	ErrTypeMissing   = status.Error(codes.InvalidArgument, "missing order type")
)

type CreateOrderReq struct {
	AccountSpec    string
	AccountID      int
	ClOrdId        string
	Action         tradovate.Action
	Symbol         string
	OrderQty       uint32
	OrderType      tradovate.OrderType
	Price          float64
	StopPrice      float64
	MaxShow        uint32
	PegDifference  float64
	TIF            tradovate.Tif
	ExpireTime     time.Time
	Text           string
	ActivationTime time.Time
	CustomTag50    string
}

func (c *CreateOrderReq) Validate() error {
	switch {
	case c.Action == 0:
		return ErrActionMissing
	case c.Symbol == "":
		return ErrSymbolMissing
	case c.OrderQty == 0:
		return ErrQtyMissing
	case c.OrderType == 0:
		return ErrTypeMissing
	}

	return nil
}

func (c *CreateOrderReq) protoV0() *ordersvc.CreateOrderRequest {
	return &ordersvc.CreateOrderRequest{
		AccountSpec:    c.AccountSpec,
		AccountId:      int64(c.AccountID),
		ClientOrderId:  c.ClOrdId,
		Action:         actionProtoV0(c.Action),
		Symbol:         c.Symbol,
		OrderQty:       c.OrderQty,
		OrderType:      orderTypeProtoV0(c.OrderType),
		Price:          c.Price,
		StopPrice:      c.StopPrice,
		MaxShow:        c.MaxShow,
		PegDifference:  c.PegDifference,
		TimeInForce:    tifProtoV0(c.TIF),
		ExpireTime:     timestamppb.New(c.ExpireTime),
		Text:           c.Text,
		ActivationTime: timestamppb.New(c.ActivationTime),
		CustomTag_50:   c.CustomTag50,
	}
}

type OrdersClient struct {
	client ordersvc.OrderServiceClient
	kv     jetstream.KeyValue
}

func NewOrdersClient(conn *grpc.ClientConn, kv jetstream.KeyValue) *OrdersClient {
	return &OrdersClient{client: ordersvc.NewOrderServiceClient(conn), kv: kv}
}

func (o *OrdersClient) Create(ctx context.Context, req *CreateOrderReq, opts ...grpc.CallOption) (int64, error) {
	if err := req.Validate(); err != nil {
		return 0, err
	}

	resp, err := o.client.CreateOrder(ctx, req.protoV0(), opts...)
	if err != nil {
		return 0, err
	}

	return resp.OrderId, nil
}

// func (o *OrdersClient) WatchOrders(ctx context.Context, fn func(*tradovate.Order), onError func(error)) ([]*tradovate.Order, error) {
// 	kw, err := o.kv.WatchAll(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	initialOrders := []*tradovate.Order{}
// 	for o := range kw.Updates() {
// 		if o == nil {
// 			break
// 		}

// 		var x ordersvc.Order
// 		if err := proto.Unmarshal(o.Value(), &x); err != nil {
// 			return nil, err
// 		}

// 		initialOrders = append(initialOrders, ProtoV0Order(&x))
// 	}

// 	go func() {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				go onError(ctx.Err())
// 				return
// 			case update := <-kw.Updates():
// 				if update == nil {
// 					go onError(ErrKeywatchStopped)
// 					return
// 				}

// 				var x ordersvc.Order
// 				if err := proto.Unmarshal(update.Value(), &x); err != nil {
// 					go onError(err)
// 					continue
// 				}

// 				go fn(ProtoV0Order(&x))
// 			}
// 		}
// 	}()

// 	return initialOrders, nil
// }
