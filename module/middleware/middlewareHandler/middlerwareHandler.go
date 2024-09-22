package middlewareHandler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kritAsawaniramol/myAnimeList-server/config"
	"github.com/kritAsawaniramol/myAnimeList-server/module/middleware/middlewareUsecase"
)

type (
	MiddlewareHandlerService interface {
		JwtAuthorization() gin.HandlerFunc
	}

	middlewareHandler struct {
		middlewareUsecase middlewareUsecase.MiddlewareUsecaseService
		cfg               *config.Config
	}
)

// JwtAuthorization implements MiddlewareHandlerService.
func (m *middlewareHandler) JwtAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		accessToken := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
		log.Printf("accessToken: %v\n", accessToken)

		userID, err := m.middlewareUsecase.JwtAuthorization(m.cfg, accessToken)
		if err != nil {
			log.Printf("error: JwtAuthorization: %s\n", err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized access",
			})
			return
		}

		ctx.Set("userID", userID)
		ctx.Next()
	}
}

func NewMiddlewareHandler(cfg *config.Config, middlewareUsecase middlewareUsecase.MiddlewareUsecaseService) MiddlewareHandlerService {
	return &middlewareHandler{
		cfg:               cfg,
		middlewareUsecase: middlewareUsecase,
	}
}

func GetHeaders(c *gin.Context) {
	headers := c.Request.Header
	for name, values := range headers {
		for _, value := range values {
			fmt.Printf("Header: %s, Value: %s\n", name, value)
		}
	}
}
