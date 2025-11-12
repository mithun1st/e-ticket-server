package vehiclehandler

import (
	vehiclemodel "e-ticket/internal/domain/vehicle/model"
	vehicleservice "e-ticket/internal/domain/vehicle/service"
	appresponse "e-ticket/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *vehicleservice.Service
}

func NewVehicleHandler(service vehicleservice.Service) *Handler {
	return &Handler{service: &service}
}

func (h *Handler) GetAllVehicle(ctx *gin.Context) {

	var uri vehiclemodel.VehicleUri
	var query vehiclemodel.VehicleQuery

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

	list, err := h.service.GetAllVehicle(uri.CompanyId, query.UserId)

	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}
	ctx.JSON(appresponse.Success(http.StatusOK, list))
}

func (h *Handler) CreateVehicle(ctx *gin.Context) {

	var uri vehiclemodel.VehicleUri
	var vehicle vehiclemodel.VehicleCreateRequest

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	err = ctx.ShouldBindBodyWithJSON(&vehicle)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	list, err := h.service.CreateVehicle(uri.CompanyId, vehicle)

	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}
	ctx.JSON(appresponse.Success(http.StatusCreated, list))
}
