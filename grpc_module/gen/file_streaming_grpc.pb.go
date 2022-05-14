// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.2
// source: file_streaming.proto

package gen

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

// FileStreamingServiceClient is the client API for FileStreamingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileStreamingServiceClient interface {
	ServerStreaming(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (FileStreamingService_ServerStreamingClient, error)
}

type fileStreamingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileStreamingServiceClient(cc grpc.ClientConnInterface) FileStreamingServiceClient {
	return &fileStreamingServiceClient{cc}
}

func (c *fileStreamingServiceClient) ServerStreaming(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (FileStreamingService_ServerStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileStreamingService_ServiceDesc.Streams[0], "/file_streaming.FileStreamingService/ServerStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileStreamingServiceServerStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileStreamingService_ServerStreamingClient interface {
	Recv() (*FileResponse, error)
	grpc.ClientStream
}

type fileStreamingServiceServerStreamingClient struct {
	grpc.ClientStream
}

func (x *fileStreamingServiceServerStreamingClient) Recv() (*FileResponse, error) {
	m := new(FileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileStreamingServiceServer is the server API for FileStreamingService service.
// All implementations must embed UnimplementedFileStreamingServiceServer
// for forward compatibility
type FileStreamingServiceServer interface {
	ServerStreaming(*FileRequest, FileStreamingService_ServerStreamingServer) error
	mustEmbedUnimplementedFileStreamingServiceServer()
}

// UnimplementedFileStreamingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileStreamingServiceServer struct {
}

func (UnimplementedFileStreamingServiceServer) ServerStreaming(*FileRequest, FileStreamingService_ServerStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreaming not implemented")
}
func (UnimplementedFileStreamingServiceServer) mustEmbedUnimplementedFileStreamingServiceServer() {}

// UnsafeFileStreamingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileStreamingServiceServer will
// result in compilation errors.
type UnsafeFileStreamingServiceServer interface {
	mustEmbedUnimplementedFileStreamingServiceServer()
}

func RegisterFileStreamingServiceServer(s grpc.ServiceRegistrar, srv FileStreamingServiceServer) {
	s.RegisterService(&FileStreamingService_ServiceDesc, srv)
}

func _FileStreamingService_ServerStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileStreamingServiceServer).ServerStreaming(m, &fileStreamingServiceServerStreamingServer{stream})
}

type FileStreamingService_ServerStreamingServer interface {
	Send(*FileResponse) error
	grpc.ServerStream
}

type fileStreamingServiceServerStreamingServer struct {
	grpc.ServerStream
}

func (x *fileStreamingServiceServerStreamingServer) Send(m *FileResponse) error {
	return x.ServerStream.SendMsg(m)
}

// FileStreamingService_ServiceDesc is the grpc.ServiceDesc for FileStreamingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileStreamingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file_streaming.FileStreamingService",
	HandlerType: (*FileStreamingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreaming",
			Handler:       _FileStreamingService_ServerStreaming_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "file_streaming.proto",
}