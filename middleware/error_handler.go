package middleware

import (
	"fmt"

	"github.com/Anixy/event-api-golang/model/web"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() { 
			if err := recover(); err != nil {
				c.JSON(500, web.WebResponse{
					Code: 500,
					Status: fmt.Sprintf("%v", err),
				})
				return
			}
		}()
		c.Next()	
	}
}