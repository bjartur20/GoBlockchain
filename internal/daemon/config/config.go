package config

import (
	"flag"
)

type Config struct {
	BindPort uint
	Routers  uint
}

func Init(c *Config) {
	flag.UintVar(&c.BindPort, "port", 5344, "Bind port for daemon api")
	flag.UintVar(&c.Routers, "routers", 0, "Specify routers for new node")
	flag.Parse()
}
