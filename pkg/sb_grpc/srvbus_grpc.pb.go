// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sb_grpc

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

// MessengerClient is the client API for Messenger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessengerClient interface {
	GetServer(ctx context.Context, in *MServerReq, opts ...grpc.CallOption) (*MServer, error)
	SendMessage(ctx context.Context, in *SendMsgRequest, opts ...grpc.CallOption) (*SendMsgResponse, error)
	GetMessages(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (Messenger_GetMessagesClient, error)
	HasQueue(ctx context.Context, in *QueueRequest, opts ...grpc.CallOption) (*QueueResponse, error)
}

type messengerClient struct {
	cc grpc.ClientConnInterface
}

func NewMessengerClient(cc grpc.ClientConnInterface) MessengerClient {
	return &messengerClient{cc}
}

func (c *messengerClient) GetServer(ctx context.Context, in *MServerReq, opts ...grpc.CallOption) (*MServer, error) {
	out := new(MServer)
	err := c.cc.Invoke(ctx, "/sb_grpc.Messenger/GetServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerClient) SendMessage(ctx context.Context, in *SendMsgRequest, opts ...grpc.CallOption) (*SendMsgResponse, error) {
	out := new(SendMsgResponse)
	err := c.cc.Invoke(ctx, "/sb_grpc.Messenger/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerClient) GetMessages(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (Messenger_GetMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Messenger_ServiceDesc.Streams[0], "/sb_grpc.Messenger/GetMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &messengerGetMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Messenger_GetMessagesClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type messengerGetMessagesClient struct {
	grpc.ClientStream
}

func (x *messengerGetMessagesClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messengerClient) HasQueue(ctx context.Context, in *QueueRequest, opts ...grpc.CallOption) (*QueueResponse, error) {
	out := new(QueueResponse)
	err := c.cc.Invoke(ctx, "/sb_grpc.Messenger/HasQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessengerServer is the server API for Messenger service.
// All implementations must embed UnimplementedMessengerServer
// for forward compatibility
type MessengerServer interface {
	GetServer(context.Context, *MServerReq) (*MServer, error)
	SendMessage(context.Context, *SendMsgRequest) (*SendMsgResponse, error)
	GetMessages(*MessageRequest, Messenger_GetMessagesServer) error
	HasQueue(context.Context, *QueueRequest) (*QueueResponse, error)
	mustEmbedUnimplementedMessengerServer()
}

// UnimplementedMessengerServer must be embedded to have forward compatible implementations.
type UnimplementedMessengerServer struct {
}

func (UnimplementedMessengerServer) GetServer(context.Context, *MServerReq) (*MServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServer not implemented")
}
func (UnimplementedMessengerServer) SendMessage(context.Context, *SendMsgRequest) (*SendMsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMessengerServer) GetMessages(*MessageRequest, Messenger_GetMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedMessengerServer) HasQueue(context.Context, *QueueRequest) (*QueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasQueue not implemented")
}
func (UnimplementedMessengerServer) mustEmbedUnimplementedMessengerServer() {}

// UnsafeMessengerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessengerServer will
// result in compilation errors.
type UnsafeMessengerServer interface {
	mustEmbedUnimplementedMessengerServer()
}

func RegisterMessengerServer(s grpc.ServiceRegistrar, srv MessengerServer) {
	s.RegisterService(&Messenger_ServiceDesc, srv)
}

func _Messenger_GetServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MServerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServer).GetServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sb_grpc.Messenger/GetServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServer).GetServer(ctx, req.(*MServerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messenger_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sb_grpc.Messenger/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServer).SendMessage(ctx, req.(*SendMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messenger_GetMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessengerServer).GetMessages(m, &messengerGetMessagesServer{stream})
}

type Messenger_GetMessagesServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type messengerGetMessagesServer struct {
	grpc.ServerStream
}

func (x *messengerGetMessagesServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Messenger_HasQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServer).HasQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sb_grpc.Messenger/HasQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServer).HasQueue(ctx, req.(*QueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Messenger_ServiceDesc is the grpc.ServiceDesc for Messenger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messenger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sb_grpc.Messenger",
	HandlerType: (*MessengerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServer",
			Handler:    _Messenger_GetServer_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Messenger_SendMessage_Handler,
		},
		{
			MethodName: "HasQueue",
			Handler:    _Messenger_HasQueue_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMessages",
			Handler:       _Messenger_GetMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/srvbus.proto",
}
