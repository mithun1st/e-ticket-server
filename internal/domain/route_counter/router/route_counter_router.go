package routecounterrouter

import (
	routecounterhandler "e-ticket/internal/domain/route_counter/handler"
	routecounterrepository "e-ticket/internal/domain/route_counter/repository"
	routecounterservice "e-ticket/internal/domain/route_counter/service"
	appdatabase "e-ticket/pkg/database"

	"github.com/gin-gonic/gin"
)

func RouteCounterRoutes(rg *gin.RouterGroup, db *appdatabase.DbEntity) {

	// Initialize dependencies
	repository := routecounterrepository.NewRouteCounterRepository(db)
	service := routecounterservice.NewRouteCounterService(*repository)
	handler := routecounterhandler.NewRouteCounterHandler(*service)

	routeCounterRouter := rg.Group("/counters")
	{
		routeCounterRouter.GET("", handler.GetCounters)
		routeCounterRouter.POST("", handler.CreateCounter)
	}
	routeCounterRouter = rg.Group("/routes")
	{
		routeCounterRouter.GET("", handler.GetRoutes)
		routeCounterRouter.POST("", handler.CreateRoute)
	}
	routeCounterRouter = rg.Group("/routecounters")
	{
		routeCounterRouter.GET("/:routeId", handler.GetRouteCounters)
		routeCounterRouter.POST("/:routeId", handler.CreateRouteCounter)
	}
}
