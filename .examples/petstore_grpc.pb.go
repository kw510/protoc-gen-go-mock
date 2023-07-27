// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: example/petstore.proto

package petstore

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

const (
	PetStore_GetAll_FullMethodName    = "/petstore.PetStore/GetAll"
	PetStore_GetPet_FullMethodName    = "/petstore.PetStore/GetPet"
	PetStore_CreatePet_FullMethodName = "/petstore.PetStore/CreatePet"
	PetStore_UpdatePet_FullMethodName = "/petstore.PetStore/UpdatePet"
	PetStore_DeletePet_FullMethodName = "/petstore.PetStore/DeletePet"
)

// PetStoreClient is the client API for PetStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PetStoreClient interface {
	GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Pets, error)
	GetPet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Pet, error)
	CreatePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Pet, error)
	UpdatePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Pet, error)
	DeletePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type petStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewPetStoreClient(cc grpc.ClientConnInterface) PetStoreClient {
	return &petStoreClient{cc}
}

func (c *petStoreClient) GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Pets, error) {
	out := new(Pets)
	err := c.cc.Invoke(ctx, PetStore_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petStoreClient) GetPet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Pet, error) {
	out := new(Pet)
	err := c.cc.Invoke(ctx, PetStore_GetPet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petStoreClient) CreatePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Pet, error) {
	out := new(Pet)
	err := c.cc.Invoke(ctx, PetStore_CreatePet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petStoreClient) UpdatePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Pet, error) {
	out := new(Pet)
	err := c.cc.Invoke(ctx, PetStore_UpdatePet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petStoreClient) DeletePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PetStore_DeletePet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetStoreServer is the server API for PetStore service.
// All implementations must embed UnimplementedPetStoreServer
// for forward compatibility
type PetStoreServer interface {
	GetAll(context.Context, *emptypb.Empty) (*Pets, error)
	GetPet(context.Context, *Pet) (*Pet, error)
	CreatePet(context.Context, *Pet) (*Pet, error)
	UpdatePet(context.Context, *Pet) (*Pet, error)
	DeletePet(context.Context, *Pet) (*emptypb.Empty, error)
	mustEmbedUnimplementedPetStoreServer()
}

// UnimplementedPetStoreServer must be embedded to have forward compatible implementations.
type UnimplementedPetStoreServer struct {
}

func (UnimplementedPetStoreServer) GetAll(context.Context, *emptypb.Empty) (*Pets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedPetStoreServer) GetPet(context.Context, *Pet) (*Pet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPet not implemented")
}
func (UnimplementedPetStoreServer) CreatePet(context.Context, *Pet) (*Pet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePet not implemented")
}
func (UnimplementedPetStoreServer) UpdatePet(context.Context, *Pet) (*Pet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePet not implemented")
}
func (UnimplementedPetStoreServer) DeletePet(context.Context, *Pet) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePet not implemented")
}
func (UnimplementedPetStoreServer) mustEmbedUnimplementedPetStoreServer() {}

// UnsafePetStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PetStoreServer will
// result in compilation errors.
type UnsafePetStoreServer interface {
	mustEmbedUnimplementedPetStoreServer()
}

func RegisterPetStoreServer(s grpc.ServiceRegistrar, srv PetStoreServer) {
	s.RegisterService(&PetStore_ServiceDesc, srv)
}

func _PetStore_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetStore_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServer).GetAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetStore_GetPet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServer).GetPet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetStore_GetPet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServer).GetPet(ctx, req.(*Pet))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetStore_CreatePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServer).CreatePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetStore_CreatePet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServer).CreatePet(ctx, req.(*Pet))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetStore_UpdatePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServer).UpdatePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetStore_UpdatePet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServer).UpdatePet(ctx, req.(*Pet))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetStore_DeletePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServer).DeletePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetStore_DeletePet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServer).DeletePet(ctx, req.(*Pet))
	}
	return interceptor(ctx, in, info, handler)
}

// PetStore_ServiceDesc is the grpc.ServiceDesc for PetStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PetStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "petstore.PetStore",
	HandlerType: (*PetStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _PetStore_GetAll_Handler,
		},
		{
			MethodName: "GetPet",
			Handler:    _PetStore_GetPet_Handler,
		},
		{
			MethodName: "CreatePet",
			Handler:    _PetStore_CreatePet_Handler,
		},
		{
			MethodName: "UpdatePet",
			Handler:    _PetStore_UpdatePet_Handler,
		},
		{
			MethodName: "DeletePet",
			Handler:    _PetStore_DeletePet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/petstore.proto",
}