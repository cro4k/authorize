package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/cro4k/authorize/internal/module"
	"net/http"
	"strings"
)

func Authorize(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if len(token) > 7 && strings.ToLower(token[:7]) == "bearer " {
		token = token[7:]
	}
	id, err := module.GetAuthService().Validate(ctx.Request.Context(), token)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.Set("id", id)
}
