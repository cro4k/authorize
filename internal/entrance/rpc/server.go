package rpc

import (
	"context"

	"github.com/cro4k/authorize/internal/entrance/rpc/handler"
	"github.com/cro4k/authorize/pkg/proto/authorization"
)

type AuthorizeServer struct {
	authorization.UnimplementedAuthorizationServer
}

func (s *AuthorizeServer) Validate(ctx context.Context, req *authorization.ValidateRequest) (*authorization.ValidateResponse, error) {
	return handler.Validate(ctx, req)
}
