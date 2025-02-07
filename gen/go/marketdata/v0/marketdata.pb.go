// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: marketdata/v0/marketdata.proto

package marketdata

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Open        float64                `protobuf:"fixed64,2,opt,name=open,proto3" json:"open,omitempty"`
	High        float64                `protobuf:"fixed64,3,opt,name=high,proto3" json:"high,omitempty"`
	Low         float64                `protobuf:"fixed64,4,opt,name=low,proto3" json:"low,omitempty"`
	Close       float64                `protobuf:"fixed64,5,opt,name=close,proto3" json:"close,omitempty"`
	UpVolume    float64                `protobuf:"fixed64,6,opt,name=up_volume,json=upVolume,proto3" json:"up_volume,omitempty"`
	DownVolume  float64                `protobuf:"fixed64,7,opt,name=down_volume,json=downVolume,proto3" json:"down_volume,omitempty"`
	UpTicks     float64                `protobuf:"fixed64,8,opt,name=up_ticks,json=upTicks,proto3" json:"up_ticks,omitempty"`
	DownTicks   float64                `protobuf:"fixed64,9,opt,name=down_ticks,json=downTicks,proto3" json:"down_ticks,omitempty"`
	BidVolume   float64                `protobuf:"fixed64,10,opt,name=bid_volume,json=bidVolume,proto3" json:"bid_volume,omitempty"`
	OfferVolume float64                `protobuf:"fixed64,11,opt,name=offer_volume,json=offerVolume,proto3" json:"offer_volume,omitempty"`
}

func (x *Bar) Reset() {
	*x = Bar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketdata_v0_marketdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bar) ProtoMessage() {}

func (x *Bar) ProtoReflect() protoreflect.Message {
	mi := &file_marketdata_v0_marketdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bar.ProtoReflect.Descriptor instead.
func (*Bar) Descriptor() ([]byte, []int) {
	return file_marketdata_v0_marketdata_proto_rawDescGZIP(), []int{0}
}

func (x *Bar) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Bar) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *Bar) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *Bar) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *Bar) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *Bar) GetUpVolume() float64 {
	if x != nil {
		return x.UpVolume
	}
	return 0
}

func (x *Bar) GetDownVolume() float64 {
	if x != nil {
		return x.DownVolume
	}
	return 0
}

func (x *Bar) GetUpTicks() float64 {
	if x != nil {
		return x.UpTicks
	}
	return 0
}

func (x *Bar) GetDownTicks() float64 {
	if x != nil {
		return x.DownTicks
	}
	return 0
}

func (x *Bar) GetBidVolume() float64 {
	if x != nil {
		return x.BidVolume
	}
	return 0
}

func (x *Bar) GetOfferVolume() float64 {
	if x != nil {
		return x.OfferVolume
	}
	return 0
}

type Tick struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RelativeTime     int64   `protobuf:"varint,2,opt,name=relative_time,json=relativeTime,proto3" json:"relative_time,omitempty"`
	RelativePrice    int64   `protobuf:"varint,3,opt,name=relative_price,json=relativePrice,proto3" json:"relative_price,omitempty"`
	Volume           int64   `protobuf:"varint,4,opt,name=volume,proto3" json:"volume,omitempty"`
	RelativeBidPrice float64 `protobuf:"fixed64,5,opt,name=relative_bid_price,json=relativeBidPrice,proto3" json:"relative_bid_price,omitempty"`
	RelativeAskPrice float64 `protobuf:"fixed64,6,opt,name=relative_ask_price,json=relativeAskPrice,proto3" json:"relative_ask_price,omitempty"`
	BidSize          float64 `protobuf:"fixed64,7,opt,name=bid_size,json=bidSize,proto3" json:"bid_size,omitempty"`
	AskSize          float64 `protobuf:"fixed64,8,opt,name=ask_size,json=askSize,proto3" json:"ask_size,omitempty"`
}

func (x *Tick) Reset() {
	*x = Tick{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketdata_v0_marketdata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tick) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tick) ProtoMessage() {}

