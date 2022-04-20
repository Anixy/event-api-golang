package services

import (
	"context"

	"github.com/Anixy/event-api-golang/model/domain"
	"github.com/Anixy/event-api-golang/model/web"
)

type EventService interface {
	Create(ctx context.Context, request web.CreateOrUpdateEventRequest, user domain.User) (domain.Event, error)
	Update(ctx context.Context, event web.CreateOrUpdateEventRequest,eventId int, user domain.User) (domain.Event, error)
	Delete(ctx context.Context, eventId int, user domain.User) (domain.Event, error)
	FindAll(ctx context.Context) ([]domain.Event, error)
	FindById(ctx context.Context, event domain.Event) (domain.Event, error)
	FindByUserId(ctx context.Context, user domain.User) ([]domain.Event, error)
	RegisterParticipant(ctx context.Context, participant domain.Participant) (domain.Participant, error)
	FindParticipantByEventId(ctx context.Context, eventId int) (domain.Event, []domain.Participant, error)
}