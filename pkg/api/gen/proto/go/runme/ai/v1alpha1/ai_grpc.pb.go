// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: runme/ai/v1alpha1/ai.proto

package aiv1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AIService_GenerateCells_FullMethodName = "/runme.ai.v1alpha1.AIService/GenerateCells"
)

// AIServiceClient is the client API for AIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The AIService service is used to provide assistant capabilities to Runme.
type AIServiceClient interface {
	// GenerateCells uses the AI to generate cells to insert into the notebook.
	GenerateCells(ctx context.Context, in *GenerateCellsRequest, opts ...grpc.CallOption) (*GenerateCellsResponse, error)
}

type aIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAIServiceClient(cc grpc.ClientConnInterface) AIServiceClient {
	return &aIServiceClient{cc}
}

func (c *aIServiceClient) GenerateCells(ctx context.Context, in *GenerateCellsRequest, opts ...grpc.CallOption) (*GenerateCellsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateCellsResponse)
	err := c.cc.Invoke(ctx, AIService_GenerateCells_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AIServiceServer is the server API for AIService service.
// All implementations must embed UnimplementedAIServiceServer
// for forward compatibility
//
// The AIService service is used to provide assistant capabilities to Runme.
type AIServiceServer interface {
	// GenerateCells uses the AI to generate cells to insert into the notebook.
	GenerateCells(context.Context, *GenerateCellsRequest) (*GenerateCellsResponse, error)
	mustEmbedUnimplementedAIServiceServer()
}

// UnimplementedAIServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAIServiceServer struct {
}

func (UnimplementedAIServiceServer) GenerateCells(context.Context, *GenerateCellsRequest) (*GenerateCellsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateCells not implemented")
}
func (UnimplementedAIServiceServer) mustEmbedUnimplementedAIServiceServer() {}

// UnsafeAIServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AIServiceServer will
// result in compilation errors.
type UnsafeAIServiceServer interface {
	mustEmbedUnimplementedAIServiceServer()
}

func RegisterAIServiceServer(s grpc.ServiceRegistrar, srv AIServiceServer) {
	s.RegisterService(&AIService_ServiceDesc, srv)
}

func _AIService_GenerateCells_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateCellsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).GenerateCells(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIService_GenerateCells_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).GenerateCells(ctx, req.(*GenerateCellsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AIService_ServiceDesc is the grpc.ServiceDesc for AIService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AIService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "runme.ai.v1alpha1.AIService",
	HandlerType: (*AIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateCells",
			Handler:    _AIService_GenerateCells_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "runme/ai/v1alpha1/ai.proto",
}
