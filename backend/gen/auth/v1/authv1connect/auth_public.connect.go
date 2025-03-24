// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: auth/v1/auth_public.proto

package authv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	v1 "opendrive/gen/auth/v1"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// PublicAuthServiceName is the fully-qualified name of the PublicAuthService service.
	PublicAuthServiceName = "auth.v1.PublicAuthService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PublicAuthServiceLoginProcedure is the fully-qualified name of the PublicAuthService's Login RPC.
	PublicAuthServiceLoginProcedure = "/auth.v1.PublicAuthService/Login"
	// PublicAuthServiceRegisterProcedure is the fully-qualified name of the PublicAuthService's
	// Register RPC.
	PublicAuthServiceRegisterProcedure = "/auth.v1.PublicAuthService/Register"
	// PublicAuthServiceRefreshTokenProcedure is the fully-qualified name of the PublicAuthService's
	// RefreshToken RPC.
	PublicAuthServiceRefreshTokenProcedure = "/auth.v1.PublicAuthService/RefreshToken"
	// PublicAuthServiceWhoamiProcedure is the fully-qualified name of the PublicAuthService's Whoami
	// RPC.
	PublicAuthServiceWhoamiProcedure = "/auth.v1.PublicAuthService/Whoami"
)

// PublicAuthServiceClient is a client for the auth.v1.PublicAuthService service.
type PublicAuthServiceClient interface {
	// 用户登录
	Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)
	// 用户注册
	Register(context.Context, *connect.Request[v1.RegisterRequest]) (*connect.Response[v1.RegisterResponse], error)
	// 刷新令牌
	RefreshToken(context.Context, *connect.Request[v1.RefreshTokenRequest]) (*connect.Response[v1.RefreshTokenResponse], error)
	// 获取当前用户信息
	Whoami(context.Context, *connect.Request[v1.WhoamiRequest]) (*connect.Response[v1.WhoamiResponse], error)
}

// NewPublicAuthServiceClient constructs a client for the auth.v1.PublicAuthService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPublicAuthServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PublicAuthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	publicAuthServiceMethods := v1.File_auth_v1_auth_public_proto.Services().ByName("PublicAuthService").Methods()
	return &publicAuthServiceClient{
		login: connect.NewClient[v1.LoginRequest, v1.LoginResponse](
			httpClient,
			baseURL+PublicAuthServiceLoginProcedure,
			connect.WithSchema(publicAuthServiceMethods.ByName("Login")),
			connect.WithClientOptions(opts...),
		),
		register: connect.NewClient[v1.RegisterRequest, v1.RegisterResponse](
			httpClient,
			baseURL+PublicAuthServiceRegisterProcedure,
			connect.WithSchema(publicAuthServiceMethods.ByName("Register")),
			connect.WithClientOptions(opts...),
		),
		refreshToken: connect.NewClient[v1.RefreshTokenRequest, v1.RefreshTokenResponse](
			httpClient,
			baseURL+PublicAuthServiceRefreshTokenProcedure,
			connect.WithSchema(publicAuthServiceMethods.ByName("RefreshToken")),
			connect.WithClientOptions(opts...),
		),
		whoami: connect.NewClient[v1.WhoamiRequest, v1.WhoamiResponse](
			httpClient,
			baseURL+PublicAuthServiceWhoamiProcedure,
			connect.WithSchema(publicAuthServiceMethods.ByName("Whoami")),
			connect.WithClientOptions(opts...),
		),
	}
}

// publicAuthServiceClient implements PublicAuthServiceClient.
type publicAuthServiceClient struct {
	login        *connect.Client[v1.LoginRequest, v1.LoginResponse]
	register     *connect.Client[v1.RegisterRequest, v1.RegisterResponse]
	refreshToken *connect.Client[v1.RefreshTokenRequest, v1.RefreshTokenResponse]
	whoami       *connect.Client[v1.WhoamiRequest, v1.WhoamiResponse]
}

// Login calls auth.v1.PublicAuthService.Login.
func (c *publicAuthServiceClient) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	return c.login.CallUnary(ctx, req)
}

// Register calls auth.v1.PublicAuthService.Register.
func (c *publicAuthServiceClient) Register(ctx context.Context, req *connect.Request[v1.RegisterRequest]) (*connect.Response[v1.RegisterResponse], error) {
	return c.register.CallUnary(ctx, req)
}

