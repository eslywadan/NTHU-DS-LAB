// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: modules/video/pb/stream.proto

package pb

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

// VideoStreamClient is the client API for VideoStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoStreamClient interface {
	HandleVideoCreated(ctx context.Context, in *HandleVideoCreatedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type videoStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoStreamClient(cc grpc.ClientConnInterface) VideoStreamClient {
	return &videoStreamClient{cc}
}

func (c *videoStreamClient) HandleVideoCreated(ctx context.Context, in *HandleVideoCreatedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/video.pb.VideoStream/HandleVideoCreated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoStreamServer is the server API for VideoStream service.
// All implementations must embed UnimplementedVideoStreamServer
// for forward compatibility
type VideoStreamServer interface {
	HandleVideoCreated(context.Context, *HandleVideoCreatedRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedVideoStreamServer()
}

// UnimplementedVideoStreamServer must be embedded to have forward compatible implementations.
type UnimplementedVideoStreamServer struct {
}

func (UnimplementedVideoStreamServer) HandleVideoCreated(context.Context, *HandleVideoCreatedRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleVideoCreated not implemented")
}
func (UnimplementedVideoStreamServer) mustEmbedUnimplementedVideoStreamServer() {}

// UnsafeVideoStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoStreamServer will
// result in compilation errors.
type UnsafeVideoStreamServer interface {
	mustEmbedUnimplementedVideoStreamServer()
}

func RegisterVideoStreamServer(s grpc.ServiceRegistrar, srv VideoStreamServer) {
	s.RegisterService(&VideoStream_ServiceDesc, srv)
}

func _VideoStream_HandleVideoCreated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleVideoCreatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoStreamServer).HandleVideoCreated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.pb.VideoStream/HandleVideoCreated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoStreamServer).HandleVideoCreated(ctx, req.(*HandleVideoCreatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoStream_ServiceDesc is the grpc.ServiceDesc for VideoStream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "video.pb.VideoStream",
	HandlerType: (*VideoStreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleVideoCreated",
			Handler:    _VideoStream_HandleVideoCreated_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/video/pb/stream.proto",
}
