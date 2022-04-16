package middleware

import (
	"strings"

	"github.com/Anixy/event-api-golang/helpers"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		barierToken := c.Request.Header["Authorization"][0]
		jwtToken := strings.Split(barierToken, " ")[1]
		err := helpers.VerifyJwtToken(jwtToken)
		helpers.PanicIfError(err) 
		c.Next()
	}
}
