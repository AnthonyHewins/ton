package ton

import (
	"github.com/AnthonyHewins/ton/gen/go/ordersvc/v0"
	"github.com/nats-io/nats.go/jetstream"
	"google.golang.org/grpc"
)

type OrdersClient struct {
	client ordersvc.OrderServiceClient
	kv     jetstream.KeyValue
}

func NewOrdersClient(conn *grpc.ClientConn, kv jetstream.KeyValue) *OrdersClient {
	return &OrdersClient{client: ordersvc.NewOrderServiceClient(conn), kv: kv}
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
