package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserList(ctx *gin.Context) {
	ctx.AbortWithError(http.StatusInternalServerError, errors.New("not implemented yet"))
}
func CreateUser(ctx *gin.Context) {
	ctx.AbortWithError(http.StatusInternalServerError, errors.New("not implemented yet"))
}
func EditUser(ctx *gin.Context) {
	ctx.AbortWithError(http.StatusInternalServerError, errors.New("not implemented yet"))
}
func SetUserStatus(ctx *gin.Context) {
	ctx.AbortWithError(http.StatusInternalServerError, errors.New("not implemented yet"))
}

func AppList(ctx *gin.Context) {
	ctx.AbortWithError(http.StatusInternalServerError, errors.New("not implemented yet"))
}
func CreateApp(ctx *gin.Context) {
	ctx.AbortWithError(http.StatusInternalServerError, errors.New("not implemented yet"))
}
func EditApp(ctx *gin.Context) {
	ctx.AbortWithError(http.StatusInternalServerError, errors.New("not implemented yet"))
}
func SetAppStatus(ctx *gin.Context) {
	ctx.AbortWithError(http.StatusInternalServerError, errors.New("not implemented yet"))
}
