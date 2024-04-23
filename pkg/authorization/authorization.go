package authorization

import (
	"context"
	"time"
)

type option struct {
	storage *secretStorage
	maxAge  time.Duration
}

type Option func(*option)

func applyOptions(opts []Option) *option {
	var o = &option{maxAge: 7 * 24 * time.Hour}
	for _, opt := range opts {
		opt(o)
	}
	if o.storage == nil {
		o.storage = newSecretStorage(&memoryStorage{}, 5, o.maxAge)
	}
	return o
}

func WithMaxAge(maxAge time.Duration) Option {
	return func(o *option) {
		o.maxAge = maxAge
	}
}

func WithSecretStorage(storage Storage) Option {
	return func(o *option) {
		o.storage = newSecretStorage(storage, 5, o.maxAge)
	}
}

type Authorization struct {
	*option
}

func NewAuthorization(opts ...Option) *Authorization {
	return &Authorization{
		option: applyOptions(opts),
	}
}

func (a *Authorization) Authorize(ctx context.Context, id string, clientId string) (string, error) {
	item, err := a.storage.Put(ctx, id, clientId)
	if err != nil {
		return "", err
	}
	token := NewToken(id, WithClientID(clientId), WithExpireAt(time.Now().Add(a.maxAge)))
	return token.Encode(item.Secret)
}

func (a *Authorization) Validate(ctx context.Context, tokenString string) (*Token, error) {
	token, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	item, err := a.storage.Get(ctx, token.UserID, token.ClientID)
	if err != nil {
		return nil, err
	}
	err = token.Validate(item.Secret)
	return token, err
}
