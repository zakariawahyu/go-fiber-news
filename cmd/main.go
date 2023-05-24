package main

import (
	"github.com/zakariawahyu/go-fiber-news/cmd/server"
	"github.com/zakariawahyu/go-fiber-news/config"
	"github.com/zakariawahyu/go-fiber-news/utils"
)

func main() {
	cfg := config.NewConfig()
	db := utils.NewDbConnection(cfg)

	repo := server.NewRepository(db)
	serv := server.NewServices(repo)
	server.NewHandler(cfg, serv)
}
