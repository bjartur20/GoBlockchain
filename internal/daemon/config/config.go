package config

import (
	"flag"
)

type Config struct {
	BindPort uint
}

func Init(c *Config) {
	flag.UintVar(&c.BindPort, "port", 5344, "Bind port for daemon api")
	flag.Parse()
}
