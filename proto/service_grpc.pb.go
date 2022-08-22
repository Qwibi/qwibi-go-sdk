// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: qwibi/service.proto

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
	// AUTH return authToken
	Auth(ctx context.Context, in *QPBxAuthRequest, opts ...grpc.CallOption) (*QPBxAuthResponse, error)
	// Join to geo object layer
	Join(ctx context.Context, in *QPBxJoinRequest, opts ...grpc.CallOption) (*QPBxJoinResponse, error)
	// Post
	Post(ctx context.Context, in *QPBxPostRequest, opts ...grpc.CallOption) (*QPBxPostResponse, error)
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

func (c *qPBxApiClient) Join(ctx context.Context, in *QPBxJoinRequest, opts ...grpc.CallOption) (*QPBxJoinResponse, error) {
	out := new(QPBxJoinResponse)
	err := c.cc.Invoke(ctx, "/QPBxApi/Join", in, out, opts...)
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

func (c *qPBxApiClient) Stream(ctx context.Context, opts ...grpc.CallOption) (QPBxApi_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &QPBxApi_ServiceDesc.Streams[0], "/QPBxApi/Stream", opts...)
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
	// AUTH return authToken
	Auth(context.Context, *QPBxAuthRequest) (*QPBxAuthResponse, error)
	// Join to geo object layer
	Join(context.Context, *QPBxJoinRequest) (*QPBxJoinResponse, error)
	// Post
	Post(context.Context, *QPBxPostRequest) (*QPBxPostResponse, error)
	// Stream
	Stream(QPBxApi_StreamServer) error
}

// UnimplementedQPBxApiServer should be embedded to have forward compatible implementations.
type UnimplementedQPBxApiServer struct {
}

func (UnimplementedQPBxApiServer) Auth(context.Context, *QPBxAuthRequest) (*QPBxAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedQPBxApiServer) Join(context.Context, *QPBxJoinRequest) (*QPBxJoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedQPBxApiServer) Post(context.Context, *QPBxPostRequest) (*QPBxPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
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

func _QPBxApi_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QPBxJoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QPBxApiServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QPBxApi/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QPBxApiServer).Join(ctx, req.(*QPBxJoinRequest))
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
			MethodName: "Join",
			Handler:    _QPBxApi_Join_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _QPBxApi_Post_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _QPBxApi_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "qwibi/service.proto",
}
