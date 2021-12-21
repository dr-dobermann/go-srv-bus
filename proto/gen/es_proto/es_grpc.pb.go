// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package es_proto

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

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventServiceClient interface {
	// checks if the topic exists on host server.
	HasTopic(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*OpResponse, error)
	// adds a new topic or a whole topics branch to the host server.
	AddTopics(ctx context.Context, in *AddTopicReq, opts ...grpc.CallOption) (*OpResponse, error)
	// Returns topic or branch from the host server
	DelTopics(ctx context.Context, in *DelTopicReq, opts ...grpc.CallOption) (*OpResponse, error)
	// adds a new event on the host server.
	AddEvent(ctx context.Context, in *EventRegistration, opts ...grpc.CallOption) (*OpResponse, error)
	// creates single or multi- subscription on the host server.
	Subscribe(ctx context.Context, in *SubscriptionRequest, opts ...grpc.CallOption) (EventService_SubscribeClient, error)
	// cancels subsciptions for one or many topics on the host server.
	UnSubscribe(ctx context.Context, in *UnsubsibeRequest, opts ...grpc.CallOption) (*OpResponse, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) HasTopic(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*OpResponse, error) {
	out := new(OpResponse)
	err := c.cc.Invoke(ctx, "/es_proto.EventService/HasTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) AddTopics(ctx context.Context, in *AddTopicReq, opts ...grpc.CallOption) (*OpResponse, error) {
	out := new(OpResponse)
	err := c.cc.Invoke(ctx, "/es_proto.EventService/AddTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) DelTopics(ctx context.Context, in *DelTopicReq, opts ...grpc.CallOption) (*OpResponse, error) {
	out := new(OpResponse)
	err := c.cc.Invoke(ctx, "/es_proto.EventService/DelTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) AddEvent(ctx context.Context, in *EventRegistration, opts ...grpc.CallOption) (*OpResponse, error) {
	out := new(OpResponse)
	err := c.cc.Invoke(ctx, "/es_proto.EventService/AddEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Subscribe(ctx context.Context, in *SubscriptionRequest, opts ...grpc.CallOption) (EventService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventService_ServiceDesc.Streams[0], "/es_proto.EventService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventService_SubscribeClient interface {
	Recv() (*EventEnvelope, error)
	grpc.ClientStream
}

type eventServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *eventServiceSubscribeClient) Recv() (*EventEnvelope, error) {
	m := new(EventEnvelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventServiceClient) UnSubscribe(ctx context.Context, in *UnsubsibeRequest, opts ...grpc.CallOption) (*OpResponse, error) {
	out := new(OpResponse)
	err := c.cc.Invoke(ctx, "/es_proto.EventService/UnSubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
// All implementations must embed UnimplementedEventServiceServer
// for forward compatibility
type EventServiceServer interface {
	// checks if the topic exists on host server.
	HasTopic(context.Context, *TopicRequest) (*OpResponse, error)
	// adds a new topic or a whole topics branch to the host server.
	AddTopics(context.Context, *AddTopicReq) (*OpResponse, error)
	// Returns topic or branch from the host server
	DelTopics(context.Context, *DelTopicReq) (*OpResponse, error)
	// adds a new event on the host server.
	AddEvent(context.Context, *EventRegistration) (*OpResponse, error)
	// creates single or multi- subscription on the host server.
	Subscribe(*SubscriptionRequest, EventService_SubscribeServer) error
	// cancels subsciptions for one or many topics on the host server.
	UnSubscribe(context.Context, *UnsubsibeRequest) (*OpResponse, error)
	mustEmbedUnimplementedEventServiceServer()
}

// UnimplementedEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (UnimplementedEventServiceServer) HasTopic(context.Context, *TopicRequest) (*OpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasTopic not implemented")
}
func (UnimplementedEventServiceServer) AddTopics(context.Context, *AddTopicReq) (*OpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTopics not implemented")
}
func (UnimplementedEventServiceServer) DelTopics(context.Context, *DelTopicReq) (*OpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelTopics not implemented")
}
func (UnimplementedEventServiceServer) AddEvent(context.Context, *EventRegistration) (*OpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEvent not implemented")
}
func (UnimplementedEventServiceServer) Subscribe(*SubscriptionRequest, EventService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedEventServiceServer) UnSubscribe(context.Context, *UnsubsibeRequest) (*OpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnSubscribe not implemented")
}
func (UnimplementedEventServiceServer) mustEmbedUnimplementedEventServiceServer() {}

// UnsafeEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServiceServer will
// result in compilation errors.
type UnsafeEventServiceServer interface {
	mustEmbedUnimplementedEventServiceServer()
}

func RegisterEventServiceServer(s grpc.ServiceRegistrar, srv EventServiceServer) {
	s.RegisterService(&EventService_ServiceDesc, srv)
}

func _EventService_HasTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).HasTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/es_proto.EventService/HasTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).HasTopic(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_AddTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTopicReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).AddTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/es_proto.EventService/AddTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).AddTopics(ctx, req.(*AddTopicReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_DelTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelTopicReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).DelTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/es_proto.EventService/DelTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).DelTopics(ctx, req.(*DelTopicReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRegistration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/es_proto.EventService/AddEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).AddEvent(ctx, req.(*EventRegistration))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscriptionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventServiceServer).Subscribe(m, &eventServiceSubscribeServer{stream})
}

type EventService_SubscribeServer interface {
	Send(*EventEnvelope) error
	grpc.ServerStream
}

type eventServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *eventServiceSubscribeServer) Send(m *EventEnvelope) error {
	return x.ServerStream.SendMsg(m)
}

func _EventService_UnSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubsibeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).UnSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/es_proto.EventService/UnSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).UnSubscribe(ctx, req.(*UnsubsibeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventService_ServiceDesc is the grpc.ServiceDesc for EventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "es_proto.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HasTopic",
			Handler:    _EventService_HasTopic_Handler,
		},
		{
			MethodName: "AddTopics",
			Handler:    _EventService_AddTopics_Handler,
		},
		{
			MethodName: "DelTopics",
			Handler:    _EventService_DelTopics_Handler,
		},
		{
			MethodName: "AddEvent",
			Handler:    _EventService_AddEvent_Handler,
		},
		{
			MethodName: "UnSubscribe",
			Handler:    _EventService_UnSubscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _EventService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/es.proto",
}
