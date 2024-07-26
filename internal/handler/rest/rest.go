package rest

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yogarn/arten/internal/handler/websocket"
	"github.com/yogarn/arten/internal/service"
	"github.com/yogarn/arten/pkg/response"
)

type Rest struct {
	router    *gin.Engine
	service   *service.Service
	wsManager *websocket.WebSocketManager
}

func NewRest(service *service.Service, wsManager *websocket.WebSocketManager) *Rest {
	return &Rest{
		router:    gin.Default(),
		service:   service,
		wsManager: wsManager,
	}
}

func MountTranslation(routerGroup *gin.RouterGroup, rest *Rest) {
	translation := routerGroup.Group("/translations")
	translation.POST("/", rest.CreateTranslation)
	translation.GET("/:id", rest.GetTranslation)
	translation.PUT("/:id", rest.UpdateTranslation)
	translation.DELETE("/:id", rest.DeleteTranslation)
}

func MountWebsocket(routerGroup *gin.RouterGroup, rest *Rest) {
	websocketGroup := routerGroup.Group("/websocket")
	websocketGroup.GET("/:chatId", func(c *gin.Context) {
		chatId, err := uuid.Parse(c.Param("chatId"))
		if err != nil {
			response.Error(c, http.StatusBadRequest, "invalid chat ID", err)
			return
		}
		rest.wsManager.HandleConnections(chatId, c.Writer, c.Request)
	})
}

func (rest *Rest) MountEndpoints() {
	rest.router.NoRoute(func(ctx *gin.Context) {
		response.Error(ctx, http.StatusNotFound, "not found", errors.New("page not found"))
	})

	routerGroup := rest.router.Group("/api/v1")
	MountTranslation(routerGroup, rest)
	MountWebsocket(routerGroup, rest)
}

func (rest *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	rest.router.Run(fmt.Sprintf("%s:%s", addr, port))
}
