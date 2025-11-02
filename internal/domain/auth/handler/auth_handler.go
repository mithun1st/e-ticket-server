package authhandler

import (
	"e-ticket/internal/config"
	authmodel "e-ticket/internal/domain/auth/model"
	authservice "e-ticket/internal/domain/auth/service"
	appresponse "e-ticket/pkg/response"
	apptoken "e-ticket/pkg/token"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *authservice.Service
}

func NewAuthHandler(service authservice.Service) *Handler {
	return &Handler{service: &service}
}

func (h *Handler) GetAuthCompany(ctx *gin.Context) {

	var request authmodel.AuthRequest

	err := ctx.ShouldBindBodyWithJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	if request.Email == "" && request.Phone == "" {
		var err error = errors.New("email or phone required")
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	auth, err := h.service.GetAuthCompany(request.Email, request.Phone, request.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusUnauthorized, err))
		return
	}

	// Load config model from env
	appConfigModel, _ := config.Load()

	token, err := apptoken.Encript(map[string]any{
		"id":    auth.UserEntity.Id,
		"email": auth.UserEntity.Email,
		"phone": auth.UserEntity.Phone,
		"time":  time.Now(),
	}, appConfigModel.Keys.JwtSecretKey)

	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
	}
	auth.Token = "Bearer " + token

	ctx.JSON(appresponse.Success(http.StatusOK, auth))
}

func (h *Handler) GetAuthVehicle(ctx *gin.Context) {

	var query authmodel.AuthQuery
	var request authmodel.AuthRequest

	err := ctx.ShouldBindQuery(&query)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	err = ctx.ShouldBindBodyWithJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	if request.Email == "" && request.Phone == "" || query.CompanyId == 0 {
		var err error = errors.New("companyId and email or phone required")
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	auth, err := h.service.GetAuthSubUser(query.CompanyId, request.Email, request.Phone, request.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusUnauthorized, err))
		return
	}

	// Load config model from env
	appConfigModel, _ := config.Load()

	token, err := apptoken.Encript(map[string]any{
		"id":    auth.UserEntity.Id,
		"email": auth.UserEntity.Email,
		"phone": auth.UserEntity.Phone,
		"time":  time.Now(),
	}, appConfigModel.Keys.JwtSecretKey)

	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
	}
	auth.Token = "Bearer " + token

	ctx.JSON(appresponse.Success(http.StatusOK, auth))
}

// func (h *Handler) GetAllAuth(ctx *gin.Context) {

// 	list, err := h.service.GetAllAuth()

// 	if err != nil {
// 		ctx.AbortWithStatusJSON(response.Error(http.StatusUnauthorized, err))
// 		return
// 	}
// 	ctx.JSON(response.Success(http.StatusOK, list))
// }
