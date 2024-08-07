package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yogarn/arten/internal/service"
	"github.com/yogarn/arten/pkg/jwt"
)

type Interface interface {
	AuthenticateUser(ctx *gin.Context)
	CheckTranslationOwnership(ctx *gin.Context)
	Logger() gin.HandlerFunc
}

type middleware struct {
	jwtAuth jwt.Interface
	service *service.Service
}

func Init(jwtAuth jwt.Interface, service *service.Service) Interface {
	return &middleware{
		jwtAuth: jwtAuth,
		service: service,
	}
}
