package app

import (
	"net/http"

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
	eventRepository := repository.NewEventRepositoryImpl()
	eventService := services.NewEventServiceImpl(eventRepository, userRepository, db)
	eventController := controllers.NewEventControllerImpl(eventService)
	r := gin.Default()
	auth := r.Group("api/v1/auth")
	auth.POST("/register", userController.Register)
	auth.POST("/login", userController.Login)
	v1 := r.Group("api/v1")
	v1.Use(middleware.AuthMiddleware())
	v1.POST("/event", eventController.Create)
	v1.GET("/event", eventController.FindAll)
	v1.PUT("/event/:eventId", func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{
			"message": "Not implemented",
		})
	})
	return r
}