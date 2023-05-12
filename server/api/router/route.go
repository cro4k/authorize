package router

import (
	"github.com/gin-gonic/gin"

	"github.com/cro4k/authorize/server/api/controller"
)

func Router(e gin.IRouter) {
	e.Use(UUID, Logger, Token)

	auth := e.Group("/api/auth")
	auth.POST("/login", controller.Login)
	auth.POST("/register", controller.Register)
	auth.POST("/logout", controller.Logout)

	oauth2 := e.Group("/api/oauth2", Auth)
	oauth2.GET("/authorize", controller.OAuth2Authorize)
	oauth2.POST("/authorize", controller.OAuth2Authorize)
	oauth2.GET("/token", controller.OAuth2Token)
	oauth2.POST("/token", controller.OAuth2Token)

}
