package middleware

import (
	"errors"
	"net/http"
	"promotion/configs"
	"promotion/pkg/response"

	"github.com/gin-gonic/gin"
)

const internalAuthHeader = "x-api-key"

type InternalAuthMiddleware struct {
	cfg *configs.Config
}

func NewInternalAuthMiddleware(cfg *configs.Config) *InternalAuthMiddleware {
	return &InternalAuthMiddleware{cfg: cfg}
}

func (m *InternalAuthMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if errHeader := verifyAuthHeader(ctx, m.cfg); errHeader == nil {
			ctx.Next()
			return
		}

		errQuery := verifyAuthTokenFromQuery(ctx, m.cfg)
		if errQuery == nil {
			ctx.Next()
			return
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.HTTPResponse{
			Status: http.StatusUnauthorized,
			Data:   errQuery.Error(),
		})
		return
	}
}

func verifyAuthHeader(ctx *gin.Context, cfg *configs.Config) error {
	authHeader := ctx.GetHeader(internalAuthHeader)
	if authHeader == "" {
		return errors.New("Missing auth header")
	}

	if authHeader != cfg.APIKey.PromotionAPIKey {
		return errors.New("Invalid auth")
	}
	return nil
}

func verifyAuthTokenFromQuery(ctx *gin.Context, cfg *configs.Config) error {
	// Mostly for the use of PubSub
	authToken, ok := ctx.GetQuery(internalAuthHeader)
	if !ok {
		return errors.New("Missing auth header")
	}
	if authToken != cfg.APIKey.PromotionAPIKey {
		return errors.New("Invalid auth")
	}
	return nil
}
