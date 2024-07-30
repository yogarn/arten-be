package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yogarn/arten/pkg/response"
)

func (m *middleware) CheckTranslationOwnership(ctx *gin.Context) {
	userId, err := m.jwtAuth.GetLoginUser(ctx)
	if err != nil {
		response.Error((ctx), 401, "Unauthorized", err)
		ctx.Abort()
		return
	}

	translationId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		response.Error(ctx, 400, "Invalid translation ID", err)
		ctx.Abort()
		return
	}

	translation, err := m.service.TranslationService.GetTranslation(ctx, translationId)
	if err != nil {
		response.Error(ctx, 500, "Failed to get translation", err)
		ctx.Abort()
		return
	}

	if translation.UserId != userId {
		response.Error(ctx, 403, "Forbidden", errors.New("you are not the owner of this translation"))
		ctx.Abort()
		return
	}
	ctx.Next()
}
