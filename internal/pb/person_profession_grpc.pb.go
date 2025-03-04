// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/person_profession.proto

package pb

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
	PersonService_CreatePerson_FullMethodName = "/pb.PersonService/CreatePerson"
	PersonService_ListPerson_FullMethodName   = "/pb.PersonService/ListPerson"
)

// PersonServiceClient is the client API for PersonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonServiceClient interface {
	CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*Person, error)
	ListPerson(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*PersonList, error)
}

type personServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonServiceClient(cc grpc.ClientConnInterface) PersonServiceClient {
	return &personServiceClient{cc}
}

func (c *personServiceClient) CreatePerson(ctx context.Context, in *CreatePersonRequest, opts ...grpc.CallOption) (*Person, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Person)
	err := c.cc.Invoke(ctx, PersonService_CreatePerson_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) ListPerson(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*PersonList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PersonList)
	err := c.cc.Invoke(ctx, PersonService_ListPerson_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonServiceServer is the server API for PersonService service.
// All implementations must embed UnimplementedPersonServiceServer
// for forward compatibility.
type PersonServiceServer interface {
	CreatePerson(context.Context, *CreatePersonRequest) (*Person, error)
	ListPerson(context.Context, *Blank) (*PersonList, error)
	mustEmbedUnimplementedPersonServiceServer()
}

// UnimplementedPersonServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPersonServiceServer struct{}

func (UnimplementedPersonServiceServer) CreatePerson(context.Context, *CreatePersonRequest) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePerson not implemented")
}
func (UnimplementedPersonServiceServer) ListPerson(context.Context, *Blank) (*PersonList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPerson not implemented")
}
func (UnimplementedPersonServiceServer) mustEmbedUnimplementedPersonServiceServer() {}
func (UnimplementedPersonServiceServer) testEmbeddedByValue()                       {}

// UnsafePersonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonServiceServer will
// result in compilation errors.
type UnsafePersonServiceServer interface {
	mustEmbedUnimplementedPersonServiceServer()
}

func RegisterPersonServiceServer(s grpc.ServiceRegistrar, srv PersonServiceServer) {
	// If the following call pancis, it indicates UnimplementedPersonServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PersonService_ServiceDesc, srv)
}

func _PersonService_CreatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).CreatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PersonService_CreatePerson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).CreatePerson(ctx, req.(*CreatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_ListPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blank)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).ListPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PersonService_ListPerson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).ListPerson(ctx, req.(*Blank))
	}
	return interceptor(ctx, in, info, handler)
}

// PersonService_ServiceDesc is the grpc.ServiceDesc for PersonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PersonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PersonService",
	HandlerType: (*PersonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePerson",
			Handler:    _PersonService_CreatePerson_Handler,
		},
		{
			MethodName: "ListPerson",
			Handler:    _PersonService_ListPerson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/person_profession.proto",
}

const (
	ProfessionService_CreateProfession_FullMethodName = "/pb.ProfessionService/CreateProfession"
	ProfessionService_ListProfession_FullMethodName   = "/pb.ProfessionService/ListProfession"
)

// ProfessionServiceClient is the client API for ProfessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfessionServiceClient interface {
	CreateProfession(ctx context.Context, in *CreateProfessionRequest, opts ...grpc.CallOption) (*Profession, error)
	ListProfession(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*ProfessionList, error)
}

type professionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfessionServiceClient(cc grpc.ClientConnInterface) ProfessionServiceClient {
	return &professionServiceClient{cc}
}

func (c *professionServiceClient) CreateProfession(ctx context.Context, in *CreateProfessionRequest, opts ...grpc.CallOption) (*Profession, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Profession)
	err := c.cc.Invoke(ctx, ProfessionService_CreateProfession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *professionServiceClient) ListProfession(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*ProfessionList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProfessionList)
	err := c.cc.Invoke(ctx, ProfessionService_ListProfession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfessionServiceServer is the server API for ProfessionService service.
// All implementations must embed UnimplementedProfessionServiceServer
// for forward compatibility.
type ProfessionServiceServer interface {
	CreateProfession(context.Context, *CreateProfessionRequest) (*Profession, error)
	ListProfession(context.Context, *Blank) (*ProfessionList, error)
	mustEmbedUnimplementedProfessionServiceServer()
}

// UnimplementedProfessionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProfessionServiceServer struct{}

func (UnimplementedProfessionServiceServer) CreateProfession(context.Context, *CreateProfessionRequest) (*Profession, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfession not implemented")
}
func (UnimplementedProfessionServiceServer) ListProfession(context.Context, *Blank) (*ProfessionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProfession not implemented")
}
func (UnimplementedProfessionServiceServer) mustEmbedUnimplementedProfessionServiceServer() {}
func (UnimplementedProfessionServiceServer) testEmbeddedByValue()                           {}

// UnsafeProfessionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfessionServiceServer will
// result in compilation errors.
type UnsafeProfessionServiceServer interface {
	mustEmbedUnimplementedProfessionServiceServer()
}

func RegisterProfessionServiceServer(s grpc.ServiceRegistrar, srv ProfessionServiceServer) {
	// If the following call pancis, it indicates UnimplementedProfessionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProfessionService_ServiceDesc, srv)
}

func _ProfessionService_CreateProfession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfessionServiceServer).CreateProfession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfessionService_CreateProfession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfessionServiceServer).CreateProfession(ctx, req.(*CreateProfessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfessionService_ListProfession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blank)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfessionServiceServer).ListProfession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfessionService_ListProfession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfessionServiceServer).ListProfession(ctx, req.(*Blank))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfessionService_ServiceDesc is the grpc.ServiceDesc for ProfessionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfessionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ProfessionService",
	HandlerType: (*ProfessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProfession",
			Handler:    _ProfessionService_CreateProfession_Handler,
		},
		{
			MethodName: "ListProfession",
			Handler:    _ProfessionService_ListProfession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/person_profession.proto",
}
