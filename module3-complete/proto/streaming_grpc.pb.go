// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: proto/streaming.proto

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

const (
	StreamingService_StreamServerTime_FullMethodName = "/streaming.StreamingService/StreamServerTime"
	StreamingService_LogStream_FullMethodName        = "/streaming.StreamingService/LogStream"
	StreamingService_Echo_FullMethodName             = "/streaming.StreamingService/Echo"
)

// StreamingServiceClient is the client API for StreamingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamingServiceClient interface {
	// StreamServerTime is an example of a server-streaming RPC.
	// It will stream the current server time back to the client at specified intervals.
	StreamServerTime(ctx context.Context, in *StreamServerTimeRequest, opts ...grpc.CallOption) (StreamingService_StreamServerTimeClient, error)
	// LogStream is an example of a client-streaming RPC.
	// It allows a client to stream log entries to a server.
	LogStream(ctx context.Context, opts ...grpc.CallOption) (StreamingService_LogStreamClient, error)
	// Echo is an example of a bidirectional-streaming RPC.
	// Where the client can send a stream of messages, which the server will echo back
	// the received messages in a stream.
	Echo(ctx context.Context, opts ...grpc.CallOption) (StreamingService_EchoClient, error)
}

type streamingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingServiceClient(cc grpc.ClientConnInterface) StreamingServiceClient {
	return &streamingServiceClient{cc}
}

func (c *streamingServiceClient) StreamServerTime(ctx context.Context, in *StreamServerTimeRequest, opts ...grpc.CallOption) (StreamingService_StreamServerTimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamingService_ServiceDesc.Streams[0], StreamingService_StreamServerTime_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingServiceStreamServerTimeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamingService_StreamServerTimeClient interface {
	Recv() (*StreamServerTimeResponse, error)
	grpc.ClientStream
}

type streamingServiceStreamServerTimeClient struct {
	grpc.ClientStream
}

func (x *streamingServiceStreamServerTimeClient) Recv() (*StreamServerTimeResponse, error) {
	m := new(StreamServerTimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamingServiceClient) LogStream(ctx context.Context, opts ...grpc.CallOption) (StreamingService_LogStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamingService_ServiceDesc.Streams[1], StreamingService_LogStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingServiceLogStreamClient{stream}
	return x, nil
}

type StreamingService_LogStreamClient interface {
	Send(*LogStreamRequest) error
	CloseAndRecv() (*LogStreamResponse, error)
	grpc.ClientStream
}

type streamingServiceLogStreamClient struct {
	grpc.ClientStream
}

func (x *streamingServiceLogStreamClient) Send(m *LogStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamingServiceLogStreamClient) CloseAndRecv() (*LogStreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LogStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamingServiceClient) Echo(ctx context.Context, opts ...grpc.CallOption) (StreamingService_EchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamingService_ServiceDesc.Streams[2], StreamingService_Echo_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingServiceEchoClient{stream}
	return x, nil
}

type StreamingService_EchoClient interface {
	Send(*EchoRequest) error
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type streamingServiceEchoClient struct {
	grpc.ClientStream
}

func (x *streamingServiceEchoClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamingServiceEchoClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingServiceServer is the server API for StreamingService service.
// All implementations must embed UnimplementedStreamingServiceServer
// for forward compatibility
type StreamingServiceServer interface {
	// StreamServerTime is an example of a server-streaming RPC.
	// It will stream the current server time back to the client at specified intervals.
	StreamServerTime(*StreamServerTimeRequest, StreamingService_StreamServerTimeServer) error
	// LogStream is an example of a client-streaming RPC.
	// It allows a client to stream log entries to a server.
	LogStream(StreamingService_LogStreamServer) error
	// Echo is an example of a bidirectional-streaming RPC.
	// Where the client can send a stream of messages, which the server will echo back
	// the received messages in a stream.
	Echo(StreamingService_EchoServer) error
	mustEmbedUnimplementedStreamingServiceServer()
}

// UnimplementedStreamingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStreamingServiceServer struct {
}

func (UnimplementedStreamingServiceServer) StreamServerTime(*StreamServerTimeRequest, StreamingService_StreamServerTimeServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamServerTime not implemented")
}
func (UnimplementedStreamingServiceServer) LogStream(StreamingService_LogStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method LogStream not implemented")
}
func (UnimplementedStreamingServiceServer) Echo(StreamingService_EchoServer) error {
	return status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedStreamingServiceServer) mustEmbedUnimplementedStreamingServiceServer() {}

// UnsafeStreamingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamingServiceServer will
// result in compilation errors.
type UnsafeStreamingServiceServer interface {
	mustEmbedUnimplementedStreamingServiceServer()
}

func RegisterStreamingServiceServer(s grpc.ServiceRegistrar, srv StreamingServiceServer) {
	s.RegisterService(&StreamingService_ServiceDesc, srv)
}

func _StreamingService_StreamServerTime_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamServerTimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamingServiceServer).StreamServerTime(m, &streamingServiceStreamServerTimeServer{stream})
}

type StreamingService_StreamServerTimeServer interface {
	Send(*StreamServerTimeResponse) error
	grpc.ServerStream
}

type streamingServiceStreamServerTimeServer struct {
	grpc.ServerStream
}

func (x *streamingServiceStreamServerTimeServer) Send(m *StreamServerTimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamingService_LogStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingServiceServer).LogStream(&streamingServiceLogStreamServer{stream})
}

type StreamingService_LogStreamServer interface {
	SendAndClose(*LogStreamResponse) error
	Recv() (*LogStreamRequest, error)
	grpc.ServerStream
}

type streamingServiceLogStreamServer struct {
	grpc.ServerStream
}

func (x *streamingServiceLogStreamServer) SendAndClose(m *LogStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamingServiceLogStreamServer) Recv() (*LogStreamRequest, error) {
	m := new(LogStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamingService_Echo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingServiceServer).Echo(&streamingServiceEchoServer{stream})
}

type StreamingService_EchoServer interface {
	Send(*EchoResponse) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type streamingServiceEchoServer struct {
	grpc.ServerStream
}

func (x *streamingServiceEchoServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamingServiceEchoServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingService_ServiceDesc is the grpc.ServiceDesc for StreamingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "streaming.StreamingService",
	HandlerType: (*StreamingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamServerTime",
			Handler:       _StreamingService_StreamServerTime_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LogStream",
			Handler:       _StreamingService_LogStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Echo",
			Handler:       _StreamingService_Echo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/streaming.proto",
}
