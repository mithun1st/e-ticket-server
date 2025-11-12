package subuserhandler

import (
	subusermodel "e-ticket/internal/domain/sub_user/model"
	subuserservice "e-ticket/internal/domain/sub_user/service"
	appresponse "e-ticket/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *subuserservice.Service
}

func NewSubUserHandler(service subuserservice.Service) *Handler {
	return &Handler{service: &service}
}

func (h *Handler) GetAllSubUser(ctx *gin.Context) {
	var uri subusermodel.SubUserUri
	var query subusermodel.SubUserQuery

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	err = ctx.ShouldBindQuery(&query)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	list, err := h.service.GetAllSubUser(uri.CompanyId, query.Role)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusUnauthorized, err))
		return
	}
	ctx.JSON(appresponse.Success(http.StatusOK, list))
}

func (h *Handler) CreateSubUser(ctx *gin.Context) {
	var uri subusermodel.SubUserUri
	var request subusermodel.SubUserCreateRequest

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	err = ctx.ShouldBindBodyWithJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	isCreated, err := h.service.CreateSubUser(uri.CompanyId, request)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusUnauthorized, err))
		return
	}
	ctx.JSON(appresponse.Success(http.StatusCreated, isCreated))
}
