package schedulerouter

import (
	schedulehandler "e-ticket/internal/domain/schedule/handler"
	schedulerepository "e-ticket/internal/domain/schedule/repository"
	scheduleservice "e-ticket/internal/domain/schedule/service"
	appdatabase "e-ticket/pkg/database"

	"github.com/gin-gonic/gin"
)

func ScheduleRoutes(rg *gin.RouterGroup, db *appdatabase.DbEntity) {

	// Initialize dependencies
	repository := schedulerepository.NewScheduleRepository(db)
	service := scheduleservice.NewScheduleService(*repository)
	handler := schedulehandler.NewScheduleHandler(*service)

	scheduleRouter := rg.Group("/schedules")
	{
		scheduleRouter.GET("", handler.GetSchedules)
		scheduleRouter.POST("", handler.CreateSchedules)
		// scheduleRouter.POST("")
		// scheduleRouter.PUT(":scheduleId")
		// scheduleRouter.DELETE("/:scheduleId")
	}
}
