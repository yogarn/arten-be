package main

import (
	"github.com/yogarn/arten/internal/handler/rest"
	"github.com/yogarn/arten/internal/handler/websocket"
	"github.com/yogarn/arten/internal/repository"
	"github.com/yogarn/arten/internal/service"
	"github.com/yogarn/arten/pkg/config"
	"github.com/yogarn/arten/pkg/database/mysql"
)

func main() {
	config.LoadEnvironment()
	db := mysql.ConnectDatabase()
	defer db.Close()

	wsManager := websocket.NewWebSocketManager()

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	rest := rest.NewRest(service, wsManager)

	rest.MountEndpoints()
	rest.Run()
}
