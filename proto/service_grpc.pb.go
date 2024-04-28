// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: service.proto

package proto

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

// QPBxApiClient is the client API for QPBxApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QPBxApiClient interface {
	// Auth
	Auth(ctx context.Context, in *QPBxAuthRequest, opts ...grpc.CallOption) (*QPBxAuthResponse, error)
	// Account
	Account(ctx context.Context, in *QPBxAccountRequest, opts ...grpc.CallOption) (*QPBxAccountResponse, error)
	// // Layer
	Layer(ctx context.Context, in *QPBxLayerRequest, opts ...grpc.CallOption) (*QPBxLayerResponse, error)
	// Object
	Post(ctx context.Context, in *QPBxPostRequest, opts ...grpc.CallOption) (*QPBxPostResponse, error)
	Get(ctx context.Context, in *QPBxGetRequest, opts ...grpc.CallOption) (*QPBxGetResponse, error)
	// Token
	Token(ctx context.Context, in *QPBxTokenRequest, opts ...grpc.CallOption) (*QPBxTokenResponse, error)
	// Command
	Command(ctx context.Context, in *QPBxCommandRequest, opts ...grpc.CallOption) (*QPBxCommandResponse, error)
	// Bot
	Bot(ctx context.Context, opts ...grpc.CallOption) (QPBxApi_BotClient, error)
	// Stream
	Stream(ctx context.Context, opts ...grpc.CallOption) (QPBxApi_StreamClient, error)
}

type qPBxApiClient struct {
	cc grpc.ClientConnInterface
}

func NewQPBxApiClient(cc grpc.ClientConnInterface) QPBxApiClient {
	return &qPBxApiClient{cc}
}

