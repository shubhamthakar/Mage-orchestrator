// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Magedatamonk

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

// UserNameClient is the client API for UserName service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserNameClient interface {
	GetMockUserData(ctx context.Context, in *Username, opts ...grpc.CallOption) (*User, error)
}

type userNameClient struct {
	cc grpc.ClientConnInterface
}

func NewUserNameClient(cc grpc.ClientConnInterface) UserNameClient {
	return &userNameClient{cc}
}

func (c *userNameClient) GetMockUserData(ctx context.Context, in *Username, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/dummy_data_service.UserName/GetMockUserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserNameServer is the server API for UserName service.
// All implementations must embed UnimplementedUserNameServer
// for forward compatibility
type UserNameServer interface {
	GetMockUserData(context.Context, *Username) (*User, error)
	mustEmbedUnimplementedUserNameServer()
}

// UnimplementedUserNameServer must be embedded to have forward compatible implementations.
type UnimplementedUserNameServer struct {
}

func (UnimplementedUserNameServer) GetMockUserData(context.Context, *Username) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMockUserData not implemented")
}
func (UnimplementedUserNameServer) mustEmbedUnimplementedUserNameServer() {}

// UnsafeUserNameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserNameServer will
// result in compilation errors.
type UnsafeUserNameServer interface {
	mustEmbedUnimplementedUserNameServer()
}

func RegisterUserNameServer(s grpc.ServiceRegistrar, srv UserNameServer) {
	s.RegisterService(&UserName_ServiceDesc, srv)
}

func _UserName_GetMockUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Username)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNameServer).GetMockUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dummy_data_service.UserName/GetMockUserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNameServer).GetMockUserData(ctx, req.(*Username))
	}
	return interceptor(ctx, in, info, handler)
}

// UserName_ServiceDesc is the grpc.ServiceDesc for UserName service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserName_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dummy_data_service.UserName",
	HandlerType: (*UserNameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMockUserData",
			Handler:    _UserName_GetMockUserData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/dummy_data_service.proto",
}
