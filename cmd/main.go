package main

import (
	"github.com/aveplen-bach/config-service/internal/config"
	"github.com/aveplen-bach/config-service/internal/server"
)

func main() {
	var cfg config.Config
	config.ReadConfig("config-service.yaml", &cfg)
	server.Start(cfg)
}
