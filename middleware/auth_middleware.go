package middleware

import (
	"net/http"

	"github.com/Anixy/event-api-golang/helpers"
	"github.com/Anixy/event-api-golang/model/web"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header["Authorization"]
		if len(bearerToken) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
				Code: http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
				Data: "need bearer token",
			})
		}
		jwtToken, err := helpers.GetJwtTokenFromBearer(bearerToken[0])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
				Code: http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
				Data: err.Error(),
			})
			return
		}
		err = helpers.VerifyJwtToken(jwtToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
				Code: http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
				Data: err.Error(),
			})
			return
		}
		c.Next()
	}
}
