package handler

import (
	"context"

	"github.com/cro4k/authorize/internal/module"
	"github.com/cro4k/authorize/pkg/proto/authorization"
)

func Validate(ctx context.Context, req *authorization.ValidateRequest) (*authorization.ValidateResponse, error) {
	id, err := module.GetAuthService().Validate(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	return &authorization.ValidateResponse{Id: id}, nil
}
