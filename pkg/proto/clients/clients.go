package clients

import (
	"context"
	"io"

	"google.golang.org/grpc"

	"github.com/cro4k/authorize/pkg/proto/authorization"
)

type AuthorizationClientCloser interface {
	authorization.AuthorizationClient
	io.Closer
}

type authorizationClientCloser struct {
	authorization.AuthorizationClient
	io.Closer
}

func New(ctx context.Context, target string, options ...grpc.DialOption) (AuthorizationClientCloser, error) {
	return NewWithDialer(ctx, grpc.DialContext, target, options...)
}

func NewWithDialer(ctx context.Context, dialer GrpcDialer, target string, options ...grpc.DialOption) (AuthorizationClientCloser, error) {
	conn, err := dialer(ctx, target, options...)
	if err != nil {
		return nil, err
	}
	return &authorizationClientCloser{
		AuthorizationClient: authorization.NewAuthorizationClient(conn),
		Closer:              conn,
	}, nil
}
