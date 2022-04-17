package controllers

import (
	"net/http"

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
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponse{
			Code: http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: err.Error(),
		})
		return
	}
	user, err := userController.userService.Register(c, userRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, web.WebResponse{
			Code: http.StatusUnprocessableEntity,
			Status: "UNPROCESSABLE ENTITY",
			Data: err.Error(),
		})
		return
	}
	c.JSON(200, web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: web.UserResponse{
			Id: user.Id,
			Name: user.Name,
			Email: user.Email,
		},
	})
}

func (userController *UserControllerImpl) Login(c *gin.Context)  {
	userRequest := web.LoginUserRequest{}
	err := c.ShouldBind(&userRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponse{
			Code: http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: err.Error(),
		})
		return
	}
	token, err := userController.userService.Login(c, userRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
			Code: http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data: err.Error(),
		})
		return
	}
	c.JSON(200, web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: token,
	})
}