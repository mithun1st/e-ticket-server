package router

import (
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
		AuthRoutes(v1Route, db)

		// Apply middleware
		v1Route.Use(middleware.AuthToken)

		HomeRoutes(v1Route, db)
		SubUserRoutes(v1Route, db)
		VehicleRoutes(v1Route, db)
	}

	// v2 := api.Group("/v2")
	// v2.Use(middleware.AuthToken)
	// {
	//
	// }

	return router
}
