package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/cro4k/authorize/internal/service"
)

func OAuth2Authorize(c *gin.Context) {
	err := service.OAuth2.HandleAuthorizeRequest(c.Writer, c.Request)
	if err != nil {
		logrus.WithContext(c).Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func OAuth2Token(c *gin.Context) {
	err := service.OAuth2.HandleTokenRequest(c.Writer, c.Request)
	if err != nil {
		logrus.WithContext(c).Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
	}
}
