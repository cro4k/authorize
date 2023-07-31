package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cro4k/authorize/server/api/controller"
)

func Router(e gin.IRouter, static ...bool) {
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

	manage := e.Group("/api/manage", ManageAuth)
	manage.POST("/user/list", controller.UserList)
	manage.POST("/user/create", controller.CreateUser)
	manage.POST("/user/edit", controller.EditUser)
	manage.POST("/user/status", controller.SetUserStatus)
	manage.POST("/app/list", controller.AppList)
	manage.POST("/app/create", controller.CreateApp)
	manage.POST("/app/edit", controller.EditApp)
	manage.POST("/app/status", controller.SetAppStatus)

	if len(static) > 0 && static[0] {
		e.GET("/v/", func(ctx *gin.Context) {
			http.StripPrefix("/v/", http.FileServer(http.Dir("static/"))).ServeHTTP(ctx.Writer, ctx.Request)
		})
	}
}
