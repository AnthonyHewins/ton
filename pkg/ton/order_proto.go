package ton

import (
	"github.com/AnthonyHewins/ton/gen/go/ordersvc/v0"
	"github.com/AnthonyHewins/tradovate"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func OrderProtoV0(o *tradovate.Order) *ordersvc.Order {
	return &ordersvc.Order{
		Id:                  int64(o.ID),
		AccountId:           uint64(o.AccountID),
		ContractId:          uint64(o.ContractID),
		SpreadDefinitionId:  uint64(o.SpreadDefinitionID),
		Timestamp:           timestamppb.New(o.Timestamp),
		Action:              actionProtoV0(o.Action),
		Status:              statusProtoV0(o.Status),
		ExecutionProviderId: uint64(o.ExecutionProviderID),
		OcoId:               uint64(o.OcoID),
		ParentId:            uint64(o.ParentID),
		LinkedId:            uint64(o.LinkedID),
		Admin:               o.Admin,
	}
}

func ProtoV0Order(o *ordersvc.Order) *tradovate.Order {
	return &tradovate.Order{
		ID:                  uint(o.Id),
		AccountID:           uint(o.AccountId),
		ContractID:          uint(o.ContractId),
		SpreadDefinitionID:  uint(o.SpreadDefinitionId),
		Timestamp:           o.Timestamp.AsTime(),
		Action:              ProtoV0Action(o.Action),
		Status:              protoV0Status(o.Status),
		ExecutionProviderID: uint(o.ExecutionProviderId),
		OcoID:               uint(o.OcoId),
		ParentID:            uint(o.ParentId),
		LinkedID:            uint(o.LinkedId),
		Admin:               o.Admin,
	}
}

func ProtoV0OtherOrder(o *ordersvc.OtherOrder) *tradovate.OtherOrder {
	return &tradovate.OtherOrder{
		Action:        ProtoV0Action(o.Action),
		ClOrdID:       o.ClientOrderId,
		OrderType:     ProtoV0OrderType(o.OrderType),
		Price:         o.Price,
		StopPrice:     o.StopPrice,
		MaxShow:       o.MaxShow,
		PegDifference: o.PegDifference,
		TimeInForce:   ProtoV0Tif(o.Tif),
		ExpireTime:    o.ExpireTime.AsTime(),
		Text:          o.Text,
	}
}

func OtherOrderProtoV0(o *tradovate.OtherOrder) *ordersvc.OtherOrder {
	return &ordersvc.OtherOrder{
		Action:        actionProtoV0(o.Action),
		ClientOrderId: o.ClOrdID,
		OrderType:     orderTypeProtoV0(o.OrderType),
		Price:         o.Price,
		StopPrice:     o.StopPrice,
		MaxShow:       o.MaxShow,
		PegDifference: o.PegDifference,
		Tif:           tifProtoV0(o.TimeInForce),
		ExpireTime:    timestamppb.New(o.ExpireTime),
		Text:          o.Text,
	}
}

func actionProtoV0(action tradovate.Action) ordersvc.Action {
	switch action {
	case tradovate.ActionBuy:
		return ordersvc.Action_ACTION_BUY
	case tradovate.ActionSell:
		return ordersvc.Action_ACTION_SELL
	default:
		return ordersvc.Action_ACTION_UNSPECIFIED
	}
}

func ProtoV0Action(action ordersvc.Action) tradovate.Action {
	switch action {
	case ordersvc.Action_ACTION_BUY:
		return tradovate.ActionBuy
	case ordersvc.Action_ACTION_SELL:
		return tradovate.ActionSell
	default:
		return tradovate.ActionUnspecified
	}
}

func statusProtoV0(status tradovate.OrderStatus) ordersvc.OrderStatus {
	switch status {
	case tradovate.OrderStatusCanceled:
		return ordersvc.OrderStatus_ORDER_STATUS_CANCELED
	case tradovate.OrderStatusCompleted:
		return ordersvc.OrderStatus_ORDER_STATUS_COMPLETED
	case tradovate.OrderStatusExpired:
		return ordersvc.OrderStatus_ORDER_STATUS_EXPIRED
	case tradovate.OrderStatusFilled:
		return ordersvc.OrderStatus_ORDER_STATUS_FILLED
	case tradovate.OrderStatusPendingCancel:
		return ordersvc.OrderStatus_ORDER_STATUS_PENDING_CANCEL
	case tradovate.OrderStatusPendingNew:
		return ordersvc.OrderStatus_ORDER_STATUS_PENDING_NEW
	case tradovate.OrderStatusPendingReplace:
		return ordersvc.OrderStatus_ORDER_STATUS_PENDING_REPLACE
	case tradovate.OrderStatusRejected:
		return ordersvc.OrderStatus_ORDER_STATUS_REJECTED
	case tradovate.OrderStatusSuspended:
		return ordersvc.OrderStatus_ORDER_STATUS_SUSPENDED
	case tradovate.OrderStatusWorking:
		return ordersvc.OrderStatus_ORDER_STATUS_WORKING
	default:
		return ordersvc.OrderStatus_ORDER_STATUS_UNSPECIFIED
	}
}

func protoV0Status(status ordersvc.OrderStatus) tradovate.OrderStatus {
	switch status {
	case ordersvc.OrderStatus_ORDER_STATUS_CANCELED:
		return tradovate.OrderStatusCanceled
	case ordersvc.OrderStatus_ORDER_STATUS_COMPLETED:
		return tradovate.OrderStatusCompleted
	case ordersvc.OrderStatus_ORDER_STATUS_EXPIRED:
		return tradovate.OrderStatusExpired
	case ordersvc.OrderStatus_ORDER_STATUS_FILLED:
		return tradovate.OrderStatusFilled
	case ordersvc.OrderStatus_ORDER_STATUS_PENDING_CANCEL:
		return tradovate.OrderStatusPendingCancel
	case ordersvc.OrderStatus_ORDER_STATUS_PENDING_NEW:
		return tradovate.OrderStatusPendingNew
	case ordersvc.OrderStatus_ORDER_STATUS_PENDING_REPLACE:
		return tradovate.OrderStatusPendingReplace
	case ordersvc.OrderStatus_ORDER_STATUS_REJECTED:
		return tradovate.OrderStatusRejected
	case ordersvc.OrderStatus_ORDER_STATUS_SUSPENDED:
		return tradovate.OrderStatusSuspended
	case ordersvc.OrderStatus_ORDER_STATUS_WORKING:
		return tradovate.OrderStatusWorking
	default:
		return tradovate.OrderStatusUnknown
	}
}

func orderTypeProtoV0(orderType tradovate.OrderType) ordersvc.OrderType {
	switch orderType {
	case tradovate.OrderTypeLimit:
		return ordersvc.OrderType_ORDER_TYPE_LIMIT
	case tradovate.OrderTypeMIT:
		return ordersvc.OrderType_ORDER_TYPE_MIT
	case tradovate.OrderTypeMarket:
		return ordersvc.OrderType_ORDER_TYPE_MARKET
	case tradovate.OrderTypeQTS:
		return ordersvc.OrderType_ORDER_TYPE_QTS
	case tradovate.OrderTypeStop:
		return ordersvc.OrderType_ORDER_TYPE_STOP
	case tradovate.OrderTypeStopLimit:
		return ordersvc.OrderType_ORDER_TYPE_STOPLIMIT
	case tradovate.OrderTypeTrailingStop:
		return ordersvc.OrderType_ORDER_TYPE_TRAILINGSTOP
	case tradovate.OrderTypeTrailingStopLimit:
		return ordersvc.OrderType_ORDER_TYPE_TRAILINGSTOPLIMIT
	default:
		return ordersvc.OrderType_ORDER_TYPE_UNSPECIFIED
	}
}

func ProtoV0OrderType(orderType ordersvc.OrderType) tradovate.OrderType {
	switch orderType {
	case ordersvc.OrderType_ORDER_TYPE_LIMIT:
		return tradovate.OrderTypeLimit
	case ordersvc.OrderType_ORDER_TYPE_MIT:
		return tradovate.OrderTypeMIT
	case ordersvc.OrderType_ORDER_TYPE_MARKET:
		return tradovate.OrderTypeMarket
	case ordersvc.OrderType_ORDER_TYPE_QTS:
		return tradovate.OrderTypeQTS
	case ordersvc.OrderType_ORDER_TYPE_STOP:
		return tradovate.OrderTypeStop
	case ordersvc.OrderType_ORDER_TYPE_STOPLIMIT:
		return tradovate.OrderTypeStopLimit
	case ordersvc.OrderType_ORDER_TYPE_TRAILINGSTOP:
		return tradovate.OrderTypeTrailingStop
	case ordersvc.OrderType_ORDER_TYPE_TRAILINGSTOPLIMIT:
		return tradovate.OrderTypeTrailingStopLimit
	default:
		return tradovate.OrderTypeUnspecified
	}
}

func tifProtoV0(tif tradovate.Tif) ordersvc.TIF {
	switch tif {
	case tradovate.TifDay:
		return ordersvc.TIF_TIF_DAY
	case tradovate.TifFOK:
		return ordersvc.TIF_TIF_FOK
	case tradovate.TifGTC:
		return ordersvc.TIF_TIF_GTC
	case tradovate.TifGTD:
		return ordersvc.TIF_TIF_GTD
	case tradovate.TifIOC:
		return ordersvc.TIF_TIF_IOC
	default:
		return ordersvc.TIF_TIF_UNSPECIFIED
	}
}

func ProtoV0Tif(tif ordersvc.TIF) tradovate.Tif {
	switch tif {
	case ordersvc.TIF_TIF_DAY:
		return tradovate.TifDay
	case ordersvc.TIF_TIF_FOK:
		return tradovate.TifFOK
	case ordersvc.TIF_TIF_GTC:
		return tradovate.TifGTC
	case ordersvc.TIF_TIF_GTD:
		return tradovate.TifGTD
	case ordersvc.TIF_TIF_IOC:
		return tradovate.TifIOC
	default:
		return tradovate.TifUnspecified
	}
}
