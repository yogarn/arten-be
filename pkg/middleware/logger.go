package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (m *middleware) Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		latency := time.Since(start)
		statusCode := ctx.Writer.Status()

		logger := log.Info()
		if statusCode >= 400 {
			logger = log.Error()
		}

		logger.Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Int("status", statusCode).
			Dur("latency", latency).
			Str("client_ip", ctx.ClientIP()).
			Str("user_agent", ctx.Request.UserAgent()).
			Msg("HTTP request")
	}
}
