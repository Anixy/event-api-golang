package controllers

import (
	"errors"

	"github.com/Anixy/event-api-golang/helpers"
	"github.com/Anixy/event-api-golang/model/web"
	"github.com/Anixy/event-api-golang/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		var ValidationError validator.ValidationErrors
		if errors.As(err, &ValidationError) {
			helpers.ValidationErrorResponse(c, ValidationError)
			return
		}
		helpers.BadRequestErrorResponse(c, err)
		return
	}
	user, err := userController.userService.Register(c, userRequest)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
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
		var ValidationError validator.ValidationErrors
		if errors.As(err, &ValidationError) {
			helpers.ValidationErrorResponse(c, ValidationError)
			return
		}
		helpers.BadRequestErrorResponse(c, err)
		return
	}
	token, err := userController.userService.Login(c, userRequest)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
		return
	}
	refreshToken, err := helpers.CreateRefreshToken(token)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
		return
	}
	c.JSON(200, web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: web.LoginUserResponse{
			Token: token,
			RefreshToken: refreshToken,
		},
	})
}

func (userController *UserControllerImpl) RefreshToken(c *gin.Context)  {
	userRequest := web.RefreshTokenRequest{}
	err := c.ShouldBind(&userRequest)
	if err != nil {
		var ValidationError validator.ValidationErrors
		if errors.As(err, &ValidationError) {
			helpers.ValidationErrorResponse(c, ValidationError)
			return
		}
		helpers.BadRequestErrorResponse(c, err)
		return
	}
	user, err := helpers.ValidateRefreshToken(userRequest.RefreshToken)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
		return
	}
	token := helpers.CreateJwtToken(user)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
		return
	}
	c.JSON(200, web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: web.LoginUserResponse{
			Token: token,
			RefreshToken: userRequest.RefreshToken,
		},
	})
}