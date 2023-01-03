package temporary

import "github.com/gin-gonic/gin"

var engine *gin.Engine

func Set(e *gin.Engine) {
	engine = e
}

func Get() *gin.Engine {
	return engine
}
