package ton

import (
	"context"

	"github.com/AnthonyHewins/ton/gen/go/ordersvc/v0"
	"github.com/AnthonyHewins/tradovate"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OsoReq struct {
	Base               CreateOrderReq
	Bracket1, Bracket2 *tradovate.OtherOrder
}

func (o *OsoReq) proto() *ordersvc.CreateOsoOrderRequest {
	return &ordersvc.CreateOsoOrderRequest{
		ClientOrderId:  o.Base.ClientID,
		Action:         actionProtoV0(o.Base.Action),
		Symbol:         o.Base.Symbol,
		OrderQty:       o.Base.OrderQty,
		OrderType:      orderTypeProtoV0(o.Base.OrderType),
		Price:          o.Base.Price,
		StopPrice:      o.Base.StopPrice,
		MaxShow:        o.Base.MaxShow,
		PegDifference:  o.Base.PegDifference,
		Tif:            tifProtoV0(o.Base.TIF),
		ExpireTime:     timestamppb.New(o.Base.ExpireTime),
		Text:           o.Base.Text,
		ActivationTime: timestamppb.New(o.Base.ActivationTime),
		CustomTag_50:   o.Base.CustomTag50,
		Bracket_1:      otherOrderProto(o.Bracket1),
		Bracket_2:      otherOrderProto(o.Bracket2),
	}
}

func (o *OrdersClient) CreateOSO(ctx context.Context, req *OsoReq, opts ...grpc.CallOption) (*tradovate.OsoResp, error) {
	resp, err := o.client.CreateOsoOrder(ctx, req.proto(), opts...)
	if err != nil {
		return nil, err
	}

	return &tradovate.OsoResp{
		OrderID: uint(resp.OrderId),
		Oso1ID:  uint(resp.Bracket1Id),
		Oso2ID:  uint(resp.Brakcet2Id),
	}, nil
}
