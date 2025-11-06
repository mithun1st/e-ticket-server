package router

import (
	subuserhandler "e-ticket/internal/domain/sub_user/handler"
	subuserrepository "e-ticket/internal/domain/sub_user/repository"
	subuserservice "e-ticket/internal/domain/sub_user/service"
	appdatabase "e-ticket/pkg/database"

	"github.com/gin-gonic/gin"
)

func SubUserRoutes(rg *gin.RouterGroup, db *appdatabase.DbEntity) {

	// Initialize dependencies
	repository := subuserrepository.NewSubUserRepository(db)
	service := subuserservice.NewSubUserService(*repository)
	handler := subuserhandler.NewSubUserHandler(*service)

	subUserRouter := rg.Group("/subuser")
	{
		// subUserRouter.GET("/:subUserId", handler.GetSubUser)
		subUserRouter.GET("/all", handler.GetAllSubUser)
		subUserRouter.POST("/create", handler.CreateSubUser)
		// subUserRouter.PUT("/:subUserId")
		// subUserRouter.DELETE("/:subUserId")
	}
}
