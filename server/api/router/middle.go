package router

import (
	"fmt"
	"github.com/cro4k/authorize/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

func UUID(c *gin.Context) {
	rid := uuid.NewString()
	cid := c.GetHeader("client_id")
	if cid == "" {
		cid, _ = c.Cookie("client_id")
	}
	if cid == "" {
		cid = uuid.NewString()
		c.SetCookie("client_id", cid, 86400*100, "/", "", false, true)
	}
	c.Set("rid", rid)
	c.Set("cid", cid)
}

func Logger(ctx *gin.Context) {
	start := time.Now()
	ctx.Next()
	info := fmt.Sprintf("| %3d | %-10v | %-16s |%6s | %s",
		ctx.Writer.Status(),
		time.Since(start),
		ctx.RemoteIP(),
		ctx.Request.Method,
		ctx.Request.RequestURI,
	)
	logrus.WithContext(ctx).Info(info)
}

func Token(ctx *gin.Context) {
	tokenStr := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
	if tokenStr == "" {
		tokenStr, _ = ctx.Cookie("token")
	}
	claims, err := service.Auth.Verify(tokenStr, ctx.GetString("cid"))
	ctx.Set("token_error", err)
	if err == nil {
		ctx.Set("claims", claims)
		ctx.Set("uid", claims.UID)
	}
}

func Auth(ctx *gin.Context) {
	err, _ := ctx.Get("token_error")
	id := ctx.GetString("uid")
	if err != nil || id == "" {
		ctx.AbortWithStatus(http.StatusForbidden)
	}
}
