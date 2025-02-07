package main

import (
	"context"
	"errors"

	"github.com/AnthonyHewins/ton/gen/go/ordersvc/v0"
	"github.com/AnthonyHewins/ton/pkg/ton"
	"github.com/AnthonyHewins/tradovate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func orderErrReasonToCode(x tradovate.OrderErrReason) codes.Code {
	switch x {
	case tradovate.OrderErrReasonAdvancedTrailingStopUnsupported,
		tradovate.OrderErrReasonUnsupported:
		return codes.Unimplemented
	case tradovate.OrderErrReasonExecutionProviderNotConfigured,
		tradovate.OrderErrReasonOtherExecutionRelated:
		return codes.FailedPrecondition
	case tradovate.OrderErrReasonAnotherCommandPending:
		return codes.Aborted
	case tradovate.OrderErrReasonAccountUnspecified,
		tradovate.OrderErrReasonBackMonthProhibited,
		tradovate.OrderErrReasonInvalidContract,
		tradovate.OrderErrReasonInvalidPrice,
		tradovate.OrderErrReasonNoQuote,
		tradovate.OrderErrReasonParentRejected,
		tradovate.OrderErrReasonLiquidationOnly,
		tradovate.OrderErrReasonLiquidationOnlyBeforeExpiration,
		tradovate.OrderErrReasonMaxOrderQtyIsNotSpecified,
		tradovate.OrderErrReasonMaxPosLimitMisconfigured,
		tradovate.OrderErrReasonTrailingStopNonOrderQtyModify:
		return codes.InvalidArgument
	case tradovate.OrderErrReasonExecutionProviderUnavailable,
		tradovate.OrderErrReasonTradingLocked,
		tradovate.OrderErrReasonAccountClosed:
		return codes.Unavailable
	case tradovate.OrderErrReasonMultipleAccountPlanRequired,
		tradovate.OrderErrReasonUnauthorized:
		return codes.PermissionDenied
	case tradovate.OrderErrReasonRiskCheckTimeout,
		tradovate.OrderErrReasonTooLate:
		return codes.DeadlineExceeded
	case tradovate.OrderErrReasonSessionClosed,
		tradovate.OrderErrReasonMaxOrderQtyLimitReached,
		tradovate.OrderErrReasonMaxTotalPosLimitReached,
		tradovate.OrderErrReasonNotEnoughLiquidity,
		tradovate.OrderErrReasonMaxPosLimitReached:
		return codes.ResourceExhausted
	default:
		return codes.Internal
	}
}

func orderErr(err error) error {
	if errors.Is(err, &tradovate.OrderErr{}) {
		x := err.(*tradovate.OrderErr)
		return status.Error(orderErrReasonToCode(x.Reason), x.Text)
	}

	return status.Error(codes.Internal, err.Error())
}

func (a *app) CreateOrder(ctx context.Context, req *ordersvc.CreateOrderRequest) (*ordersvc.CreateOrderResponse, error) {
	r := &tradovate.OrderReq{
		AccountSpec:    req.AccountSpec,
		AccountID:      int(req.AccountId),
		ClOrdId:        req.ClientOrderId,
		Action:         ton.ProtoV0Action(req.Action),
		Symbol:         req.Symbol,
		OrderQty:       req.MaxShow,
		OrderType:      ton.ProtoV0OrderType(req.OrderType),
		Price:          req.Price,
		StopPrice:      req.StopPrice,
		MaxShow:        req.MaxShow,
		PegDifference:  req.PegDifference,
		TimeInForce:    ton.ProtoV0Tif(req.TimeInForce),
		ExpireTime:     req.ExpireTime.AsTime(),
		Text:           req.Text,
		ActivationTime: req.ActivationTime.AsTime(),
		CustomTag50:    req.CustomTag_50,
		IsAutomated:    true,
	}

	id, err := a.ws.PlaceOrder(ctx, r)
	if err != nil {
		a.Logger.ErrorContext(ctx, "failed create order", "err", err, "req", req)
		return nil, orderErr(err)
	}

	a.Logger.DebugContext(ctx, "created order", "id", id)
	return &ordersvc.CreateOrderResponse{OrderId: int64(id)}, nil
}

func (a *app) CreateOcoOrder(ctx context.Context, req *ordersvc.CreateOcoOrderRequest) (*ordersvc.CreateOcoOrderResponse, error) {
	r := &tradovate.OcoReq{
		AccountSpec:    req.AccountSpec,
		AccountID:      uint(req.AccountId),
		ClOrdID:        req.ClientOrderId,
		Action:         ton.ProtoV0Action(req.Action),
		Symbol:         req.Symbol,
		OrderQty:       uint(req.MaxShow),
		OrderType:      ton.ProtoV0OrderType(req.OrderType),
		Price:          req.Price,
		StopPrice:      req.StopPrice,
		MaxShow:        req.MaxShow,
		PegDifference:  req.PegDifference,
		TimeInForce:    ton.ProtoV0Tif(req.Tif),
		ExpireTime:     req.ExpireTime.AsTime(),
		Text:           req.Text,
		ActivationTime: req.ActivationTime.AsTime(),
		CustomTag50:    req.CustomTag_50,
		Other:          ton.ProtoV0OtherOrder(req.Other),
		IsAutomated:    true,
	}

	resp, err := a.ws.OCO(ctx, r)
	if err != nil {
		a.Logger.ErrorContext(ctx, "failed placing oco", "req", r)
		return nil, orderErr(err)
	}

	a.Logger.DebugContext(ctx, "placed OCO", "resp", resp)
	return &ordersvc.CreateOcoOrderResponse{
		OrderId: uint64(resp.OrderID),
		OcoId:   uint64(resp.OcoID),
	}, nil
}

func (a *app) CreateOsoOrder(ctx context.Context, req *ordersvc.CreateOsoOrderRequest) (*ordersvc.CreateOsoOrderResponse, error) {
	r := &tradovate.OsoReq{
		AccountSpec:    req.AccountSpec,
		AccountID:      uint(req.AccountId),
		ClOrdID:        req.ClientOrderId,
		Action:         ton.ProtoV0Action(req.Action),
		Symbol:         req.Symbol,
		OrderQty:       uint(req.MaxShow),
		OrderType:      ton.ProtoV0OrderType(req.OrderType),
		Price:          req.Price,
		StopPrice:      req.StopPrice,
		MaxShow:        req.MaxShow,
		PegDifference:  req.PegDifference,
		TimeInForce:    ton.ProtoV0Tif(req.Tif),
		ExpireTime:     req.ExpireTime.AsTime(),
		Text:           req.Text,
		ActivationTime: req.ActivationTime.AsTime(),
		CustomTag50:    req.CustomTag_50,
		Bracket1:       ton.ProtoV0OtherOrder(req.Bracket1),
		Bracket2:       ton.ProtoV0OtherOrder(req.Bracket2),
		IsAutomated:    true,
	}

	resp, err := a.ws.OSO(ctx, r)
	if err != nil {
		a.Logger.ErrorContext(ctx, "failed placing oco", "req", r)
		return nil, orderErr(err)
	}

	a.Logger.DebugContext(ctx, "placed OCO", "resp", resp)
	return &ordersvc.CreateOsoOrderResponse{
		OrderId:    uint64(resp.OrderID),
		Bracket1Id: uint64(resp.Oso1ID),
		Brakcet2Id: uint64(resp.Oso2ID),
	}, nil
}
