package response

import (
	"net/http"
	"promotion/pkg/failure"

	"github.com/gin-gonic/gin"
)

const DefaultResponseSuccess = "success"

type HTTPResponse struct {
	Status failure.ErrorCode `json:"status"`
	Data   any               `json:"data"`
}

func ErrorBinding(ctx *gin.Context, err *failure.BindJSONError) {
	_ = ctx.Error(err)
}

func ErrorApp(ctx *gin.Context, err *failure.AppError) {
	_ = ctx.Error(err)
}

func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, HTTPResponse{
		Status: http.StatusOK,
		Data:   data,
	})
}
