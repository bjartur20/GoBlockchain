package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addPingRoutes(rg *gin.RouterGroup) {
	g := rg.Group("/ping")

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}
