package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func addPingRoutes(rg *gin.RouterGroup) {
	g := rg.Group("/ping")

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}

func addSubmitRoutes(rg *gin.RouterGroup) {
	g := rg.Group("/documents")

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "you submitted!")
	})
}

func getRoutes() {
	v1 := r.Group("/v1")
	addPingRoutes(v1)
}

func Run() {
	getRoutes()
	r.Run(fmt.Sprintf(":6060"))
}

func main() {
	Run()
}
