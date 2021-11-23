// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package auction

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CommunicationClient is the client API for Communication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommunicationClient interface {
	Result(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ResultMessage, error)
	Bid(ctx context.Context, in *BidMessage, opts ...grpc.CallOption) (*MessageAck, error)
	// reset bids,clock, ongoing = true eg. all relevant fields.
	StartAuction(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error)
	// set ongoing = false eg. when times runs out
	EndAuction(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error)
	// list of ips of replicas (serf)
	GetReplicas(ctx context.Context, in *Void, opts ...grpc.CallOption) (*IpMessage, error)
}

type communicationClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunicationClient(cc grpc.ClientConnInterface) CommunicationClient {
	return &communicationClient{cc}
}

func (c *communicationClient) Result(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ResultMessage, error) {
	out := new(ResultMessage)
	err := c.cc.Invoke(ctx, "/auction.Communication/result", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationClient) Bid(ctx context.Context, in *BidMessage, opts ...grpc.CallOption) (*MessageAck, error) {
	out := new(MessageAck)
	err := c.cc.Invoke(ctx, "/auction.Communication/bid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationClient) StartAuction(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/auction.Communication/startAuction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationClient) EndAuction(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/auction.Communication/endAuction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationClient) GetReplicas(ctx context.Context, in *Void, opts ...grpc.CallOption) (*IpMessage, error) {
	out := new(IpMessage)
	err := c.cc.Invoke(ctx, "/auction.Communication/getReplicas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunicationServer is the server API for Communication service.
// All implementations must embed UnimplementedCommunicationServer
// for forward compatibility
type CommunicationServer interface {
	Result(context.Context, *Void) (*ResultMessage, error)
	Bid(context.Context, *BidMessage) (*MessageAck, error)
	// reset bids,clock, ongoing = true eg. all relevant fields.
	StartAuction(context.Context, *Void) (*Void, error)
	// set ongoing = false eg. when times runs out
	EndAuction(context.Context, *Void) (*Void, error)
	// list of ips of replicas (serf)
	GetReplicas(context.Context, *Void) (*IpMessage, error)
	mustEmbedUnimplementedCommunicationServer()
}

// UnimplementedCommunicationServer must be embedded to have forward compatible implementations.
type UnimplementedCommunicationServer struct {
}

func (UnimplementedCommunicationServer) Result(context.Context, *Void) (*ResultMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedCommunicationServer) Bid(context.Context, *BidMessage) (*MessageAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedCommunicationServer) StartAuction(context.Context, *Void) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartAuction not implemented")
}
func (UnimplementedCommunicationServer) EndAuction(context.Context, *Void) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndAuction not implemented")
}
func (UnimplementedCommunicationServer) GetReplicas(context.Context, *Void) (*IpMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReplicas not implemented")
}
func (UnimplementedCommunicationServer) mustEmbedUnimplementedCommunicationServer() {}

// UnsafeCommunicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommunicationServer will
// result in compilation errors.
type UnsafeCommunicationServer interface {
	mustEmbedUnimplementedCommunicationServer()
}

func RegisterCommunicationServer(s grpc.ServiceRegistrar, srv CommunicationServer) {
	s.RegisterService(&Communication_ServiceDesc, srv)
}

func _Communication_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auction.Communication/result",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).Result(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Communication_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auction.Communication/bid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).Bid(ctx, req.(*BidMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Communication_StartAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).StartAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auction.Communication/startAuction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).StartAuction(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Communication_EndAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).EndAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auction.Communication/endAuction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).EndAuction(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Communication_GetReplicas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).GetReplicas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auction.Communication/getReplicas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).GetReplicas(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

// Communication_ServiceDesc is the grpc.ServiceDesc for Communication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Communication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auction.Communication",
	HandlerType: (*CommunicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "result",
			Handler:    _Communication_Result_Handler,
		},
		{
			MethodName: "bid",
			Handler:    _Communication_Bid_Handler,
		},
		{
			MethodName: "startAuction",
			Handler:    _Communication_StartAuction_Handler,
		},
		{
			MethodName: "endAuction",
			Handler:    _Communication_EndAuction_Handler,
		},
		{
			MethodName: "getReplicas",
			Handler:    _Communication_GetReplicas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auction/auction.proto",
}