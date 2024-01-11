// Code generated by protoc-gen-go-mock. DO NOT EDIT.
// versions:
// - protoc-gen-go-mock v0.2.4
// - protoc             v4.25.1
// source: .examples/petstore.proto

package petstore

import (
	context "context"
	fmt "fmt"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var _ PetStoreClient = &PetStoreMockClient{}

type PetStoreMockClient struct {
	PetStoreGetAllHandler
	PetStoreGetPetHandler
	PetStoreCreatePetHandler
	PetStoreUpdatePetHandler
	PetStoreDeletePetHandler
}

type PetStoreGetAllHandler func(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Pets, error)
type PetStoreGetPetHandler func(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Pet, error)
type PetStoreCreatePetHandler func(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Pet, error)
type PetStoreUpdatePetHandler func(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Pet, error)
type PetStoreDeletePetHandler func(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*emptypb.Empty, error)

func NewPetStoreMockClient() *PetStoreMockClient {
	return &PetStoreMockClient{}
}

func (c *PetStoreMockClient) GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Pets, error) {
	handler := c.PetStoreGetAllHandler
	if handler == nil {
		return nil, fmt.Errorf("no handler registered for GetAll")
	}

	return handler(ctx, in, opts...)
}

func (c *PetStoreMockClient) GetPet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Pet, error) {
	handler := c.PetStoreGetPetHandler
	if handler == nil {
		return nil, fmt.Errorf("no handler registered for GetPet")
	}

	return handler(ctx, in, opts...)
}

func (c *PetStoreMockClient) CreatePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Pet, error) {
	handler := c.PetStoreCreatePetHandler
	if handler == nil {
		return nil, fmt.Errorf("no handler registered for CreatePet")
	}

	return handler(ctx, in, opts...)
}

func (c *PetStoreMockClient) UpdatePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Pet, error) {
	handler := c.PetStoreUpdatePetHandler
	if handler == nil {
		return nil, fmt.Errorf("no handler registered for UpdatePet")
	}

	return handler(ctx, in, opts...)
}

func (c *PetStoreMockClient) DeletePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	handler := c.PetStoreDeletePetHandler
	if handler == nil {
		return nil, fmt.Errorf("no handler registered for DeletePet")
	}

	return handler(ctx, in, opts...)
}
