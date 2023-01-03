package router

import (
	"github.com/cro4k/authorize/doc/temporary"
	"github.com/cro4k/authorize/server/api/controller"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	e := gin.New()
	e.Use(UUID, Logger, Token)

	auth := e.Group("/api/auth")
	auth.POST("/login", controller.Login)
	auth.POST("/register", controller.Register)
	auth.POST("/logout", controller.Logout)

	temporary.Set(e)

	return e
}
