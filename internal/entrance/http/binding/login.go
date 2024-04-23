package binding

import (
	"context"
	"github.com/go-playground/validator/v10"
)

var _validator_ = validator.New()

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (request *LoginRequest) Validate(ctx context.Context) error {
	return _validator_.StructCtx(ctx, request)
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (request *RegisterRequest) Validate(ctx context.Context) error {
	return _validator_.StructCtx(ctx, request)
}

type RegisterResponse struct {
}
