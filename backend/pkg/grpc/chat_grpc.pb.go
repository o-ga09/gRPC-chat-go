// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: chat.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessagingServiseClient is the client API for MessagingServise service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagingServiseClient interface {
	// サービスが持つメソッドの定義
	HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Status, error)
	SendMessage(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgResponse, error)
	ReceiveMessage(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (MessagingServise_ReceiveMessageClient, error)
}

type messagingServiseClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagingServiseClient(cc grpc.ClientConnInterface) MessagingServiseClient {
	return &messagingServiseClient{cc}
}

func (c *messagingServiseClient) HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ChatRPC.MessagingServise/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiseClient) SendMessage(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgResponse, error) {
	out := new(MsgResponse)
	err := c.cc.Invoke(ctx, "/ChatRPC.MessagingServise/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiseClient) ReceiveMessage(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (MessagingServise_ReceiveMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessagingServise_ServiceDesc.Streams[0], "/ChatRPC.MessagingServise/ReceiveMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &messagingServiseReceiveMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessagingServise_ReceiveMessageClient interface {
	Recv() (*MsgResponse, error)
	grpc.ClientStream
}

type messagingServiseReceiveMessageClient struct {
	grpc.ClientStream
}

func (x *messagingServiseReceiveMessageClient) Recv() (*MsgResponse, error) {
	m := new(MsgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessagingServiseServer is the server API for MessagingServise service.
// All implementations must embed UnimplementedMessagingServiseServer
// for forward compatibility
type MessagingServiseServer interface {
	// サービスが持つメソッドの定義
	HealthCheck(context.Context, *emptypb.Empty) (*Status, error)
	SendMessage(context.Context, *MsgRequest) (*MsgResponse, error)
	ReceiveMessage(*MsgRequest, MessagingServise_ReceiveMessageServer) error
	mustEmbedUnimplementedMessagingServiseServer()
}

// UnimplementedMessagingServiseServer must be embedded to have forward compatible implementations.
type UnimplementedMessagingServiseServer struct {
}

func (UnimplementedMessagingServiseServer) HealthCheck(context.Context, *emptypb.Empty) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedMessagingServiseServer) SendMessage(context.Context, *MsgRequest) (*MsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMessagingServiseServer) ReceiveMessage(*MsgRequest, MessagingServise_ReceiveMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveMessage not implemented")
}
func (UnimplementedMessagingServiseServer) mustEmbedUnimplementedMessagingServiseServer() {}

// UnsafeMessagingServiseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagingServiseServer will
// result in compilation errors.
type UnsafeMessagingServiseServer interface {
	mustEmbedUnimplementedMessagingServiseServer()
}

func RegisterMessagingServiseServer(s grpc.ServiceRegistrar, srv MessagingServiseServer) {
	s.RegisterService(&MessagingServise_ServiceDesc, srv)
}

func _MessagingServise_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiseServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatRPC.MessagingServise/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiseServer).HealthCheck(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingServise_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiseServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatRPC.MessagingServise/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiseServer).SendMessage(ctx, req.(*MsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingServise_ReceiveMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MsgRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessagingServiseServer).ReceiveMessage(m, &messagingServiseReceiveMessageServer{stream})
}

type MessagingServise_ReceiveMessageServer interface {
	Send(*MsgResponse) error
	grpc.ServerStream
}

type messagingServiseReceiveMessageServer struct {
	grpc.ServerStream
}

func (x *messagingServiseReceiveMessageServer) Send(m *MsgResponse) error {
	return x.ServerStream.SendMsg(m)
}

// MessagingServise_ServiceDesc is the grpc.ServiceDesc for MessagingServise service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessagingServise_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChatRPC.MessagingServise",
	HandlerType: (*MessagingServiseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _MessagingServise_HealthCheck_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _MessagingServise_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveMessage",
			Handler:       _MessagingServise_ReceiveMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat.proto",
}