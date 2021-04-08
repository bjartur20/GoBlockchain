package main

import (
	"fmt"
	"github.com/bjartur20/GoBlockchain/internal/daemon/discovery"
	"github.com/bjartur20/GoBlockchain/internal/daemon/routes"
	"github.com/bjartur20/GoBlockchain/internal/daemon/config"
	"github.com/bjartur20/GoBlockchain/internal/daemon/names"
)

func main() {
	// Parse cli arguments and fill in the config struct.
	var c config.Config
	config.Init(&c);

	// Start the naming service
	n := names.Make()
	fmt.Printf("%+v\n", n)

	// Start the DHT node.
	go discovery.Run("")

	// Start the router.
	routes.Run(c)
}
