package schedulehandler

import (
	schedulemodel "e-ticket/internal/domain/schedule/model"
	scheduleservice "e-ticket/internal/domain/schedule/service"
	appresponse "e-ticket/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *scheduleservice.Service
}

func NewScheduleHandler(service scheduleservice.Service) *Handler {
	return &Handler{service: &service}
}

func (h *Handler) GetSchedules(ctx *gin.Context) {

	var uri schedulemodel.ScheduleUri

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	schedules, err := h.service.GetSchedulesById(uri.CompanyId)

	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}
	ctx.JSON(appresponse.Success(http.StatusOK, schedules))
}

func (h *Handler) CreateSchedules(ctx *gin.Context) {

	var uri schedulemodel.ScheduleUri
	var request schedulemodel.ScheduleCreateEntity

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

	isCerated, err := h.service.CreateSchedules(uri.CompanyId, request)

	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}
	ctx.JSON(appresponse.Success(http.StatusOK, isCerated))

}
