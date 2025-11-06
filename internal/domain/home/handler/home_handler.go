package homehandler

import (
	homemodel "e-ticket/internal/domain/home/model"
	homeservice "e-ticket/internal/domain/home/service"
	appresponse "e-ticket/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *homeservice.Service
}

func NewHomeHandler(service homeservice.Service) *Handler {
	return &Handler{service: &service}
}

func (h *Handler) GetCompanyHome(ctx *gin.Context) {

	var query homemodel.HomeQuery

	err := ctx.ShouldBindQuery(&query)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	home, err := h.service.GetCompanyHomeData(query.UserId)

	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}
	ctx.JSON(appresponse.Success(http.StatusOK, home))
}

func (h *Handler) GetOwnerHome(ctx *gin.Context) {

	var query homemodel.HomeQuery

	err := ctx.ShouldBindQuery(&query)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	home, err := h.service.GetOwnerHomeData(query.UserId)

	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}
	ctx.JSON(appresponse.Success(http.StatusOK, home))
}
