package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-chocolate/chocolate/pkg/chocolate/errorx"
	"github.com/cro4k/authorize/internal/dao/model"
	"github.com/cro4k/authorize/internal/dao/operator"
	"github.com/cro4k/authorize/internal/entrance/http/binding"
	"github.com/cro4k/authorize/internal/module"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"

	"github.com/go-chocolate/chocolate/pkg/chocolate/chocohttp/chocomux"
)

func Login(ctx *gin.Context) {
	login(chocomux.WithStd(ctx.Writer, ctx.Request), ctx)
}

func login(ctx chocomux.Context, ginCtx *gin.Context) {
	request := new(binding.LoginRequest)
	if err := ctx.Bind(request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorx.Code(400, err.Error()))
		return
	}
	user := &model.User{Username: request.Username}
	op := operator.NewUserOperator(ctx, user)
	if err := op.Load(ctx); err != nil {
		ctx.Error(err)
		return
	}
	if user.Password != request.Password {
		ctx.OkJSON(errorx.Code(0, "invalid username or password"))
		return
	}
	token, err := module.GetAuthService().Token(ctx, user.ID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.OkJSON(&binding.LoginResponse{Token: token})
}

func Register(ctx *gin.Context) {
	register(chocomux.WithStd(ctx.Writer, ctx.Request), ctx)
}

func register(ctx chocomux.Context, ginCtx *gin.Context) {
	request := new(binding.LoginRequest)
	if err := ctx.Bind(request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorx.Code(400, err.Error()))
		return
	}
	user := &model.User{Username: request.Username}
	op := operator.NewUserOperator(ctx, user)
	if err := op.Load(ctx); err != nil && err != gorm.ErrRecordNotFound {
		ctx.Error(err)
		return
	}
	if user.ID > 0 {
		ctx.OkJSON(errorx.Code(0, "username has been used"))
		return
	}
	user.ID = int64(uuid.New().ID())
	user.Username = request.Username
	user.Password = request.Password
	if _, err := op.Create(ctx); err != nil {
		ctx.Error(err)
		return
	}
	ctx.OkJSON(&binding.RegisterResponse{})
}

func Hello(ctx *gin.Context) {
	hello(chocomux.WithStd(ctx.Writer, ctx.Request), ctx)
}

func hello(ctx chocomux.Context, ginCtx *gin.Context) {
	id, ok := ginCtx.Value("id").(int64)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}
	user := &model.User{ID: id}
	op := operator.NewUserOperator(ctx, user)
	if err := op.Load(ctx); err != nil && err != gorm.ErrRecordNotFound {
		ctx.Error(err)
		return
	}
	ctx.Writer().Write([]byte(fmt.Sprintf("Hello, %s!", user.Username)))
}
