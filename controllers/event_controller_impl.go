package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Anixy/event-api-golang/helpers"
	"github.com/Anixy/event-api-golang/model/domain"
	"github.com/Anixy/event-api-golang/model/web"
	"github.com/Anixy/event-api-golang/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		helpers.UnprocessableEntityErrorResponse(c, err)
		return
	}
	user, err := helpers.GetJwtClaim(jwtToken)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
		return
	}
	eventRequest := web.CreateOrUpdateEventRequest{}
	err = c.ShouldBind(&eventRequest)
	if err != nil {
		var ValidationError validator.ValidationErrors
		if errors.As(err, &ValidationError) {
			helpers.ValidationErrorResponse(c, ValidationError)
			return
		}
		helpers.BadRequestErrorResponse(c, err)
		return
	}

	event, err := eventController.EventService.Create(c, eventRequest, user)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
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
		helpers.BadRequestErrorResponse(c, err)
		return
	}
	bearerToken := c.Request.Header["Authorization"][0]
	jwtToken, err := helpers.GetJwtTokenFromBearer(bearerToken)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
		return
	}
	user, err := helpers.GetJwtClaim(jwtToken)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
		return
	}
	eventRequest := web.CreateOrUpdateEventRequest{}
	err = c.ShouldBind(&eventRequest)
	if err != nil {
		var ValidationError validator.ValidationErrors
		if errors.As(err, &ValidationError) {
			helpers.ValidationErrorResponse(c, ValidationError)
			return
		}
		helpers.BadRequestErrorResponse(c, err)
		return
	}
	event, err := eventController.EventService.Update(c, eventRequest, eventIdInt, user)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
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
		helpers.UnprocessableEntityErrorResponse(c, err)
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
		helpers.BadRequestErrorResponse(c, err)
		return
	}

	event := domain.Event{
		Id: eventIdInt,
	}

	event, err = eventController.EventService.FindById(c, event)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
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

func (eventController *EventControllerImpl) Delete(c *gin.Context)  {
	eventId := c.Params.ByName("eventId")
	eventIdInt, err := strconv.Atoi(eventId)
	if err != nil {
		helpers.BadRequestErrorResponse(c, err)
		return
	}

	bearerToken := c.Request.Header["Authorization"][0]
	jwtToken, err := helpers.GetJwtTokenFromBearer(bearerToken)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
		return
	}
	user, err := helpers.GetJwtClaim(jwtToken)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
		return
	}
	event, err := eventController.EventService.Delete(c, eventIdInt, user) 
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
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

func (eventController *EventControllerImpl) FindByUserId(c *gin.Context)  {

	bearerToken := c.Request.Header["Authorization"][0]
	jwtToken, err := helpers.GetJwtTokenFromBearer(bearerToken)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
		return
	}
	user, err := helpers.GetJwtClaim(jwtToken)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
		return
	}
	events, err := eventController.EventService.FindByUserId(c, user)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
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

func (eventController *EventControllerImpl) RegisterParticipant(c *gin.Context)  {
	eventId := c.Params.ByName("eventId")
	eventIdInt, err := strconv.Atoi(eventId)
	if err != nil {
		helpers.BadRequestErrorResponse(c, err)
		return
	}
	bearerToken := c.Request.Header["Authorization"][0]
	jwtToken, err := helpers.GetJwtTokenFromBearer(bearerToken)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
		return
	}
	user, err := helpers.GetJwtClaim(jwtToken)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
		return
	}
	participant := domain.Participant{
		User: user,
		Event: domain.Event{
			Id: eventIdInt,
		},
	}
	participant, err = eventController.EventService.RegisterParticipant(c, participant)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
		return
	}
	c.JSON(200, web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: web.ParticipantResponse{
			User: web.UserResponse{
				Id: participant.User.Id,
				Name: participant.User.Name,
				Email: participant.User.Email,
			},
			Event: web.EventResponse{
				Id: participant.Event.Id,
				Title: participant.Event.Title,
				StartDate: participant.Event.StartDate,
				EndDate: participant.Event.EndDate,
				Description: participant.Event.Description,
				Type: participant.Event.Type,
				User: web.UserResponse{
					Id: participant.Event.User.Id,
					Name: participant.Event.User.Name,
					Email: participant.Event.User.Email,
				},
			},
		},
	})
}

func (eventController *EventControllerImpl) FindParticipantByEventId(c *gin.Context)  {
	eventId := c.Params.ByName("eventId")
	eventIdInt, err := strconv.Atoi(eventId)
	if err != nil {
		helpers.BadRequestErrorResponse(c, err)
		return
	}
	event, participants, err := eventController.EventService.FindParticipantByEventId(c, eventIdInt)
	if err != nil {
		helpers.UnprocessableEntityErrorResponse(c, err)
		return
	}

	eventResponse := web.EventResponse{
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

	participantsResponse := []web.UserParticipant{}
	for _, participant := range participants {
		participantResponse := web.UserParticipant{
			Id: participant.User.Id,
			ParticipantId: participant.Id,
			Name: participant.User.Name,
			Email: participant.User.Email,
		}
		participantsResponse = append(participantsResponse, participantResponse)
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: web.EventParticipantResponse{
			Event: eventResponse,
			Participants: participantsResponse,
		},
	})

}