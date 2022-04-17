package services

import (
	"context"

	"github.com/Anixy/event-api-golang/model/domain"
	"github.com/Anixy/event-api-golang/model/web"
)

type EventService interface {
	Create(ctx context.Context, request web.CreateOrUpdateEventRequest, user domain.User) (domain.Event, error)
	Update(ctx context.Context, request web.CreateOrUpdateEventRequest, user domain.Event) (domain.Event, error)
	FindAll(ctx context.Context) ([]domain.Event, error)
	FindById(ctx context.Context, event domain.Event) (domain.Event, error)
}