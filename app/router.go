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
	eventRepository := repository.NewEventRepositoryImpl()
	eventService := services.NewEventServiceImpl(eventRepository, userRepository, db)
	eventController := controllers.NewEventControllerImpl(eventService)
	r := gin.Default()
	auth := r.Group("api/v1/auth")
	v1 := r.Group("api/v1")
	v1.Use(middleware.AuthMiddleware())
	auth.POST("/register", userController.Register)
	auth.POST("/login", userController.Login)
	v1.POST("/event", eventController.Create)
	v1.GET("/event", eventController.FindAll)
	v1.GET("/event/:eventId", eventController.FindById)
	v1.PUT("/event/:eventId", eventController.Update)
	v1.DELETE("/event/:eventId", eventController.Delete)
	v1.GET("/event/my-event", eventController.FindByUserId)
	return r
}