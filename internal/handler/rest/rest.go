package rest

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yogarn/arten/internal/service"
	"github.com/yogarn/arten/pkg/response"
)

type Rest struct {
	router  *gin.Engine
	service *service.Service
}

func NewRest(service *service.Service) *Rest {
	return &Rest{
		router:  gin.Default(),
		service: service,
	}
}

func MountTranslation(routerGroup *gin.RouterGroup, rest *Rest) {
	translation := routerGroup.Group("/translations")
	translation.POST("/", rest.CreateTranslation)
	translation.GET("/:id", rest.GetTranslation)
	translation.PUT("/:id", rest.UpdateTranslation)
	translation.DELETE("/:id", rest.DeleteTranslation)
}

func (rest *Rest) MountEndpoints() {
	rest.router.NoRoute(func(ctx *gin.Context) {
		response.Error(ctx, http.StatusNotFound, "not found", errors.New("page not found"))
	})

	routerGroup := rest.router.Group("/api/v1")
	MountTranslation(routerGroup, rest)
}

func (rest *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	rest.router.Run(fmt.Sprintf("%s:%s", addr, port))
}
