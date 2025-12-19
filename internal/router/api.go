package router

import (
	authrouter "e-ticket/internal/domain/auth/router"
	routecounterrouter "e-ticket/internal/domain/route_counter/router"
	schedulerouter "e-ticket/internal/domain/schedule/router"
	subuserrouter "e-ticket/internal/domain/sub_user/router"
	vehiclerouter "e-ticket/internal/domain/vehicle/router"
	"e-ticket/internal/middleware"
	appdatabase "e-ticket/pkg/database"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *appdatabase.DbEntity) *gin.Engine {

	var router *gin.Engine = gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "E-Ticket")
	})

	v1Route := router.Group("/v1")
	{
		authrouter.AuthRoutes(v1Route, db)

		// Apply middleware
		v1Route.Use(middleware.AuthToken)

		companyRoutes := v1Route.Group("company/:companyId")
		{
			HomeRoutes(companyRoutes, db)
			subuserrouter.SubUserRoutes(companyRoutes, db)
			vehiclerouter.VehicleRoutes(companyRoutes, db)
			routecounterrouter.RouteCounterRoutes(companyRoutes, db)
			schedulerouter.ScheduleRoutes(companyRoutes, db)
		}
	}

	return router
}
