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

	counters, err := h.service.CreateCounters(uri.CompanyId, counter)

	if err != nil {
		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
		return
	}
	ctx.JSON(appresponse.Success(http.StatusOK, counters))
}

// func (h *Handler) GetRouterCounters(ctx *gin.Context) {

// 	var header routecountermodel.RouteCounterHeader
// 	var uri routecountermodel.RouteCounterUri
// 	var query routecountermodel.RouteCounterQuery
// 	var request routecountermodel.RouteCounterRequest

// 	err := ctx.ShouldBindHeader(&header)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
// 		return
// 	}

// 	err = ctx.ShouldBindUri(&uri)
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

// 	routeCounter, err := h.service.GetRouteCounter(request.Id)

// 	if err != nil {
// 		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
// 		return
// 	}
// 	ctx.JSON(appresponse.Success(http.StatusOK, map[string]any{
// 		"header":  header,
// 		"query":   query,
// 		"request": request,
// 		"result":  routeCounter,
// 	}))
// }

// func (h *Handler) GetAllRouteCounter(ctx *gin.Context) {

// 	list, err := h.service.GetAllRouteCounter()

// 	if err != nil {
// 		ctx.AbortWithStatusJSON(appresponse.Error(http.StatusBadRequest, err))
// 		return
// 	}
// 	ctx.JSON(appresponse.Success(http.StatusOK, list))
// }