func (c *qPBxApiClient) Auth(ctx context.Context, in *QPBxAuthRequest, opts ...grpc.CallOption) (*QPBxAuthResponse, error) {
	out := new(QPBxAuthResponse)
	err := c.cc.Invoke(ctx, "/QPBxApi/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qPBxApiClient) Account(ctx context.Context, in *QPBxAccountRequest, opts ...grpc.CallOption) (*QPBxAccountResponse, error) {
	out := new(QPBxAccountResponse)
	err := c.cc.Invoke(ctx, "/QPBxApi/Account", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qPBxApiClient) Layer(ctx context.Context, in *QPBxLayerRequest, opts ...grpc.CallOption) (*QPBxLayerResponse, error) {
	out := new(QPBxLayerResponse)
	err := c.cc.Invoke(ctx, "/QPBxApi/Layer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qPBxApiClient) Post(ctx context.Context, in *QPBxPostRequest, opts ...grpc.CallOption) (*QPBxPostResponse, error) {
	out := new(QPBxPostResponse)
	err := c.cc.Invoke(ctx, "/QPBxApi/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qPBxApiClient) Get(ctx context.Context, in *QPBxGetRequest, opts ...grpc.CallOption) (*QPBxGetResponse, error) {
	out := new(QPBxGetResponse)
	err := c.cc.Invoke(ctx, "/QPBxApi/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qPBxApiClient) Token(ctx context.Context, in *QPBxTokenRequest, opts ...grpc.CallOption) (*QPBxTokenResponse, error) {
	out := new(QPBxTokenResponse)
	err := c.cc.Invoke(ctx, "/QPBxApi/Token", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qPBxApiClient) Command(ctx context.Context, in *QPBxCommandRequest, opts ...grpc.CallOption) (*QPBxCommandResponse, error) {
	out := new(QPBxCommandResponse)
	err := c.cc.Invoke(ctx, "/QPBxApi/Command", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qPBxApiClient) Bot(ctx context.Context, opts ...grpc.CallOption) (QPBxApi_BotClient, error) {
	stream, err := c.cc.NewStream(ctx, &QPBxApi_ServiceDesc.Streams[0], "/QPBxApi/Bot", opts...)
	if err != nil {
		return nil, err
	}
	x := &qPBxApiBotClient{stream}
	return x, nil
}

type QPBxApi_BotClient interface {
	Send(*QPBxCommandResponse) error
	Recv() (*QPBxCommandRequest, error)
	grpc.ClientStream
}

type qPBxApiBotClient struct {
	grpc.ClientStream
}

func (x *qPBxApiBotClient) Send(m *QPBxCommandResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *qPBxApiBotClient) Recv() (*QPBxCommandRequest, error) {
	m := new(QPBxCommandRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *qPBxApiClient) Stream(ctx context.Context, opts ...grpc.CallOption) (QPBxApi_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &QPBxApi_ServiceDesc.Streams[1], "/QPBxApi/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &qPBxApiStreamClient{stream}
	return x, nil
}

type QPBxApi_StreamClient interface {
	Send(*QPBxStreamRequest) error
	Recv() (*QPBxStreamResponse, error)
	grpc.ClientStream
}

type qPBxApiStreamClient struct {
	grpc.ClientStream
}

func (x *qPBxApiStreamClient) Send(m *QPBxStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *qPBxApiStreamClient) Recv() (*QPBxStreamResponse, error) {
	m := new(QPBxStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QPBxApiServer is the server API for QPBxApi service.
// All implementations should embed UnimplementedQPBxApiServer
// for forward compatibility
type QPBxApiServer interface {
	// Auth
	Auth(context.Context, *QPBxAuthRequest) (*QPBxAuthResponse, error)
	// Account
	Account(context.Context, *QPBxAccountRequest) (*QPBxAccountResponse, error)
	// // Layer
	Layer(context.Context, *QPBxLayerRequest) (*QPBxLayerResponse, error)
	// Object
	Post(context.Context, *QPBxPostRequest) (*QPBxPostResponse, error)
	Get(context.Context, *QPBxGetRequest) (*QPBxGetResponse, error)
	// Token
	Token(context.Context, *QPBxTokenRequest) (*QPBxTokenResponse, error)
	// Command
	Command(context.Context, *QPBxCommandRequest) (*QPBxCommandResponse, error)
	// Bot
	Bot(QPBxApi_BotServer) error
	// Stream
	Stream(QPBxApi_StreamServer) error
}

// UnimplementedQPBxApiServer should be embedded to have forward compatible implementations.
type UnimplementedQPBxApiServer struct {
}

func (UnimplementedQPBxApiServer) Auth(context.Context, *QPBxAuthRequest) (*QPBxAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedQPBxApiServer) Account(context.Context, *QPBxAccountRequest) (*QPBxAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Account not implemented")
}
func (UnimplementedQPBxApiServer) Layer(context.Context, *QPBxLayerRequest) (*QPBxLayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Layer not implemented")
}
func (UnimplementedQPBxApiServer) Post(context.Context, *QPBxPostRequest) (*QPBxPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedQPBxApiServer) Get(context.Context, *QPBxGetRequest) (*QPBxGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedQPBxApiServer) Token(context.Context, *QPBxTokenRequest) (*QPBxTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Token not implemented")
}
func (UnimplementedQPBxApiServer) Command(context.Context, *QPBxCommandRequest) (*QPBxCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Command not implemented")
}
func (UnimplementedQPBxApiServer) Bot(QPBxApi_BotServer) error {
	return status.Errorf(codes.Unimplemented, "method Bot not implemented")
}
func (UnimplementedQPBxApiServer) Stream(QPBxApi_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}

// UnsafeQPBxApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QPBxApiServer will
// result in compilation errors.
type UnsafeQPBxApiServer interface {
	mustEmbedUnimplementedQPBxApiServer()
}

func RegisterQPBxApiServer(s grpc.ServiceRegistrar, srv QPBxApiServer) {
	s.RegisterService(&QPBxApi_ServiceDesc, srv)
}

func _QPBxApi_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QPBxAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QPBxApiServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QPBxApi/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QPBxApiServer).Auth(ctx, req.(*QPBxAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QPBxApi_Account_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QPBxAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QPBxApiServer).Account(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QPBxApi/Account",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QPBxApiServer).Account(ctx, req.(*QPBxAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QPBxApi_Layer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QPBxLayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QPBxApiServer).Layer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QPBxApi/Layer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QPBxApiServer).Layer(ctx, req.(*QPBxLayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QPBxApi_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QPBxPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QPBxApiServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QPBxApi/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QPBxApiServer).Post(ctx, req.(*QPBxPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QPBxApi_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QPBxGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QPBxApiServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QPBxApi/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QPBxApiServer).Get(ctx, req.(*QPBxGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QPBxApi_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QPBxTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QPBxApiServer).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QPBxApi/Token",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QPBxApiServer).Token(ctx, req.(*QPBxTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QPBxApi_Command_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QPBxCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QPBxApiServer).Command(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QPBxApi/Command",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QPBxApiServer).Command(ctx, req.(*QPBxCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QPBxApi_Bot_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(QPBxApiServer).Bot(&qPBxApiBotServer{stream})
}

type QPBxApi_BotServer interface {
	Send(*QPBxCommandRequest) error
	Recv() (*QPBxCommandResponse, error)
	grpc.ServerStream
}

type qPBxApiBotServer struct {
	grpc.ServerStream
}

func (x *qPBxApiBotServer) Send(m *QPBxCommandRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *qPBxApiBotServer) Recv() (*QPBxCommandResponse, error) {
	m := new(QPBxCommandResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _QPBxApi_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(QPBxApiServer).Stream(&qPBxApiStreamServer{stream})
}

type QPBxApi_StreamServer interface {
	Send(*QPBxStreamResponse) error
	Recv() (*QPBxStreamRequest, error)
	grpc.ServerStream
}

type qPBxApiStreamServer struct {
	grpc.ServerStream
}

func (x *qPBxApiStreamServer) Send(m *QPBxStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *qPBxApiStreamServer) Recv() (*QPBxStreamRequest, error) {
	m := new(QPBxStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QPBxApi_ServiceDesc is the grpc.ServiceDesc for QPBxApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QPBxApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "QPBxApi",
	HandlerType: (*QPBxApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _QPBxApi_Auth_Handler,
		},
		{
			MethodName: "Account",
			Handler:    _QPBxApi_Account_Handler,
		},
		{
			MethodName: "Layer",
			Handler:    _QPBxApi_Layer_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _QPBxApi_Post_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _QPBxApi_Get_Handler,
		},
		{
			MethodName: "Token",
			Handler:    _QPBxApi_Token_Handler,
		},
		{
			MethodName: "Command",
			Handler:    _QPBxApi_Command_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Bot",
			Handler:       _QPBxApi_Bot_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Stream",
			Handler:       _QPBxApi_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}
