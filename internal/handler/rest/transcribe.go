package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yogarn/arten/pkg/response"
)

func (rest *Rest) EnglishTranscribe(ctx *gin.Context) {
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid request", err)
		return
	}

	resp, err := rest.service.TranscribeService.TranscribeEnglish(ctx, file)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid request", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success", resp)
}

func (rest *Rest) IndonesianTranscribe(ctx *gin.Context) {
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid request", err)
		return
	}

	resp, err := rest.service.TranscribeService.TranscribeIndonesian(ctx, file)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid request", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success", resp)
}
