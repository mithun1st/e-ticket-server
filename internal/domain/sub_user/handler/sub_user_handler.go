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

// func (h *Handler) GetSubUser(ctx *gin.Context) {

// var header subusermodel.SubUserHeader
// var path subusermodel.SubUserPath
// 	var query subusermodel.SubUserQuery
// 	var request subusermodel.SubUserRequest

// 	err := ctx.ShouldBindHeader(&header)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
// 		return
// 	}

// 	err = ctx.ShouldBindUri(&path)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
// 		return
// 	}

// 	err = ctx.ShouldBindQuery(&query)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
// 		return
// 	}

// 	err = ctx.ShouldBindBodyWithJSON(&request)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
// 		return
// 	}

// 	subUser, err := h.service.GetSubUser(request.Id)

// 	if err != nil {
// 		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusUnauthorized, err))
// 		return
// 	}
// 	ctx.JSON(appresponse.Success(http.StatusOK, map[string]any{
// 		"header":  header,
// 		"query":   query,
// 		"request": request,
// 		"result":  subUser,
// 	}))
// }

func (h *Handler) GetAllSubUser(ctx *gin.Context) {
	var query subusermodel.SubUserQuery

	err := ctx.ShouldBindQuery(&query)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	list, err := h.service.GetAllSubUser(query.CompanyId, *query.Role)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusUnauthorized, err))
		return
	}
	ctx.JSON(appresponse.Success(http.StatusOK, list))
}

func (h *Handler) CreateSubUser(ctx *gin.Context) {
	var request subusermodel.SubUserCreateRequest

	err := ctx.ShouldBindBodyWithJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	isCreated, err := h.service.CreateSubUser(request)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusUnauthorized, err))
		return
	}
	ctx.JSON(appresponse.Success(http.StatusCreated, isCreated))
}
