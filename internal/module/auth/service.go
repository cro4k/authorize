package auth

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-chocolate/chocolate/pkg/kv"

	"github.com/cro4k/authorize/pkg/authorization"
)

type Service interface {
	Token(ctx context.Context, id int64, client string) (string, error)
	Validate(ctx context.Context, token string) (int64, error)
}

type UnimplementedService struct{}

func (s *UnimplementedService) Token(ctx context.Context, id int64, client string) (string, error) {
	return "", errors.New("unimplemented")
}
func (s *UnimplementedService) Validate(ctx context.Context, token string) (int64, error) {
	return 0, errors.New("unimplemented")
}

func NewService(kv kv.Storage) Service {
	return &service{auth: authorization.NewAuthorization(authorization.WithSecretStorage(kv))}
}

type service struct {
	auth *authorization.Authorization
}

func (s *service) Token(ctx context.Context, id int64, client string) (string, error) {
	return s.auth.Authorize(ctx, strconv.FormatInt(id, 10), client)
}

func (s *service) Validate(ctx context.Context, tokenString string) (int64, error) {
	token, err := s.auth.Validate(ctx, tokenString)
	if err != nil {
		return 0, err
	}
	id, _ := strconv.ParseInt(token.UserID, 10, 64)
	return id, nil
}
