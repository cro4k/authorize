package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/cro4k/authorize/internal/model"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

type service struct {
	storage Storage
}

func NewService() *service {
	return &service{storage: newStorage()}
}

func (s *service) GenToken(acc *model.Account, cid string) (string, error) {
	nonce, err := s.storage.Gen(acc.ID, cid)
	if err != nil {
		return "", err
	}
	cla := &Claims{}
	cla.UID = acc.ID
	cla.CID = cid
	cla.Username = acc.Username
	cla.CertificateStatus = int(acc.CertificateStatus)
	//if len(org) > 0 && org[0] != "" {
	//	organization, err := dao.Organization.Get(acc.ID, org[0])
	//	if err != nil {
	//		return "", err
	//	}
	//	cla.OrganizationID = org[0]
	//	cla.OrganizationRole, _ = dao.Organization.Role(acc.ID, org[0])
	//	cla.OrganizationCertificateStatus = int(organization.CertificateStatus)
	//}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
	return token.SignedString([]byte(nonce))
}

func (s *service) Verify(tokenStr string, cid string) (*Claims, error) {
	temp := strings.Split(tokenStr, ".")
	if len(temp) != 3 {
		return nil, errors.New("invalid token")
	}
	claimsData, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(temp[1])
	if err != nil {
		return nil, err
	}
	var claims = new(Claims)
	if err := json.Unmarshal(claimsData, claims); err != nil {
		return nil, err
	}
	if claims.CID != cid {
		return nil, errors.New("invalid token")
	}
	nonce, err := s.storage.Get(claims.UID, claims.CID)
	if err != nil {
		return nil, err
	}
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(nonce), nil
	})
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

func (s *service) Logout(uid string, cid ...string) error {
	return s.storage.Del(uid, cid...)
}
