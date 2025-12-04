package middleware

import (
	"e-ticket/internal/config"
	appresponse "e-ticket/pkg/response"
	apptoken "e-ticket/pkg/token"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func _isExpire(dt time.Time, expireHour int) bool {
	tem := dt.Add(time.Hour * time.Duration(expireHour))
	return tem.Before(time.Now())
}

func AuthToken(ctx *gin.Context) {
	type AuthHeaders struct {
		Authorization string `header:"Authorization" binding:"required"`
	}

	var authHeaders AuthHeaders
	err := ctx.ShouldBindHeader(&authHeaders)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	if !strings.HasPrefix(authHeaders.Authorization, "Bearer ") {
		err := errors.New("bearer is required")
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	// Load config model from env
	appConfigModel, _ := config.Load()
	tokeSubstring := authHeaders.Authorization[7:]

	isValid := apptoken.IsTokenValid(tokeSubstring, appConfigModel.Keys.JwtSecretKey)
	if !isValid {
		err := errors.New("invalid token")
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusUnauthorized, err))
		return
	}

	decodedData, _ := apptoken.Decript(tokeSubstring, appConfigModel.Keys.JwtSecretKey)

	byteData, _ := json.Marshal(decodedData)
	var tokenEntity TokenEnitty
	json.Unmarshal(byteData, &tokenEntity)

	isExpire := _isExpire(tokenEntity.Time, appConfigModel.Keys.JwtExpiryHour)
	if isExpire {
		err := errors.New("token expired")
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusUnauthorized, err))
		return
	}

	ctx.Next()
}
