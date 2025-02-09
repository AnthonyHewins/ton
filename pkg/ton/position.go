package ton

import (
	"github.com/AnthonyHewins/ton/gen/go/positionpb/v0"
	"github.com/AnthonyHewins/tradovate"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PositionProtoV0(x *tradovate.Position) *positionpb.Position {
	return &positionpb.Position{
		Id:          int64(x.ID),
		AccountId:   int64(x.AccountID),
		ContractId:  int64(x.ContractID),
		Timestamp:   timestamppb.New(x.Timestamp),
		TradeDate:   timestamppb.New(x.TradeDate),
		NetPos:      int64(x.NetPos),
		NetPrice:    x.NetPrice,
		Bought:      int64(x.Bought),
		BoughtValue: x.BoughtValue,
		Sold:        int64(x.Sold),
		SoldValue:   x.SoldValue,
		PrevPos:     int64(x.PrevPos),
		PrevPrice:   x.PrevPrice,
	}
}

func ProtoV0Position(x *positionpb.Position) *tradovate.Position {
	return &tradovate.Position{
		ID:          int(x.Id),
		AccountID:   int(x.AccountId),
		ContractID:  int(x.ContractId),
		Timestamp:   x.Timestamp.AsTime(),
		TradeDate:   x.TradeDate.AsTime(),
		NetPos:      int(x.NetPos),
		NetPrice:    x.NetPrice,
		Bought:      int(x.Bought),
		BoughtValue: x.BoughtValue,
		Sold:        int(x.Sold),
		SoldValue:   x.SoldValue,
		PrevPos:     int(x.PrevPos),
		PrevPrice:   x.PrevPrice,
	}
}
