package middleware

import (
	"e-ticket/internal/config"
	appresponse "e-ticket/pkg/response"
	apptoken "e-ticket/pkg/token"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthToken(ctx *gin.Context) {
	type AuthHeaders struct {
		Authorization string `header:"Authorization" binding:"required"`
		// UserAgent     string `header:"User-Agent"`
		// ContentType   string `header:"Content-Type" binding:"required,contains=application/json"`
	}
	var authHeaders AuthHeaders
	err := ctx.ShouldBindHeader(&authHeaders)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	if !strings.HasPrefix(authHeaders.Authorization, "Bearer ") {
		err := errors.New("bearer is required")
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusUnauthorized, err))
		return
	}

	// Load config model from env
	appConfigModel, _ := config.Load()

	isValid := apptoken.IsTokenValid(authHeaders.Authorization[7:], appConfigModel.Keys.JwtSecretKey)
	if !isValid {
		err := errors.New("invalid token")
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusUnauthorized, err))
		return
	}

	ctx.Next()
}
