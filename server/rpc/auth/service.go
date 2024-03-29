package auth

import (
	"context"

	"github.com/cro4k/authorize/internal/dao"
	"github.com/cro4k/authorize/internal/db"
	"github.com/cro4k/authorize/internal/service"
	"github.com/cro4k/authorize/rpc/authrpc"
)

type authService struct {
	authrpc.UnimplementedAuthServiceServer
}

var _ authrpc.AuthServiceServer = new(authService)

func (s *authService) VerifyToken(ctx context.Context, req *authrpc.VerifyTokenRequest) (*authrpc.VerifyTokenResponse, error) {
	claims, err := service.Auth.Verify(req.Token, req.Cid)
	if err != nil {
		return nil, err
	}
	profile, err := dao.Account(db.DB()).Profile(claims.UID)
	if err != nil {
		return nil, err
	}

	rsp := &authrpc.VerifyTokenResponse{}
	rsp.Profile = &authrpc.Profile{
		Id:       claims.UID,
		Nickname: profile.Nickname,
		Avatar:   profile.Avatar.Thumbnail,
		Gender:   int32(profile.Gender),
		Bio:      profile.Bio,
	}
	return rsp, nil
}

func (s *authService) AccountInfo(ctx context.Context, req *authrpc.AccountInfoRequest) (*authrpc.AccountInfoResponse, error) {
	claims, err := service.Auth.Verify(req.Token, req.Cid)
	if err != nil {
		return nil, err
	}
	acc, err := dao.Account(db.DB()).GetByID(claims.Id)
	if err != nil {
		return nil, err
	}
	rsp := &authrpc.AccountInfoResponse{}
	rsp.Info = &authrpc.AccountInfo{
		Id:                claims.UID,
		Username:          acc.Username,
		Cellphone:         string(acc.Cellphone),
		Email:             string(acc.Email),
		CertificateStatus: int32(acc.CertificateStatus),
	}
	return rsp, nil
}

func NewAuthService() authrpc.AuthServiceServer {
	return new(authService)
}
