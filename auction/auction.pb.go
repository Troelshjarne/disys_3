// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: auction/auction.proto

package auction

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BidMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientName string `protobuf:"bytes,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
	Amount     int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *BidMessage) Reset() {
	*x = BidMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_auction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidMessage) ProtoMessage() {}

func (x *BidMessage) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidMessage.ProtoReflect.Descriptor instead.
func (*BidMessage) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{0}
}

func (x *BidMessage) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

func (x *BidMessage) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// {fail, success or exception}
type MessageAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *MessageAck) Reset() {
	*x = MessageAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_auction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageAck) ProtoMessage() {}

func (x *MessageAck) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageAck.ProtoReflect.Descriptor instead.
func (*MessageAck) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{1}
}

func (x *MessageAck) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ResultMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ongoing       bool   `protobuf:"varint,1,opt,name=ongoing,proto3" json:"ongoing,omitempty"`
	WinningClient string `protobuf:"bytes,2,opt,name=winningClient,proto3" json:"winningClient,omitempty"`
	HighestBid    int32  `protobuf:"varint,3,opt,name=highestBid,proto3" json:"highestBid,omitempty"`
}

func (x *ResultMessage) Reset() {
	*x = ResultMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_auction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultMessage) ProtoMessage() {}

func (x *ResultMessage) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultMessage.ProtoReflect.Descriptor instead.
func (*ResultMessage) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{2}
}

func (x *ResultMessage) GetOngoing() bool {
	if x != nil {
		return x.Ongoing
	}
	return false
}

func (x *ResultMessage) GetWinningClient() string {
	if x != nil {
		return x.WinningClient
	}
	return ""
}

func (x *ResultMessage) GetHighestBid() int32 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

type IpMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ips []string `protobuf:"bytes,1,rep,name=ips,proto3" json:"ips,omitempty"`
}

func (x *IpMessage) Reset() {
	*x = IpMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_auction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpMessage) ProtoMessage() {}

func (x *IpMessage) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpMessage.ProtoReflect.Descriptor instead.
func (*IpMessage) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{3}
}

func (x *IpMessage) GetIps() []string {
	if x != nil {
		return x.Ips
	}
	return nil
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FakeInfo bool `protobuf:"varint,1,opt,name=fakeInfo,proto3" json:"fakeInfo,omitempty"`
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_auction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{4}
}

func (x *Void) GetFakeInfo() bool {
	if x != nil {
		return x.FakeInfo
	}
	return false
}

var File_auction_auction_proto protoreflect.FileDescriptor

var file_auction_auction_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x44, 0x0a, 0x0a, 0x62, 0x69, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x24, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x41, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x6f, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x6f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x77, 0x69, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x77, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x22, 0x1d, 0x0a,
	0x09, 0x69, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x70,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x70, 0x73, 0x22, 0x22, 0x0a, 0x04,
	0x76, 0x6f, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x6b, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x61, 0x6b, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x32, 0x87, 0x02, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0d, 0x2e, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x6f, 0x69, 0x64, 0x1a, 0x16, 0x2e, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x12, 0x13, 0x2e, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x69, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x6f, 0x69, 0x64, 0x1a, 0x0d, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x6f, 0x69, 0x64, 0x1a, 0x0d, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x0d, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x6f, 0x69, 0x64, 0x1a, 0x12, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69,
	0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x54, 0x72, 0x6f, 0x65, 0x6c, 0x73, 0x68, 0x6a, 0x61, 0x72, 0x6e, 0x65, 0x2f, 0x64, 0x69,
	0x73, 0x79, 0x73, 0x5f, 0x33, 0x3b, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auction_auction_proto_rawDescOnce sync.Once
	file_auction_auction_proto_rawDescData = file_auction_auction_proto_rawDesc
)

func file_auction_auction_proto_rawDescGZIP() []byte {
	file_auction_auction_proto_rawDescOnce.Do(func() {
		file_auction_auction_proto_rawDescData = protoimpl.X.CompressGZIP(file_auction_auction_proto_rawDescData)
	})
	return file_auction_auction_proto_rawDescData
}

var file_auction_auction_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_auction_auction_proto_goTypes = []interface{}{
	(*BidMessage)(nil),    // 0: auction.bidMessage
	(*MessageAck)(nil),    // 1: auction.messageAck
	(*ResultMessage)(nil), // 2: auction.resultMessage
	(*IpMessage)(nil),     // 3: auction.ipMessage
	(*Void)(nil),          // 4: auction.void
}
var file_auction_auction_proto_depIdxs = []int32{
	4, // 0: auction.Communication.result:input_type -> auction.void
	0, // 1: auction.Communication.bid:input_type -> auction.bidMessage
	4, // 2: auction.Communication.startAuction:input_type -> auction.void
	4, // 3: auction.Communication.endAuction:input_type -> auction.void
	4, // 4: auction.Communication.getReplicas:input_type -> auction.void
	2, // 5: auction.Communication.result:output_type -> auction.resultMessage
	1, // 6: auction.Communication.bid:output_type -> auction.messageAck
	4, // 7: auction.Communication.startAuction:output_type -> auction.void
	4, // 8: auction.Communication.endAuction:output_type -> auction.void
	3, // 9: auction.Communication.getReplicas:output_type -> auction.ipMessage
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auction_auction_proto_init() }
func file_auction_auction_proto_init() {
	if File_auction_auction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auction_auction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidMessage); i {
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
		file_auction_auction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageAck); i {
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
		file_auction_auction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultMessage); i {
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
		file_auction_auction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpMessage); i {
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
		file_auction_auction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
			RawDescriptor: file_auction_auction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auction_auction_proto_goTypes,
		DependencyIndexes: file_auction_auction_proto_depIdxs,
		MessageInfos:      file_auction_auction_proto_msgTypes,
	}.Build()
	File_auction_auction_proto = out.File
	file_auction_auction_proto_rawDesc = nil
	file_auction_auction_proto_goTypes = nil
	file_auction_auction_proto_depIdxs = nil
}
