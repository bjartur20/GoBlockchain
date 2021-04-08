package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/bjartur20/GoBlockchain/internal/daemon/config"
)

var (
	r = gin.Default()
)

func getRoutes() {
	v1 := r.Group("/v1")
	addPingRoutes(v1)
}

// Start the API server.
func Run(c config.Config) {
	getRoutes()
	r.Run(fmt.Sprintf(":%d", c.BindPort))
}
