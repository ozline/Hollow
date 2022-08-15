// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.1
// - protoc             v3.21.5
// source: hollow/v1/user.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUsersLogin = "/user.v1.Users/Login"
const OperationUsersRegister = "/user.v1.Users/Register"

type UsersHTTPServer interface {
	Login(context.Context, *LoginUserRequest) (*LoginUserReply, error)
	Register(context.Context, *RegisterUserRequest) (*RegisterUserReply, error)
}

func RegisterUsersHTTPServer(s *http.Server, srv UsersHTTPServer) {
	r := s.Route("/")
	r.POST("/api/user/register", _Users_Register0_HTTP_Handler(srv))
	r.POST("/api/user/login", _Users_Login0_HTTP_Handler(srv))
}

func _Users_Register0_HTTP_Handler(srv UsersHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUsersRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterUserReply)
		return ctx.Result(200, reply)
	}
}

func _Users_Login0_HTTP_Handler(srv UsersHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUsersLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginUserReply)
		return ctx.Result(200, reply)
	}
}

type UsersHTTPClient interface {
	Login(ctx context.Context, req *LoginUserRequest, opts ...http.CallOption) (rsp *LoginUserReply, err error)
	Register(ctx context.Context, req *RegisterUserRequest, opts ...http.CallOption) (rsp *RegisterUserReply, err error)
}

type UsersHTTPClientImpl struct {
	cc *http.Client
}

func NewUsersHTTPClient(client *http.Client) UsersHTTPClient {
	return &UsersHTTPClientImpl{client}
}

func (c *UsersHTTPClientImpl) Login(ctx context.Context, in *LoginUserRequest, opts ...http.CallOption) (*LoginUserReply, error) {
	var out LoginUserReply
	pattern := "/api/user/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUsersLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UsersHTTPClientImpl) Register(ctx context.Context, in *RegisterUserRequest, opts ...http.CallOption) (*RegisterUserReply, error) {
	var out RegisterUserReply
	pattern := "/api/user/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUsersRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
