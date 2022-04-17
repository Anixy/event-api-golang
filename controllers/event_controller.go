package controllers

import "github.com/gin-gonic/gin"

type EventController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
}