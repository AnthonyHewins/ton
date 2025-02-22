package main

import (
	"context"

	"github.com/AnthonyHewins/ton/gen/go/ordersvc/v0"
	"github.com/AnthonyHewins/ton/pkg/ton"
	"github.com/AnthonyHewins/tradovate"
)

func (a *app) CreateOrder(ctx context.Context, req *ordersvc.CreateOrderRequest) (*ordersvc.CreateOrderResponse, error) {
	id, err := a.controller.Orders.PlaceOrder(ctx, &tradovate.OrderReq{
		AccountSpec:    req.AccountSpec,
		AccountID:      int(req.AccountId),
		ClientOrderID:  req.ClientOrderId,
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
	})

	if err != nil {
		return nil, err
	}

	return &ordersvc.CreateOrderResponse{OrderId: int64(id)}, nil
}

func (a *app) CreateOcoOrder(ctx context.Context, req *ordersvc.CreateOcoOrderRequest) (*ordersvc.CreateOcoOrderResponse, error) {
	resp, err := a.controller.Orders.PlaceOCO(ctx, &tradovate.OcoReq{
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
	})

	if err != nil {
		return nil, err
	}

	return &ordersvc.CreateOcoOrderResponse{
		OrderId: uint64(resp.OrderID),
		OcoId:   uint64(resp.OcoID),
	}, nil
}

func (a *app) CreateOsoOrder(ctx context.Context, req *ordersvc.CreateOsoOrderRequest) (*ordersvc.CreateOsoOrderResponse, error) {
	resp, err := a.controller.Orders.PlaceOSO(ctx, &tradovate.OsoReq{
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
	})

	if err != nil {
		return nil, err
	}

	return &ordersvc.CreateOsoOrderResponse{
		OrderId:    uint64(resp.OrderID),
		Bracket1Id: uint64(resp.Oso1ID),
		Brakcet2Id: uint64(resp.Oso2ID),
	}, nil
}