func (x *Tick) ProtoReflect() protoreflect.Message {
	mi := &file_marketdata_v0_marketdata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tick.ProtoReflect.Descriptor instead.
func (*Tick) Descriptor() ([]byte, []int) {
	return file_marketdata_v0_marketdata_proto_rawDescGZIP(), []int{1}
}

func (x *Tick) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tick) GetRelativeTime() int64 {
	if x != nil {
		return x.RelativeTime
	}
	return 0
}

func (x *Tick) GetRelativePrice() int64 {
	if x != nil {
		return x.RelativePrice
	}
	return 0
}

func (x *Tick) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Tick) GetRelativeBidPrice() float64 {
	if x != nil {
		return x.RelativeBidPrice
	}
	return 0
}

func (x *Tick) GetRelativeAskPrice() float64 {
	if x != nil {
		return x.RelativeAskPrice
	}
	return 0
}

func (x *Tick) GetBidSize() float64 {
	if x != nil {
		return x.BidSize
	}
	return 0
}

func (x *Tick) GetAskSize() float64 {
	if x != nil {
		return x.AskSize
	}
	return 0
}

type TickChart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TradeDate     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=trade_date,json=tradeDate,proto3" json:"trade_date,omitempty"`
	EndOfHistory  bool                   `protobuf:"varint,3,opt,name=end_of_history,json=endOfHistory,proto3" json:"end_of_history,omitempty"`
	Source        string                 `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	BasePrice     int64                  `protobuf:"varint,5,opt,name=base_price,json=basePrice,proto3" json:"base_price,omitempty"`
	BaseTimestamp *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=base_timestamp,json=baseTimestamp,proto3" json:"base_timestamp,omitempty"`
	TickSize      float64                `protobuf:"fixed64,7,opt,name=tick_size,json=tickSize,proto3" json:"tick_size,omitempty"`
	Ticks         []*Tick                `protobuf:"bytes,8,rep,name=ticks,proto3" json:"ticks,omitempty"`
}

func (x *TickChart) Reset() {
	*x = TickChart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketdata_v0_marketdata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TickChart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TickChart) ProtoMessage() {}

func (x *TickChart) ProtoReflect() protoreflect.Message {
	mi := &file_marketdata_v0_marketdata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TickChart.ProtoReflect.Descriptor instead.
func (*TickChart) Descriptor() ([]byte, []int) {
	return file_marketdata_v0_marketdata_proto_rawDescGZIP(), []int{2}
}

func (x *TickChart) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TickChart) GetTradeDate() *timestamppb.Timestamp {
	if x != nil {
		return x.TradeDate
	}
	return nil
}

func (x *TickChart) GetEndOfHistory() bool {
	if x != nil {
		return x.EndOfHistory
	}
	return false
}

func (x *TickChart) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *TickChart) GetBasePrice() int64 {
	if x != nil {
		return x.BasePrice
	}
	return 0
}

func (x *TickChart) GetBaseTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.BaseTimestamp
	}
	return nil
}

func (x *TickChart) GetTickSize() float64 {
	if x != nil {
		return x.TickSize
	}
	return 0
}

func (x *TickChart) GetTicks() []*Tick {
	if x != nil {
		return x.Ticks
	}
	return nil
}

type BarChart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TradeDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=trade_date,json=tradeDate,proto3" json:"trade_date,omitempty"`
	Bars      []*Bar                 `protobuf:"bytes,3,rep,name=bars,proto3" json:"bars,omitempty"`
}

func (x *BarChart) Reset() {
	*x = BarChart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketdata_v0_marketdata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarChart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarChart) ProtoMessage() {}

func (x *BarChart) ProtoReflect() protoreflect.Message {
	mi := &file_marketdata_v0_marketdata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarChart.ProtoReflect.Descriptor instead.
func (*BarChart) Descriptor() ([]byte, []int) {
	return file_marketdata_v0_marketdata_proto_rawDescGZIP(), []int{3}
}

func (x *BarChart) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BarChart) GetTradeDate() *timestamppb.Timestamp {
	if x != nil {
		return x.TradeDate
	}
	return nil
}

func (x *BarChart) GetBars() []*Bar {
	if x != nil {
		return x.Bars
	}
	return nil
}

var File_marketdata_v0_marketdata_proto protoreflect.FileDescriptor

