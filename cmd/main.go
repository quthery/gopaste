package main

import (
	"gopaste/internal/config"
	"gopaste/internal/gorm/connection"
	"gopaste/internal/gorm/models"
	"gopaste/internal/http-server/router"
	"gopaste/internal/http-server/server"
	"gopaste/internal/service/user"
)

func main() {
	cfg := config.MustLoad("config/config.yaml")
	db := connection.MySQLConn(cfg)
	db.Migrate(&models.User{}, &models.CodeNote{})
	userService := user.New(db.DB)
	r := router.New(userService)
	panic(server.Run(cfg, r))
}
