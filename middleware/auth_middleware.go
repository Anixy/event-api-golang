package middleware

import (
	"net/http"

	"github.com/Anixy/event-api-golang/helpers"
	"github.com/Anixy/event-api-golang/model/web"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header["Authorization"][0]
		jwtToken, err := helpers.GetJwtTokenFromBearer(bearerToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
				Code: http.StatusUnauthorized,
				Status: "UNAUTORIZED",
				Data: err.Error(),
			})
			return
		}
		err = helpers.VerifyJwtToken(jwtToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
				Code: http.StatusUnauthorized,
				Status: "UNAUTORIZED",
				Data: err.Error(),
			})
			return
		}
		c.Next()
	}
}
