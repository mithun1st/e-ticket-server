package authhandler

import (
	authmodel "e-ticket/internal/domain/auth/model"
	authservice "e-ticket/internal/domain/auth/service"
	appresponse "e-ticket/pkg/response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *authservice.Service
}

func NewAuthHandler(service authservice.Service) *Handler {
	return &Handler{service: &service}
}

func (h *Handler) GetAuthCompany(ctx *gin.Context) {

	var request authmodel.AuthSubUserRequest

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

	ctx.JSON(appresponse.Success(http.StatusOK, auth))
}

func (h *Handler) GetAuthVehicle(ctx *gin.Context) {

	var request authmodel.AuthCompanyRequest

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

	auth, err := h.service.GetAuthSubUser(request.CompanyId, request.Email, request.Phone, request.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusUnauthorized, err))
		return
	}

	ctx.JSON(appresponse.Success(http.StatusOK, auth))
}
