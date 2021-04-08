package main

import (
	"github.com/bjartur20/GoBlockchain/internal/daemon/config"
	"github.com/bjartur20/GoBlockchain/internal/daemon/discovery"
	"github.com/bjartur20/GoBlockchain/internal/daemon/routes"
)

func main() {
	// Parse cli arguments and fill in the config struct.
	var c config.Config
	config.Init(&c)

	// Start the DHT node.
	go discovery.Run(c.Routers)

	// Start the router.
	routes.Run(c)
}
