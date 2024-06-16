package response

import "github.com/gin-gonic/gin"

type Status struct {
	Code      int  `json:"code"`
	IsSuccess bool `json:"isSuccess"`
}

type Response struct {
	Status  Status      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(code, Response{
		Status: Status{
			Code:      code,
			IsSuccess: true,
		},
		Message: message,
		Data:    data,
	})
}

func Error(ctx *gin.Context, code int, message string, err error) {
	ctx.JSON(code, Response{
		Status: Status{
			Code:      code,
			IsSuccess: false,
		},
		Message: message,
		Data:    err.Error(),
	})
}
