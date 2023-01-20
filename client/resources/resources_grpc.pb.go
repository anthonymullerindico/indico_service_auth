// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protos/resources.proto

package resource

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ResourceServiceClient is the client API for ResourceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceServiceClient interface {
	CreateResource(ctx context.Context, in *CreateResourceRequest, opts ...grpc.CallOption) (*Resource, error)
	CreateResourceScope(ctx context.Context, in *CreateResourceScopeRequest, opts ...grpc.CallOption) (*CreateResourceScopeRequest, error)
	ListResources(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListResource, error)
	ListClientScopes(ctx context.Context, in *ListClientScopesRequest, opts ...grpc.CallOption) (*ListResourceScope, error)
	GetResourceScope(ctx context.Context, in *GetResourceScopeRequest, opts ...grpc.CallOption) (*ListResourceScope, error)
}

type resourceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceServiceClient(cc grpc.ClientConnInterface) ResourceServiceClient {
	return &resourceServiceClient{cc}
}

func (c *resourceServiceClient) CreateResource(ctx context.Context, in *CreateResourceRequest, opts ...grpc.CallOption) (*Resource, error) {
	out := new(Resource)
	err := c.cc.Invoke(ctx, "/ResourceService/CreateResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) CreateResourceScope(ctx context.Context, in *CreateResourceScopeRequest, opts ...grpc.CallOption) (*CreateResourceScopeRequest, error) {
	out := new(CreateResourceScopeRequest)
	err := c.cc.Invoke(ctx, "/ResourceService/CreateResourceScope", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) ListResources(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListResource, error) {
	out := new(ListResource)
	err := c.cc.Invoke(ctx, "/ResourceService/ListResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) ListClientScopes(ctx context.Context, in *ListClientScopesRequest, opts ...grpc.CallOption) (*ListResourceScope, error) {
	out := new(ListResourceScope)
	err := c.cc.Invoke(ctx, "/ResourceService/ListClientScopes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) GetResourceScope(ctx context.Context, in *GetResourceScopeRequest, opts ...grpc.CallOption) (*ListResourceScope, error) {
	out := new(ListResourceScope)
	err := c.cc.Invoke(ctx, "/ResourceService/GetResourceScope", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceServiceServer is the server API for ResourceService service.
// All implementations must embed UnimplementedResourceServiceServer
// for forward compatibility
type ResourceServiceServer interface {
	CreateResource(context.Context, *CreateResourceRequest) (*Resource, error)
	CreateResourceScope(context.Context, *CreateResourceScopeRequest) (*CreateResourceScopeRequest, error)
	ListResources(context.Context, *empty.Empty) (*ListResource, error)
	ListClientScopes(context.Context, *ListClientScopesRequest) (*ListResourceScope, error)
	GetResourceScope(context.Context, *GetResourceScopeRequest) (*ListResourceScope, error)
	mustEmbedUnimplementedResourceServiceServer()
}

// UnimplementedResourceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResourceServiceServer struct {
}

func (UnimplementedResourceServiceServer) CreateResource(context.Context, *CreateResourceRequest) (*Resource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResource not implemented")
}
func (UnimplementedResourceServiceServer) CreateResourceScope(context.Context, *CreateResourceScopeRequest) (*CreateResourceScopeRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResourceScope not implemented")
}
func (UnimplementedResourceServiceServer) ListResources(context.Context, *empty.Empty) (*ListResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResources not implemented")
}
func (UnimplementedResourceServiceServer) ListClientScopes(context.Context, *ListClientScopesRequest) (*ListResourceScope, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClientScopes not implemented")
}
func (UnimplementedResourceServiceServer) GetResourceScope(context.Context, *GetResourceScopeRequest) (*ListResourceScope, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResourceScope not implemented")
}
func (UnimplementedResourceServiceServer) mustEmbedUnimplementedResourceServiceServer() {}

// UnsafeResourceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceServiceServer will
// result in compilation errors.
type UnsafeResourceServiceServer interface {
	mustEmbedUnimplementedResourceServiceServer()
}

func RegisterResourceServiceServer(s grpc.ServiceRegistrar, srv ResourceServiceServer) {
	s.RegisterService(&ResourceService_ServiceDesc, srv)
}

func _ResourceService_CreateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).CreateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService/CreateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).CreateResource(ctx, req.(*CreateResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_CreateResourceScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResourceScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).CreateResourceScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService/CreateResourceScope",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).CreateResourceScope(ctx, req.(*CreateResourceScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_ListResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).ListResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService/ListResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).ListResources(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_ListClientScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientScopesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).ListClientScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService/ListClientScopes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).ListClientScopes(ctx, req.(*ListClientScopesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_GetResourceScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourceScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).GetResourceScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService/GetResourceScope",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).GetResourceScope(ctx, req.(*GetResourceScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceService_ServiceDesc is the grpc.ServiceDesc for ResourceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ResourceService",
	HandlerType: (*ResourceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResource",
			Handler:    _ResourceService_CreateResource_Handler,
		},
		{
			MethodName: "CreateResourceScope",
			Handler:    _ResourceService_CreateResourceScope_Handler,
		},
		{
			MethodName: "ListResources",
			Handler:    _ResourceService_ListResources_Handler,
		},
		{
			MethodName: "ListClientScopes",
			Handler:    _ResourceService_ListClientScopes_Handler,
		},
		{
			MethodName: "GetResourceScope",
			Handler:    _ResourceService_GetResourceScope_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/resources.proto",
}
