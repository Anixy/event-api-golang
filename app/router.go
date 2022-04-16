package app

import (
	"github.com/Anixy/event-api-golang/controllers"
	"github.com/Anixy/event-api-golang/middleware"
	"github.com/Anixy/event-api-golang/repository"
	"github.com/Anixy/event-api-golang/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	db := GetDBConnection()
	userRepository := repository.NewUserRepositoryImpl()
	userService := services.NewUserServiceImpl(userRepository, db)
	userController := controllers.NewUserControllerImpl(userService)
	r := gin.Default()

	auth := r.Group("api/v1/auth")
	auth.POST("/register", userController.Register)
	auth.POST("/login", userController.Login)
	v1 := r.Group("api/v1")
	v1.Use(middleware.AuthMiddleware())
	v1.Use(gin.Recovery())
	v1.GET("/secret-data", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Secret data",
		})
	})
	return r
}