package router

import (
	vehiclehandler "e-ticket/internal/domain/vehicle/handler"
	vehiclerepository "e-ticket/internal/domain/vehicle/repository"
	vehicleservice "e-ticket/internal/domain/vehicle/service"
	appdatabase "e-ticket/pkg/database"

	"github.com/gin-gonic/gin"
)

func VehicleRoutes(rg *gin.RouterGroup, db *appdatabase.DbEntity) {

	// Initialize dependencies
	repository := vehiclerepository.NewVehicleRepository(db)
	service := vehicleservice.NewVehicleService(*repository)
	handler := vehiclehandler.NewVehicleHandler(*service)

	vehicleRouter := rg.Group("/vehicles")
	{
		// vehicleRouter.GET("/:vehicleId", handler.GetVehicle)
		vehicleRouter.GET("", handler.GetAllVehicle)
		vehicleRouter.POST("", handler.CreateVehicle)
		// vehicleRouter.PUT("/:vehicleId")
		// vehicleRouter.DELETE("/:vehicleId")
	}
}
