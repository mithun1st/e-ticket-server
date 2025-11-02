package router

import (
	homehandler "e-ticket/internal/domain/home/handler"
	homerepository "e-ticket/internal/domain/home/repository"
	homeservice "e-ticket/internal/domain/home/service"
	appdatabase "e-ticket/pkg/database"

	"github.com/gin-gonic/gin"
)

func HomeRoutes(rg *gin.RouterGroup, db *appdatabase.DbEntity) {

	// Initialize dependencies
	repository := homerepository.NewHomeRepository(db)
	service := homeservice.NewHomeService(*repository)
	handler := homehandler.NewHomeHandler(*service)

	homeRouter := rg.Group("/home")
	{
		homeRouter.GET("company", handler.GetCompanyHome)
		homeRouter.GET("owner", handler.GetOwnerHome)
	}
}
