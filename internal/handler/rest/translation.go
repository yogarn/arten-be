package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yogarn/arten/entity"
	"github.com/yogarn/arten/pkg/response"
)

func (rest *Rest) CreateTranslation(ctx *gin.Context) {
	translation := &entity.Translation{}
	if err := ctx.BindJSON(translation); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid request", err)
		return
	}

	if err := rest.service.TranslationService.CreateTranslation(ctx, translation); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create translation", err)
		return
	}

	chatId := translation.ChatId
	rest.wsManager.Broadcast(chatId, []byte(translation.Translate))

	response.Success(ctx, http.StatusCreated, "Translation created", translation)
}

func (rest *Rest) GetTranslation(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid ID", err)
		return
	}

	translation, err := rest.service.TranslationService.GetTranslation(ctx, id)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Translation not found", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Translation found", translation)
}

func (rest *Rest) UpdateTranslation(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid ID", err)
		return
	}

	translation := &entity.Translation{}
	if err := ctx.BindJSON(translation); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid request", err)
		return
	}

	if err := rest.service.TranslationService.UpdateTranslation(ctx, id, translation); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update translation", err)
		return
	}

	chatId := translation.ChatId
	rest.wsManager.Broadcast(chatId, []byte(translation.Translate))

	response.Success(ctx, http.StatusOK, "Translation updated", translation)
}

func (rest *Rest) DeleteTranslation(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid ID", err)
		return
	}

	if err := rest.service.TranslationService.DeleteTranslation(ctx, id); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to delete translation", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Translation deleted", nil)
}
