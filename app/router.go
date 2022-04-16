package app

import (
	"github.com/Anixy/event-api-golang/controllers"
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
	r.POST("api/v1/register", userController.Register)
	r.POST("api/v1/login", userController.Login)
	return r
}