var file_marketdata_v0_marketdata_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x30, 0x2f,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x30, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc9, 0x02, 0x0a, 0x03, 0x42, 0x61, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f,
	0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x75, 0x70, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x64, 0x6f, 0x77, 0x6e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x75, 0x70, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x07, 0x75, 0x70, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x6f, 0x77, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x64, 0x6f, 0x77, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69,
	0x64, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x62, 0x69, 0x64, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x66, 0x66,
	0x65, 0x72, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0b, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0x8c, 0x02, 0x0a,
	0x04, 0x54, 0x69, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x62, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x42,
	0x69, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x10, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x41, 0x73, 0x6b,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x69, 0x64, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x62, 0x69, 0x64, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x07, 0x61, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xbe, 0x02, 0x0a, 0x09,
	0x54, 0x69, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x66, 0x5f, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x65, 0x6e,
	0x64, 0x4f, 0x66, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x41, 0x0a, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x29, 0x0a, 0x05, 0x74, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x30,
	0x2e, 0x54, 0x69, 0x63, 0x6b, 0x52, 0x05, 0x74, 0x69, 0x63, 0x6b, 0x73, 0x22, 0x7d, 0x0a, 0x08,
	0x42, 0x61, 0x72, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x72, 0x61, 0x64, 0x65, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x62, 0x61, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76,
	0x30, 0x2e, 0x42, 0x61, 0x72, 0x52, 0x04, 0x62, 0x61, 0x72, 0x73, 0x42, 0x3e, 0x5a, 0x3c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x74, 0x68, 0x6f, 0x6e,
	0x79, 0x48, 0x65, 0x77, 0x69, 0x6e, 0x73, 0x2f, 0x74, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x30,
	0x3b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_marketdata_v0_marketdata_proto_rawDescOnce sync.Once
	file_marketdata_v0_marketdata_proto_rawDescData = file_marketdata_v0_marketdata_proto_rawDesc
)

func file_marketdata_v0_marketdata_proto_rawDescGZIP() []byte {
	file_marketdata_v0_marketdata_proto_rawDescOnce.Do(func() {
		file_marketdata_v0_marketdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_marketdata_v0_marketdata_proto_rawDescData)
	})
	return file_marketdata_v0_marketdata_proto_rawDescData
}

var file_marketdata_v0_marketdata_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_marketdata_v0_marketdata_proto_goTypes = []interface{}{
	(*Bar)(nil),                   // 0: marketdata.v0.Bar
	(*Tick)(nil),                  // 1: marketdata.v0.Tick
	(*TickChart)(nil),             // 2: marketdata.v0.TickChart
	(*BarChart)(nil),              // 3: marketdata.v0.BarChart
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_marketdata_v0_marketdata_proto_depIdxs = []int32{
	4, // 0: marketdata.v0.Bar.timestamp:type_name -> google.protobuf.Timestamp
	4, // 1: marketdata.v0.TickChart.trade_date:type_name -> google.protobuf.Timestamp
	4, // 2: marketdata.v0.TickChart.base_timestamp:type_name -> google.protobuf.Timestamp
	1, // 3: marketdata.v0.TickChart.ticks:type_name -> marketdata.v0.Tick
	4, // 4: marketdata.v0.BarChart.trade_date:type_name -> google.protobuf.Timestamp
	0, // 5: marketdata.v0.BarChart.bars:type_name -> marketdata.v0.Bar
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_marketdata_v0_marketdata_proto_init() }
func file_marketdata_v0_marketdata_proto_init() {
	if File_marketdata_v0_marketdata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_marketdata_v0_marketdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bar); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_marketdata_v0_marketdata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tick); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_marketdata_v0_marketdata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TickChart); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_marketdata_v0_marketdata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarChart); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_marketdata_v0_marketdata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_marketdata_v0_marketdata_proto_goTypes,
		DependencyIndexes: file_marketdata_v0_marketdata_proto_depIdxs,
		MessageInfos:      file_marketdata_v0_marketdata_proto_msgTypes,
	}.Build()
	File_marketdata_v0_marketdata_proto = out.File
	file_marketdata_v0_marketdata_proto_rawDesc = nil
	file_marketdata_v0_marketdata_proto_goTypes = nil
	file_marketdata_v0_marketdata_proto_depIdxs = nil
}
