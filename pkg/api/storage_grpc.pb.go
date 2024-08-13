// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: api/storage.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Storage_Get_FullMethodName        = "/storage.Storage/Get"
	Storage_Find_FullMethodName       = "/storage.Storage/Find"
	Storage_Replace_FullMethodName    = "/storage.Storage/Replace"
	Storage_DeleteOne_FullMethodName  = "/storage.Storage/DeleteOne"
	Storage_DeleteMany_FullMethodName = "/storage.Storage/DeleteMany"
)

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
	Replace(ctx context.Context, in *ReplaceRequest, opts ...grpc.CallOption) (*ReplaceResponse, error)
	DeleteOne(ctx context.Context, in *DeleteOneRequest, opts ...grpc.CallOption) (*DeleteOneResponse, error)
	DeleteMany(ctx context.Context, in *DeleteManyRequest, opts ...grpc.CallOption) (*DeleteManyResponse, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, Storage_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindResponse)
	err := c.cc.Invoke(ctx, Storage_Find_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Replace(ctx context.Context, in *ReplaceRequest, opts ...grpc.CallOption) (*ReplaceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReplaceResponse)
	err := c.cc.Invoke(ctx, Storage_Replace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) DeleteOne(ctx context.Context, in *DeleteOneRequest, opts ...grpc.CallOption) (*DeleteOneResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteOneResponse)
	err := c.cc.Invoke(ctx, Storage_DeleteOne_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) DeleteMany(ctx context.Context, in *DeleteManyRequest, opts ...grpc.CallOption) (*DeleteManyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteManyResponse)
	err := c.cc.Invoke(ctx, Storage_DeleteMany_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility.
type StorageServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Find(context.Context, *FindRequest) (*FindResponse, error)
	Replace(context.Context, *ReplaceRequest) (*ReplaceResponse, error)
	DeleteOne(context.Context, *DeleteOneRequest) (*DeleteOneResponse, error)
	DeleteMany(context.Context, *DeleteManyRequest) (*DeleteManyResponse, error)
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStorageServer struct{}

func (UnimplementedStorageServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStorageServer) Find(context.Context, *FindRequest) (*FindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedStorageServer) Replace(context.Context, *ReplaceRequest) (*ReplaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Replace not implemented")
}
func (UnimplementedStorageServer) DeleteOne(context.Context, *DeleteOneRequest) (*DeleteOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOne not implemented")
}
func (UnimplementedStorageServer) DeleteMany(context.Context, *DeleteManyRequest) (*DeleteManyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMany not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}
func (UnimplementedStorageServer) testEmbeddedByValue()                 {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	// If the following call pancis, it indicates UnimplementedStorageServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Storage_ServiceDesc, srv)
}

func _Storage_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Find_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Find(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Replace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Replace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Replace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Replace(ctx, req.(*ReplaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_DeleteOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).DeleteOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_DeleteOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).DeleteOne(ctx, req.(*DeleteOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_DeleteMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).DeleteMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_DeleteMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).DeleteMany(ctx, req.(*DeleteManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Storage_ServiceDesc is the grpc.ServiceDesc for Storage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Storage_Get_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _Storage_Find_Handler,
		},
		{
			MethodName: "Replace",
			Handler:    _Storage_Replace_Handler,
		},
		{
			MethodName: "DeleteOne",
			Handler:    _Storage_DeleteOne_Handler,
		},
		{
			MethodName: "DeleteMany",
			Handler:    _Storage_DeleteMany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/storage.proto",
}
