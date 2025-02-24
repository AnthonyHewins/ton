package ton

import (
	"context"
	"errors"
	"fmt"

	"github.com/AnthonyHewins/ton/gen/go/ordersvc/v0"
	"github.com/AnthonyHewins/tradovate"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var ErrMissingBracket = errors.New("missing bracket")

type OcoReq struct {
	Base    CreateOrderReq
	Bracket *tradovate.OtherOrder
}

func (o *OcoReq) Validate() error {
	if err := o.Base.Validate(); err != nil {
		return err
	}

	if err := validateOther(o.Bracket); err != nil {
		return fmt.Errorf("error in other order: %w", err)
	}

	return nil
}

func validateOther(o *tradovate.OtherOrder) error {
	if o == nil {
		return ErrMissingBracket
	}

	switch {
	case o.Action == tradovate.ActionUnspecified:
		return ErrActionMissing
	case o.OrderType == tradovate.OrderTypeUnspecified:
		return ErrTypeMissing
	default:
		return nil
	}
}

func otherOrderProto(c *tradovate.OtherOrder) *ordersvc.OtherOrder {
	return &ordersvc.OtherOrder{
		Action:        actionProtoV0(c.Action),
		ClientOrderId: c.ClOrdID,
		OrderType:     orderTypeProtoV0(c.OrderType),
		Price:         c.Price,
		StopPrice:     c.StopPrice,
		MaxShow:       c.MaxShow,
		PegDifference: c.PegDifference,
		Tif:           tifProtoV0(c.TimeInForce),
		ExpireTime:    timestamppb.New(c.ExpireTime),
		Text:          c.Text,
	}
}

func (c *OcoReq) proto() *ordersvc.CreateOcoOrderRequest {
	return &ordersvc.CreateOcoOrderRequest{
		ClientOrderId:  c.Base.ClientID,
		Action:         actionProtoV0(c.Base.Action),
		Symbol:         c.Base.Symbol,
		OrderQty:       c.Base.OrderQty,
		OrderType:      orderTypeProtoV0(c.Base.OrderType),
		Price:          c.Base.Price,
		StopPrice:      c.Base.StopPrice,
		MaxShow:        c.Base.MaxShow,
		PegDifference:  c.Base.PegDifference,
		Tif:            tifProtoV0(c.Base.TIF),
		ExpireTime:     timestamppb.New(c.Base.ExpireTime),
		Text:           c.Base.Text,
		ActivationTime: timestamppb.New(c.Base.ActivationTime),
		CustomTag_50:   c.Base.CustomTag50,
		Other:          &ordersvc.OtherOrder{},
	}
}

func (o *OrdersClient) CreateOCO(ctx context.Context, req *OcoReq, opts ...grpc.CallOption) (*tradovate.OcoResp, error) {
	resp, err := o.client.CreateOcoOrder(ctx, req.proto(), opts...)
	if err != nil {
		return nil, err
	}

	return &tradovate.OcoResp{OrderID: uint(resp.OrderId), OcoID: uint(resp.OcoId)}, nil
}
