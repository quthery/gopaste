package main

import (
	"gopaste/internal/config"
	"gopaste/internal/http-server/router"
	"gopaste/internal/http-server/server"
)

func main() {
	cfg := config.MustLoad("config/config.yaml")
	r := router.New()
	panic(server.Run(cfg, r))
}
