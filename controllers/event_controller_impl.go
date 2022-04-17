package controllers

import (
	"net/http"
	"strconv"

	"github.com/Anixy/event-api-golang/helpers"
	"github.com/Anixy/event-api-golang/model/domain"
	"github.com/Anixy/event-api-golang/model/web"
	"github.com/Anixy/event-api-golang/services"
	"github.com/gin-gonic/gin"
)

type EventControllerImpl struct {
	EventService services.EventService
}

func NewEventControllerImpl(eventService services.EventService) EventController {
	return &EventControllerImpl{
		EventService: eventService,
	}
}

func (eventController *EventControllerImpl) Create(c *gin.Context)  {
	bearerToken := c.Request.Header["Authorization"][0]
	jwtToken, err := helpers.GetJwtTokenFromBearer(bearerToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, web.WebResponse{
			Code: http.StatusUnprocessableEntity,
			Status: "UNPROCESSABLE ENTITY",
			Data: err.Error(),
		})
		return
	}
	user, err := helpers.GetJwtClaim(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, web.WebResponse{
			Code: http.StatusUnprocessableEntity,
			Status: "UNPROCESSABLE ENTITY",
			Data: err.Error(),
		})
		return
	}
	eventRequest := web.CreateOrUpdateEventRequest{}
	err = c.ShouldBind(&eventRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponse{
			Code: http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: err.Error(),
		})
		return
	}

	event, err := eventController.EventService.Create(c, eventRequest, user)
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
		Data: web.EventResponse{
			Id: event.Id,
			Title: event.Title,
			User: web.UserResponse{
				Id: event.User.Id,
				Name: event.User.Name,
				Email: event.User.Email,
			},
			StartDate: event.StartDate,
			EndDate: event.EndDate,
			Description: event.Description,
			Type: event.Type,
		},
	})
}

func (eventController *EventControllerImpl) Update(c *gin.Context)  {
	eventId := c.Params.ByName("eventId")
	eventIdInt, err := strconv.Atoi(eventId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponse{
			Code: http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: err.Error(),
		})
		return
	}
	bearerToken := c.Request.Header["Authorization"][0]
	jwtToken, err := helpers.GetJwtTokenFromBearer(bearerToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, web.WebResponse{
			Code: http.StatusUnprocessableEntity,
			Status: "UNPROCESSABLE ENTITY",
			Data: err.Error(),
		})
		return
	}
	user, err := helpers.GetJwtClaim(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, web.WebResponse{
			Code: http.StatusUnprocessableEntity,
			Status: "UNPROCESSABLE ENTITY",
			Data: err.Error(),
		})
		return
	}
	eventRequest := web.CreateOrUpdateEventRequest{}
	err = c.ShouldBind(&eventRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponse{
			Code: http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: err.Error(),
		})
		return
	}

	event := domain.Event{
		Id: eventIdInt,
		User: user,
	}

	event, err = eventController.EventService.Update(c, eventRequest, event)
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
		Data: web.EventResponse{
			Id: event.Id,
			Title: event.Title,
			User: web.UserResponse{
				Id: event.User.Id,
				Name: event.User.Name,
				Email: event.User.Email,
			},
			StartDate: event.StartDate,
			EndDate: event.EndDate,
			Description: event.Description,
			Type: event.Type,
		},
	})
}

func (eventController *EventControllerImpl) FindAll(c *gin.Context)  {
	events, err := eventController.EventService.FindAll(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, web.WebResponse{
			Code: http.StatusUnprocessableEntity,
			Status: "UNPROCESSABLE ENTITY",
			Data: err.Error(),
		})
		return
	}
	eventsResponses := []web.EventResponse{}
	for _, event := range events {
		eventResponses := web.EventResponse{
			Id: event.Id,
			Title: event.Title,
			User: web.UserResponse{
				Id: event.User.Id,
				Name: event.User.Name,
				Email: event.User.Email,
			},
			StartDate: event.StartDate,
			EndDate: event.EndDate,
			Description: event.Description,
			Type: event.Type,
		}
		eventsResponses = append(eventsResponses, eventResponses)
	}

	c.JSON(200, web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: eventsResponses,
	})
}

func (eventController *EventControllerImpl) FindById(c *gin.Context)  {
	eventId := c.Params.ByName("eventId")
	eventIdInt, err := strconv.Atoi(eventId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponse{
			Code: http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: err.Error(),
		})
		return
	}

	event := domain.Event{
		Id: eventIdInt,
	}

	event, err = eventController.EventService.FindById(c, event)
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
		Data: web.EventResponse{
			Id: event.Id,
			Title: event.Title,
			User: web.UserResponse{
				Id: event.User.Id,
				Name: event.User.Name,
				Email: event.User.Email,
			},
			StartDate: event.StartDate,
			EndDate: event.EndDate,
			Description: event.Description,
			Type: event.Type,
		},
		
	})
}