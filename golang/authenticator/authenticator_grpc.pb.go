// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: authenticator.proto

package authenticator

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

const (
	AuthenticatorService_AccountCreate_FullMethodName = "/mex.authenticator.v1.AuthenticatorService/AccountCreate"
	AuthenticatorService_AccountDelete_FullMethodName = "/mex.authenticator.v1.AuthenticatorService/AccountDelete"
	AuthenticatorService_TokenVerify_FullMethodName   = "/mex.authenticator.v1.AuthenticatorService/TokenVerify"
)

// AuthenticatorServiceClient is the client API for AuthenticatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticatorServiceClient interface {
	AccountCreate(ctx context.Context, in *AccountCreateRequest, opts ...grpc.CallOption) (*AccountCreateResponse, error)
	AccountDelete(ctx context.Context, in *AccountDeleteRequest, opts ...grpc.CallOption) (*AccountDeleteResponse, error)
	TokenVerify(ctx context.Context, in *TokenVerifyRequest, opts ...grpc.CallOption) (*TokenVerifyResponse, error)
}

type authenticatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticatorServiceClient(cc grpc.ClientConnInterface) AuthenticatorServiceClient {
	return &authenticatorServiceClient{cc}
}

func (c *authenticatorServiceClient) AccountCreate(ctx context.Context, in *AccountCreateRequest, opts ...grpc.CallOption) (*AccountCreateResponse, error) {
	out := new(AccountCreateResponse)
	err := c.cc.Invoke(ctx, AuthenticatorService_AccountCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticatorServiceClient) AccountDelete(ctx context.Context, in *AccountDeleteRequest, opts ...grpc.CallOption) (*AccountDeleteResponse, error) {
	out := new(AccountDeleteResponse)
	err := c.cc.Invoke(ctx, AuthenticatorService_AccountDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticatorServiceClient) TokenVerify(ctx context.Context, in *TokenVerifyRequest, opts ...grpc.CallOption) (*TokenVerifyResponse, error) {
	out := new(TokenVerifyResponse)
	err := c.cc.Invoke(ctx, AuthenticatorService_TokenVerify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticatorServiceServer is the server API for AuthenticatorService service.
// All implementations must embed UnimplementedAuthenticatorServiceServer
// for forward compatibility
type AuthenticatorServiceServer interface {
	AccountCreate(context.Context, *AccountCreateRequest) (*AccountCreateResponse, error)
	AccountDelete(context.Context, *AccountDeleteRequest) (*AccountDeleteResponse, error)
	TokenVerify(context.Context, *TokenVerifyRequest) (*TokenVerifyResponse, error)
	mustEmbedUnimplementedAuthenticatorServiceServer()
}

// UnimplementedAuthenticatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticatorServiceServer struct {
}

func (UnimplementedAuthenticatorServiceServer) AccountCreate(context.Context, *AccountCreateRequest) (*AccountCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountCreate not implemented")
}
func (UnimplementedAuthenticatorServiceServer) AccountDelete(context.Context, *AccountDeleteRequest) (*AccountDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountDelete not implemented")
}
func (UnimplementedAuthenticatorServiceServer) TokenVerify(context.Context, *TokenVerifyRequest) (*TokenVerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenVerify not implemented")
}
func (UnimplementedAuthenticatorServiceServer) mustEmbedUnimplementedAuthenticatorServiceServer() {}

// UnsafeAuthenticatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticatorServiceServer will
// result in compilation errors.
type UnsafeAuthenticatorServiceServer interface {
	mustEmbedUnimplementedAuthenticatorServiceServer()
}

func RegisterAuthenticatorServiceServer(s grpc.ServiceRegistrar, srv AuthenticatorServiceServer) {
	s.RegisterService(&AuthenticatorService_ServiceDesc, srv)
}

func _AuthenticatorService_AccountCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServiceServer).AccountCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticatorService_AccountCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServiceServer).AccountCreate(ctx, req.(*AccountCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticatorService_AccountDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServiceServer).AccountDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticatorService_AccountDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServiceServer).AccountDelete(ctx, req.(*AccountDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticatorService_TokenVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServiceServer).TokenVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticatorService_TokenVerify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServiceServer).TokenVerify(ctx, req.(*TokenVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticatorService_ServiceDesc is the grpc.ServiceDesc for AuthenticatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mex.authenticator.v1.AuthenticatorService",
	HandlerType: (*AuthenticatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AccountCreate",
			Handler:    _AuthenticatorService_AccountCreate_Handler,
		},
		{
			MethodName: "AccountDelete",
			Handler:    _AuthenticatorService_AccountDelete_Handler,
		},
		{
			MethodName: "TokenVerify",
			Handler:    _AuthenticatorService_TokenVerify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authenticator.proto",
}
