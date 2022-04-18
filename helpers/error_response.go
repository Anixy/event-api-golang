package helpers

import (
	"net/http"

	"github.com/Anixy/event-api-golang/model/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidationErrorResponse(c *gin.Context, validationError validator.ValidationErrors) {
	out := make([]web.ValidationErrorResponse, len(validationError))
	for i, fe := range validationError {
		out[i] = web.ValidationErrorResponse{
			Params: fe.Field(), 
			Message: ValidationError(fe),
		}
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponse{
		Code: http.StatusBadRequest,
		Status: "BAD REQUEST",
		Data: out,
	})
}

func BadRequestErrorResponse(c *gin.Context, err error)  {
	c.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponse{
		Code: http.StatusBadRequest,
		Status: "BAD REQUEST",
		Data: err.Error(),
	})
}

func UnprocessableEntityErrorResponse(c *gin.Context, err error)  {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, web.WebResponse{
		Code: http.StatusUnprocessableEntity,
		Status: "UNPROCESSABLE ENTITY",
		Data: err.Error(),
	})
}

func ForbiddenErrorResponse(c *gin.Context, err error)  {
	c.AbortWithStatusJSON(http.StatusForbidden, web.WebResponse{
		Code: http.StatusForbidden,
		Status: "FORBIDDEN",
		Data: err.Error(),
	})
}

func UnauthorizedErrorResponse(c *gin.Context, err error)  {
	c.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
		Code: http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
		Data: err.Error(),
	})
}