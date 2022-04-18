package middleware

import (
	"errors"

	"github.com/Anixy/event-api-golang/helpers"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header["Authorization"]
		if len(bearerToken) == 0 {
			helpers.UnauthorizedErrorResponse(c, errors.New("need bearer token"))
			return
		}
		jwtToken, err := helpers.GetJwtTokenFromBearer(bearerToken[0])
		if err != nil {
			helpers.UnauthorizedErrorResponse(c, err)
			return
		}
		err = helpers.VerifyJwtToken(jwtToken)
		if err != nil {
			helpers.UnauthorizedErrorResponse(c, err)
			return
		}
		c.Next()
	}
}
