package router

import (
	authhandler "e-ticket/internal/domain/auth/handler"
	authrepository "e-ticket/internal/domain/auth/repository"
	authservice "e-ticket/internal/domain/auth/service"
	appdatabase "e-ticket/pkg/database"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup, db *appdatabase.DbEntity) {

	// Initialize dependencies
	repository := authrepository.NewAuthRepository(db)
	service := authservice.NewAuthService(*repository)
	handler := authhandler.NewAuthHandler(*service)

	authRouter := rg.Group("/auth")
	{
		loginRouter := authRouter.Group("/login")
		{
			loginRouter.POST("/company", func(ctx *gin.Context) {
				handler.GetAuthCompany(ctx)
			})
			loginRouter.POST("/subuser", func(ctx *gin.Context) {
				handler.GetAuthVehicle(ctx)
			})
		}
		// authRouter.GET("", handler.GetAllAuth)
		// authRouter.POST("")
		// authRouter.PUT("/:authId")
		// authRouter.DELETE("/:authId")
	}
}
