package main

import (
	"github.com/yogarn/arten/internal/handler/rest"
	"github.com/yogarn/arten/internal/handler/websocket"
	"github.com/yogarn/arten/internal/repository"
	"github.com/yogarn/arten/internal/service"
	"github.com/yogarn/arten/pkg/bcrypt"
	"github.com/yogarn/arten/pkg/config"
	"github.com/yogarn/arten/pkg/database/mysql"
	"github.com/yogarn/arten/pkg/jwt"
	"github.com/yogarn/arten/pkg/middleware"
)

func main() {
	config.LoadEnvironment()
	db := mysql.ConnectDatabase()
	defer db.Close()

	wsManager := websocket.NewWebSocketManager()

	jwt := jwt.Init()
	bcrypt := bcrypt.Init()

	repository := repository.NewRepository(db)
	service := service.NewService(repository, bcrypt, jwt)
	middleware := middleware.Init(jwt, service)
	rest := rest.NewRest(service, wsManager, middleware, jwt, bcrypt)

	rest.MountEndpoints()
	rest.Run()
}
