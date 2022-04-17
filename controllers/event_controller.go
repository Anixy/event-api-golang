package controllers

import "github.com/gin-gonic/gin"

type EventController interface {
	Create(c *gin.Context)
}