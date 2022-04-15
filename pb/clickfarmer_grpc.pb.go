// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// ClickFarmerClient is the client API for ClickFarmer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClickFarmerClient interface {
	GetClicks(ctx context.Context, in *GetClicksRequest, opts ...grpc.CallOption) (*GetClicksResponse, error)
	SetClicks(ctx context.Context, in *SetClicksRequest, opts ...grpc.CallOption) (*SetClicksResponse, error)
}

type clickFarmerClient struct {
	cc grpc.ClientConnInterface
}

func NewClickFarmerClient(cc grpc.ClientConnInterface) ClickFarmerClient {
	return &clickFarmerClient{cc}
}

func (c *clickFarmerClient) GetClicks(ctx context.Context, in *GetClicksRequest, opts ...grpc.CallOption) (*GetClicksResponse, error) {
	out := new(GetClicksResponse)
	err := c.cc.Invoke(ctx, "/clickfarmer.ClickFarmer/GetClicks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clickFarmerClient) SetClicks(ctx context.Context, in *SetClicksRequest, opts ...grpc.CallOption) (*SetClicksResponse, error) {
	out := new(SetClicksResponse)
	err := c.cc.Invoke(ctx, "/clickfarmer.ClickFarmer/SetClicks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClickFarmerServer is the server API for ClickFarmer service.
// All implementations must embed UnimplementedClickFarmerServer
// for forward compatibility
type ClickFarmerServer interface {
	GetClicks(context.Context, *GetClicksRequest) (*GetClicksResponse, error)
	SetClicks(context.Context, *SetClicksRequest) (*SetClicksResponse, error)
	mustEmbedUnimplementedClickFarmerServer()
}

// UnimplementedClickFarmerServer must be embedded to have forward compatible implementations.
type UnimplementedClickFarmerServer struct {
}

func (UnimplementedClickFarmerServer) GetClicks(context.Context, *GetClicksRequest) (*GetClicksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClicks not implemented")
}
func (UnimplementedClickFarmerServer) SetClicks(context.Context, *SetClicksRequest) (*SetClicksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetClicks not implemented")
}
func (UnimplementedClickFarmerServer) mustEmbedUnimplementedClickFarmerServer() {}

// UnsafeClickFarmerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClickFarmerServer will
// result in compilation errors.
type UnsafeClickFarmerServer interface {
	mustEmbedUnimplementedClickFarmerServer()
}

func RegisterClickFarmerServer(s grpc.ServiceRegistrar, srv ClickFarmerServer) {
	s.RegisterService(&ClickFarmer_ServiceDesc, srv)
}

func _ClickFarmer_GetClicks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClicksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClickFarmerServer).GetClicks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clickfarmer.ClickFarmer/GetClicks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClickFarmerServer).GetClicks(ctx, req.(*GetClicksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClickFarmer_SetClicks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetClicksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClickFarmerServer).SetClicks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clickfarmer.ClickFarmer/SetClicks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClickFarmerServer).SetClicks(ctx, req.(*SetClicksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClickFarmer_ServiceDesc is the grpc.ServiceDesc for ClickFarmer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClickFarmer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clickfarmer.ClickFarmer",
	HandlerType: (*ClickFarmerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClicks",
			Handler:    _ClickFarmer_GetClicks_Handler,
		},
		{
			MethodName: "SetClicks",
			Handler:    _ClickFarmer_SetClicks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clickfarmer.proto",
}