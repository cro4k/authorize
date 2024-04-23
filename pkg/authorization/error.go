package authorization

import "errors"

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
	ErrNotFound     = errors.New("not found")
)
