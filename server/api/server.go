package api

import (
	"github.com/gin-gonic/gin"

	"github.com/cro4k/authorize/common/services/api"
	"github.com/cro4k/authorize/doc/temporary"
	"github.com/cro4k/authorize/server/api/router"
)

func NewServer(c api.Config) *api.GINServer {
	srv := api.NewGINServer(c)
	srv.RegisterRouter(func(e *gin.Engine) {
		router.Router(e)
		temporary.Set(e)
	})
	return srv
}
