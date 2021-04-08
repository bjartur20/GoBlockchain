package main

import (
	"fmt"

	"github.com/bjartur20/GoBlockchain/internal/daemon/config"
	"github.com/bjartur20/GoBlockchain/internal/daemon/discovery"
	"github.com/bjartur20/GoBlockchain/internal/daemon/names"
	"github.com/bjartur20/GoBlockchain/internal/daemon/routes"
)

func main() {
	// Parse cli arguments and fill in the config struct.
	var c config.Config
	config.Init(&c)

	// Start the naming service
	n := names.Make()
	fmt.Printf("%+v\n", n)

	// Check for routers
	var routers string
	if c.Routers != 0 {
		routers = fmt.Sprintf("localhost:%d", c.Routers)
	}
	// Start the DHT node.
	go discovery.Run(routers)

	// Start the router.
	routes.Run(c)
}
