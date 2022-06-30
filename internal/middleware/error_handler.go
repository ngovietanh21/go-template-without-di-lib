package middleware

import (
	"net/http"
	"promotion/pkg/failure"
	"promotion/pkg/logger"
	"promotion/pkg/response"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(log *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		for _, err := range ctx.Errors {
			handleContextErr(ctx, log, err)
		}
	}
}

func handleContextErr(ctx *gin.Context, log *logger.Logger, err *gin.Error) {
	if appError, ok := err.Err.(*failure.BindJSONError); ok {
		if appError.OriginalErr != nil {
			log.Error(appError.OriginalErr)
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.HTTPResponse{
			Status: appError.Code,
			Data:   appError.Error(),
		})
		return
	}

	if appError, ok := err.Err.(*failure.AppError); ok {
		if appError.OriginalErr != nil {
			log.Error(appError.OriginalErr)
		}
		ctx.AbortWithStatusJSON(appError.HTTPCode(), response.HTTPResponse{
			Status: appError.Code,
			Data:   appError.Error(),
		})
		return
	}

	log.Error(err.Err)
	ctx.JSON(http.StatusBadRequest, response.HTTPResponse{
		Status: http.StatusBadRequest,
		Data:   err.Error(),
	})
}
