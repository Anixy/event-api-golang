package services

import (
	"context"

	"github.com/Anixy/event-api-golang/model/domain"
	"github.com/Anixy/event-api-golang/model/web"
)

type EventService interface {
	Create(ctx context.Context, request web.CreateEventRequest, user domain.User) (domain.Event, error)
	FindAll(ctx context.Context) ([]domain.Event, error)
}