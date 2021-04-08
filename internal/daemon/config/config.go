package config

import (
	"flag"
)

type Config struct {
	BindPort uint
	Routers  string
}

func Init(c *Config) {
	flag.UintVar(&c.BindPort, "port", 5344, "Bind port for daemon api")
	flag.StringVar(&c.Routers, "routers", "", "Specify routers to connect to")
	flag.Parse()
}
