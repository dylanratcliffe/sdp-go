// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: auth0support.proto

package sdpconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	sdp_go "github.com/overmindtech/sdp-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// Auth0SupportName is the fully-qualified name of the Auth0Support service.
	Auth0SupportName = "auth0support.Auth0Support"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// Auth0SupportCreateUserProcedure is the fully-qualified name of the Auth0Support's CreateUser RPC.
	Auth0SupportCreateUserProcedure = "/auth0support.Auth0Support/CreateUser"
	// Auth0SupportKeepaliveSourcesProcedure is the fully-qualified name of the Auth0Support's
	// KeepaliveSources RPC.
	Auth0SupportKeepaliveSourcesProcedure = "/auth0support.Auth0Support/KeepaliveSources"
)

// Auth0SupportClient is a client for the auth0support.Auth0Support service.
type Auth0SupportClient interface {
	// create a new user on first login
	CreateUser(context.Context, *connect.Request[sdp_go.Auth0CreateUserRequest]) (*connect.Response[sdp_go.Auth0CreateUserResponse], error)
	// Updates sources to keep them running in the background. This is called on
	// login by auth0 to give us a chance to boot up sources while the app is
	// loading.
	KeepaliveSources(context.Context, *connect.Request[sdp_go.AdminKeepaliveSourcesRequest]) (*connect.Response[sdp_go.KeepaliveSourcesResponse], error)
}

// NewAuth0SupportClient constructs a client for the auth0support.Auth0Support service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuth0SupportClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) Auth0SupportClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &auth0SupportClient{
		createUser: connect.NewClient[sdp_go.Auth0CreateUserRequest, sdp_go.Auth0CreateUserResponse](
			httpClient,
			baseURL+Auth0SupportCreateUserProcedure,
			opts...,
		),
		keepaliveSources: connect.NewClient[sdp_go.AdminKeepaliveSourcesRequest, sdp_go.KeepaliveSourcesResponse](
			httpClient,
			baseURL+Auth0SupportKeepaliveSourcesProcedure,
			opts...,
		),
	}
}

// auth0SupportClient implements Auth0SupportClient.
type auth0SupportClient struct {
	createUser       *connect.Client[sdp_go.Auth0CreateUserRequest, sdp_go.Auth0CreateUserResponse]
	keepaliveSources *connect.Client[sdp_go.AdminKeepaliveSourcesRequest, sdp_go.KeepaliveSourcesResponse]
}

// CreateUser calls auth0support.Auth0Support.CreateUser.
func (c *auth0SupportClient) CreateUser(ctx context.Context, req *connect.Request[sdp_go.Auth0CreateUserRequest]) (*connect.Response[sdp_go.Auth0CreateUserResponse], error) {
	return c.createUser.CallUnary(ctx, req)
}

// KeepaliveSources calls auth0support.Auth0Support.KeepaliveSources.
func (c *auth0SupportClient) KeepaliveSources(ctx context.Context, req *connect.Request[sdp_go.AdminKeepaliveSourcesRequest]) (*connect.Response[sdp_go.KeepaliveSourcesResponse], error) {
	return c.keepaliveSources.CallUnary(ctx, req)
}

// Auth0SupportHandler is an implementation of the auth0support.Auth0Support service.
type Auth0SupportHandler interface {
	// create a new user on first login
	CreateUser(context.Context, *connect.Request[sdp_go.Auth0CreateUserRequest]) (*connect.Response[sdp_go.Auth0CreateUserResponse], error)
	// Updates sources to keep them running in the background. This is called on
	// login by auth0 to give us a chance to boot up sources while the app is
	// loading.
	KeepaliveSources(context.Context, *connect.Request[sdp_go.AdminKeepaliveSourcesRequest]) (*connect.Response[sdp_go.KeepaliveSourcesResponse], error)
}

// NewAuth0SupportHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuth0SupportHandler(svc Auth0SupportHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	auth0SupportCreateUserHandler := connect.NewUnaryHandler(
		Auth0SupportCreateUserProcedure,
		svc.CreateUser,
		opts...,
	)
	auth0SupportKeepaliveSourcesHandler := connect.NewUnaryHandler(
		Auth0SupportKeepaliveSourcesProcedure,
		svc.KeepaliveSources,
		opts...,
	)
	return "/auth0support.Auth0Support/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case Auth0SupportCreateUserProcedure:
			auth0SupportCreateUserHandler.ServeHTTP(w, r)
		case Auth0SupportKeepaliveSourcesProcedure:
			auth0SupportKeepaliveSourcesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuth0SupportHandler returns CodeUnimplemented from all methods.
type UnimplementedAuth0SupportHandler struct{}

func (UnimplementedAuth0SupportHandler) CreateUser(context.Context, *connect.Request[sdp_go.Auth0CreateUserRequest]) (*connect.Response[sdp_go.Auth0CreateUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth0support.Auth0Support.CreateUser is not implemented"))
}

func (UnimplementedAuth0SupportHandler) KeepaliveSources(context.Context, *connect.Request[sdp_go.AdminKeepaliveSourcesRequest]) (*connect.Response[sdp_go.KeepaliveSourcesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth0support.Auth0Support.KeepaliveSources is not implemented"))
}