package controller

import (
	"errors"

	"github.com/cro4k/common/crypto/hashutil"
	"github.com/cro4k/common/randx"
	"github.com/gin-gonic/gin"

	"github.com/cro4k/authorize/internal/dao"
	"github.com/cro4k/authorize/internal/db"
	"github.com/cro4k/authorize/internal/model"
	"github.com/cro4k/authorize/internal/model/resource"
	"github.com/cro4k/authorize/internal/service"
	"github.com/cro4k/authorize/server/api/ginx"
	"github.com/cro4k/authorize/utils/reg"
)

type LoginRequest struct {
	ginx.Empty
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID          string          `json:"id"`           //
	Token       string          `json:"token"`        //
	TokenExpire int             `json:"token_expire"` //
	Username    string          `json:"username"`     //
	Nickname    string          `json:"nickname"`     //
	Avatar      resource.Simple `json:"avatar"`       //
	BIO         string          `json:"bio"`          //
}

// Login
// @req [LoginRequest]
// @rsp [LoginResponse]
// @header Content-type:application/json
func Login(c *gin.Context) {
	ctx, err := ginx.With[LoginRequest](c).Bind()
	if err != nil {
		ctx.Logger().Error(err)
		ctx.FailError(err)
		return
	}
	acc, err := dao.Account(db.DB()).Find(ctx.Body.Username)
	if err != nil {
		ctx.Logger().Error(err)
		ctx.Fail(ginx.ErrIncorrectUsername)
		return
	}
	pwd := hashutil.MD5s(acc.Secret + ctx.Body.Password + acc.Secret)
	if acc.Password != pwd {
		ctx.Fail(ginx.ErrIncorrectUsername)
		return
	}
	token, err := service.Auth.GenToken(acc, ctx.CID)
	if err != nil {
		ctx.Logger().Error(err)
		ctx.Fail("login failed")
		return
	}
	var profile, _ = dao.Account(db.DB()).Profile(acc.ID)
	ctx.OK(LoginResponse{
		ID:          acc.ID,
		Token:       token,
		TokenExpire: 0,
		Username:    acc.Username,
		Nickname:    profile.Nickname,
		Avatar:      profile.Avatar.Simple(),
		BIO:         profile.Bio,
	})
}

const (
	CaptchaCellphone = 1
	CaptchaEmail     = 2
)

type RegisterRequest struct {
	ginx.Empty
	Username     string `json:"username"      doc:"must"` //
	Password     string `json:"password"      doc:"must"` //
	CaptchaType  int    `json:"captcha_type"  doc:"must"` //
	CaptchaID    string `json:"captcha_id"    doc:"must"` //
	CaptchaName  string `json:"captcha_name"  doc:"must"` //
	CaptchaValue string `json:"captcha_value" doc:"must"` //
}

func (r RegisterRequest) Valid(ctx *gin.Context) error {
	if r.Username == "" {
		return errors.New("username empty")
	}
	if !reg.Username.MatchString(r.Username) {
		return errors.New("invalid username")
	}
	if !reg.MD5.MatchString(r.Password) {
		return errors.New("invalid password")
	}

	if r.CaptchaName == "" {
		return errors.New("invalid cellphone/email")
	}
	if r.CaptchaID == "" || r.CaptchaValue == "" {
		return errors.New("captcha empty")
	}
	switch r.CaptchaType {
	case CaptchaCellphone:
		if !reg.Cellphone.MatchString(r.CaptchaName) {
			return errors.New("invalid cellphone/email")
		}
	case CaptchaEmail:
		if !reg.Email.MatchString(r.CaptchaName) {
			return errors.New("invalid cellphone/email")
		}
	default:
		return errors.New("invalid captcha type")
	}
	return nil
}

type RegisterResponse struct {
	ID string `json:"id"` //
}

// Register
// @req [RegisterRequest]
// @rsp [RegisterResponse]
func Register(c *gin.Context) {
	ctx, err := ginx.With[RegisterRequest](c).Bind()
	if err != nil {
		ctx.Logger().Error(err)
		ctx.FailError(err)
		return
	}
	acc := &model.Account{}
	acc.Username = ctx.Body.Username
	acc.Secret = randx.String(64)
	acc.Password = hashutil.MD5s(acc.Secret + ctx.Body.Password + acc.Secret)
	switch ctx.Body.CaptchaType {
	case CaptchaCellphone:
		acc.Cellphone = model.CipherText(ctx.Body.CaptchaName)
		acc.CellphoneHash = hashutil.MD5s(ctx.Body.CaptchaName)
	case CaptchaEmail:
		acc.Email = model.CipherText(ctx.Body.CaptchaName)
		acc.EmailHash = hashutil.MD5s(ctx.Body.CaptchaName)
	}

	var profile = &model.AccountProfile{
		Nickname: "用户_" + randx.String(6),
	}

	if err := dao.Account(db.DB()).Register(acc, profile); err != nil {
		ctx.Logger().Error(err)
		ctx.Fail()
		return
	}
	ctx.OK(RegisterResponse{
		ID: acc.ID,
	})
}

type LogoutRequest struct {
	ginx.Empty
	CID []string `json:"cid"`
}
type LogoutResponse struct{}

// Logout
// @header Content-type:application/json
// @header Authorization:TOKEN
// @req [LogoutRequest]
// @rsp [LogoutResponse]
func Logout(c *gin.Context) {
	ctx, err := ginx.With[LogoutRequest](c).Bind()
	if err != nil {
		ctx.Logger().Error(err)
		ctx.FailError(err)
		return
	}
	if len(ctx.Body.CID) > 0 {
		_ = service.Auth.Logout(ctx.UID, ctx.Body.CID...)
	} else {
		_ = service.Auth.Logout(ctx.UID, ctx.CID)
	}
	ctx.OK(LogoutResponse{})
}