// RefreshToken calls auth.v1.PublicAuthService.RefreshToken.
func (c *publicAuthServiceClient) RefreshToken(ctx context.Context, req *connect.Request[v1.RefreshTokenRequest]) (*connect.Response[v1.RefreshTokenResponse], error) {
	return c.refreshToken.CallUnary(ctx, req)
}

// Whoami calls auth.v1.PublicAuthService.Whoami.
func (c *publicAuthServiceClient) Whoami(ctx context.Context, req *connect.Request[v1.WhoamiRequest]) (*connect.Response[v1.WhoamiResponse], error) {
	return c.whoami.CallUnary(ctx, req)
}

// PublicAuthServiceHandler is an implementation of the auth.v1.PublicAuthService service.
type PublicAuthServiceHandler interface {
	// 用户登录
	Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)
	// 用户注册
	Register(context.Context, *connect.Request[v1.RegisterRequest]) (*connect.Response[v1.RegisterResponse], error)
	// 刷新令牌
	RefreshToken(context.Context, *connect.Request[v1.RefreshTokenRequest]) (*connect.Response[v1.RefreshTokenResponse], error)
	// 获取当前用户信息
	Whoami(context.Context, *connect.Request[v1.WhoamiRequest]) (*connect.Response[v1.WhoamiResponse], error)
}

// NewPublicAuthServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPublicAuthServiceHandler(svc PublicAuthServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	publicAuthServiceMethods := v1.File_auth_v1_auth_public_proto.Services().ByName("PublicAuthService").Methods()
	publicAuthServiceLoginHandler := connect.NewUnaryHandler(
		PublicAuthServiceLoginProcedure,
		svc.Login,
		connect.WithSchema(publicAuthServiceMethods.ByName("Login")),
		connect.WithHandlerOptions(opts...),
	)
	publicAuthServiceRegisterHandler := connect.NewUnaryHandler(
		PublicAuthServiceRegisterProcedure,
		svc.Register,
		connect.WithSchema(publicAuthServiceMethods.ByName("Register")),
		connect.WithHandlerOptions(opts...),
	)
	publicAuthServiceRefreshTokenHandler := connect.NewUnaryHandler(
		PublicAuthServiceRefreshTokenProcedure,
		svc.RefreshToken,
		connect.WithSchema(publicAuthServiceMethods.ByName("RefreshToken")),
		connect.WithHandlerOptions(opts...),
	)
	publicAuthServiceWhoamiHandler := connect.NewUnaryHandler(
		PublicAuthServiceWhoamiProcedure,
		svc.Whoami,
		connect.WithSchema(publicAuthServiceMethods.ByName("Whoami")),
		connect.WithHandlerOptions(opts...),
	)
	return "/auth.v1.PublicAuthService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PublicAuthServiceLoginProcedure:
			publicAuthServiceLoginHandler.ServeHTTP(w, r)
		case PublicAuthServiceRegisterProcedure:
			publicAuthServiceRegisterHandler.ServeHTTP(w, r)
		case PublicAuthServiceRefreshTokenProcedure:
			publicAuthServiceRefreshTokenHandler.ServeHTTP(w, r)
		case PublicAuthServiceWhoamiProcedure:
			publicAuthServiceWhoamiHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPublicAuthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPublicAuthServiceHandler struct{}

func (UnimplementedPublicAuthServiceHandler) Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.PublicAuthService.Login is not implemented"))
}

func (UnimplementedPublicAuthServiceHandler) Register(context.Context, *connect.Request[v1.RegisterRequest]) (*connect.Response[v1.RegisterResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.PublicAuthService.Register is not implemented"))
}

func (UnimplementedPublicAuthServiceHandler) RefreshToken(context.Context, *connect.Request[v1.RefreshTokenRequest]) (*connect.Response[v1.RefreshTokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.PublicAuthService.RefreshToken is not implemented"))
}

func (UnimplementedPublicAuthServiceHandler) Whoami(context.Context, *connect.Request[v1.WhoamiRequest]) (*connect.Response[v1.WhoamiResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.PublicAuthService.Whoami is not implemented"))
}
