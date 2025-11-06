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

	var query vehiclemodel.VehicleQuery

	err := ctx.ShouldBindQuery(&query)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	list, err := h.service.GetAllVehicle(query.CompanyId, query.UserId)

	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}
	ctx.JSON(appresponse.Success(http.StatusOK, list))
}

func (h *Handler) CreateVehicle(ctx *gin.Context) {

	var vehicle vehiclemodel.VehicleCreateRequest

	err := ctx.ShouldBindBodyWithJSON(&vehicle)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	list, err := h.service.CreateVehicle(vehicle)

	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}
	ctx.JSON(appresponse.Success(http.StatusCreated, list))
}
