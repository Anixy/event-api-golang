package controllers

import (
	"github.com/Anixy/event-api-golang/helpers"
	"github.com/Anixy/event-api-golang/model/domain"
	"github.com/Anixy/event-api-golang/model/web"
	"github.com/Anixy/event-api-golang/services"
	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	userService services.UserService	
}

func NewUserControllerImpl(userService services.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}


func (userController *UserControllerImpl) Register(c *gin.Context)  {
	userRequest := web.RegisterUserRequest{}
	err := c.ShouldBind(&userRequest)
	helpers.PanicIfError(err)
	user := userController.userService.Register(c, userRequest)
	c.JSON(200, web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: domain.User{
			Name: user.Name,
			Email: user.Email,
		},
	})
}

func (userController *UserControllerImpl) Login(c *gin.Context)  {
	userRequest := web.LoginUserRequest{}
	err := c.ShouldBind(&userRequest)
	helpers.PanicIfError(err)
	token := userController.userService.Login(c, userRequest)
	c.JSON(200, web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: token,
	})
}