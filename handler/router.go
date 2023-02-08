package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetApiRoutes(e *gin.Engine) *gin.Engine {
	v1 := e.Group("/v1")
	{
		v1.GET("/", HealthCheck)
	}
	return e
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
