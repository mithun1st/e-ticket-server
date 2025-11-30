package routecounterhandler

import (
	routecountermodel "e-ticket/internal/domain/route_counter/model"
	routecounterservice "e-ticket/internal/domain/route_counter/service"
	appresponse "e-ticket/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *routecounterservice.Service
}

func NewRouteCounterHandler(service routecounterservice.Service) *Handler {
	return &Handler{service: &service}
}

func (h *Handler) GetCounters(ctx *gin.Context) {

	var uri routecountermodel.RouteCounterUri

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	counters, err := h.service.GetCounters(uri.CompanyId)

	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}
	ctx.JSON(appresponse.Success(http.StatusOK, counters))
}

func (h *Handler) CreateCounter(ctx *gin.Context) {

	var uri routecountermodel.RouteCounterUri
	var counter routecountermodel.CounterCreateEntity

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	err = ctx.ShouldBindBodyWithJSON(&counter)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	isCreated, err := h.service.CreateCounters(uri.CompanyId, counter)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}
	ctx.JSON(appresponse.Success(http.StatusCreated, isCreated))
}

func (h *Handler) GetRoutes(ctx *gin.Context) {

	var uri routecountermodel.RouteCounterUri

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	routes, err := h.service.GetRoutes(uri.CompanyId)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}
	ctx.JSON(appresponse.Success(http.StatusOK, routes))
}

func (h *Handler) CreateRoute(ctx *gin.Context) {

	var uri routecountermodel.RouteCounterUri
	var route routecountermodel.RouteCreateEntity

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	err = ctx.ShouldBindBodyWithJSON(&route)
	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}

	isCreated, err := h.service.CreateRoute(uri.CompanyId, route)

	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}
	ctx.JSON(appresponse.Success(http.StatusCreated, isCreated))
}
