package authorization

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"strings"
	"time"
)

type TokenHeader struct {
	UserID    string
	ClientID  string
	CreatedAt int64
	ExpireAt  int64
	Nonce     string
}

type Token struct {
	TokenHeader
	signature string
	text      string
	claims    any

	headerBytes []byte
	claimsBytes []byte

	sum hash.Hash
}

type TokenOption func(t *Token)

func applyTokenOptions(t *Token, options []TokenOption) {
	for _, option := range options {
		option(t)
	}
}

func WithClaims(claims any) TokenOption {
	return func(t *Token) {
		t.claims = claims
	}
}

func WithClientID(clientID string) TokenOption {
	return func(t *Token) {
		t.ClientID = clientID
	}
}

func WithExpireAt(expire time.Time) TokenOption {
	return func(t *Token) {
		t.ExpireAt = expire.UnixMilli()
	}
}

func WithSum(sum hash.Hash) TokenOption {
	return func(t *Token) {
		t.sum = sum
	}
}

func NewToken(id string, options ...TokenOption) *Token {
	t := &Token{
		TokenHeader: TokenHeader{
			UserID:    id,
			CreatedAt: time.Now().UnixMilli(),
			ExpireAt:  time.Now().Add(7 * 24 * time.Hour).UnixMilli(),
		},
		sum: md5.New(),
	}
	applyTokenOptions(t, options)
	return t
}

func (t *Token) Encode(secret string) (string, error) {
	t.text = ""
	b64 := base64.StdEncoding.WithPadding(base64.NoPadding)
	if b, err := json.Marshal(t.TokenHeader); err == nil {
		t.headerBytes = b
		t.text += b64.EncodeToString(b) + "."
	} else {
		return "", err
	}
	if b, err := json.Marshal(t.claims); err == nil {
		t.claimsBytes = b
		t.text += b64.EncodeToString(b) + "."
	} else {
		return "", err
	}
	t.signature = t.sign(secret)
	t.text += t.signature
	return t.text, nil
}

func ParseToken(tokenString string, options ...TokenOption) (*Token, error) {
	token := &Token{sum: md5.New(), text: tokenString}
	applyTokenOptions(token, options)

	segments := strings.Split(tokenString, ".")
	if len(segments) != 3 {
		return nil, ErrInvalidToken
	}
	b64 := base64.StdEncoding.WithPadding(base64.NoPadding)
	if b, err := b64.DecodeString(segments[0]); err == nil {
		if err := json.Unmarshal(b, &token.TokenHeader); err != nil {
			return nil, ErrInvalidToken
		}
		token.headerBytes = b
	} else {
		return nil, fmt.Errorf("decode token header on error: %v", err)
	}

	if token.claims != nil {
		if b, err := b64.DecodeString(segments[1]); err == nil {
			if err := json.Unmarshal(b, token.claims); err != nil {
				return nil, ErrInvalidToken
			}
			token.claimsBytes = b
		} else {
			return nil, fmt.Errorf("decode token claims on error: %v", err)
		}
	}
	token.signature = segments[2]
	return token, nil
}

func (t *Token) sign(secret string) string {
	t.sum.Reset()
	t.sum.Write(t.headerBytes)
	t.sum.Write(t.claimsBytes)
	t.sum.Write([]byte(secret))
	return hex.EncodeToString(t.sum.Sum(nil))
}

func (t *Token) Validate(secret string) error {
	signature := t.sign(secret)
	if t.signature != signature {
		return ErrInvalidToken
	}
	if time.Now().UnixMilli() >= t.TokenHeader.ExpireAt {
		return ErrTokenExpired
	}
	return nil
}
