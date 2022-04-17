package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Anixy/event-api-golang/model/domain"
	"github.com/Anixy/event-api-golang/model/web"
	"github.com/Anixy/event-api-golang/repository"
)

type EventServiceImpl struct {
	EventRepository repository.EventRepository
	UserRepository repository.UserRepository
	DB *sql.DB
}

func NewEventServiceImpl(evenRepository repository.EventRepository, userRepository repository.UserRepository, db *sql.DB) EventService {
	return &EventServiceImpl{
		EventRepository: evenRepository,
		UserRepository: userRepository,
		DB: db,
	}
}

func (eventService *EventServiceImpl) Create(ctx context.Context, request web.CreateOrUpdateEventRequest, user domain.User) (domain.Event, error) {
	tx, err := eventService.DB.Begin()
	event := domain.Event{
		Title: request.Title,
		StartDate: request.StartDate,
		EndDate: request.EndDate,
		Description: request.Description,
		Type: request.Type,
	}
	if err != nil {
		return event, err
	}
	user, err = eventService.UserRepository.FindById(ctx, tx, user)
	if err != nil {
		tx.Rollback()
		return event, err
	}
	event.User = user
	event, err = eventService.EventRepository.Save(ctx, tx, event)
	if err != nil {
		tx.Rollback()
		return event, err
	}
	tx.Commit()
	return event, nil
}

func (eventService *EventServiceImpl) Update(ctx context.Context, request web.CreateOrUpdateEventRequest, eventRequest domain.Event) (domain.Event, error) {
	tx, err := eventService.DB.Begin()
	eventRequest.Title = request.Title
	eventRequest.StartDate = request.StartDate
	eventRequest.EndDate = request.EndDate
	eventRequest.Description = request.Description
	eventRequest.Type = request.Type
	if err != nil {
		return eventRequest, err
	}
	event, err := eventService.EventRepository.FindById(ctx, tx, eventRequest)
	if err != nil {
		tx.Rollback()
		return event, err
	}
	if event.User.Id != eventRequest.User.Id {
		tx.Rollback()
		return event, errors.New("not allowed request")
	}
	event, err = eventService.EventRepository.Update(ctx, tx, eventRequest)
	if err != nil {
		tx.Rollback()
		return event, err
	}
	tx.Commit()
	return event, nil
}

func (eventService *EventServiceImpl) FindAll(ctx context.Context) ([]domain.Event, error) {
	tx, err := eventService.DB.Begin()
	if err != nil {
		return nil, err
	}
	events, err := eventService.EventRepository.FindAll(ctx, tx)
	if err != nil {
		tx.Rollback()
		return events, err
	}
	tx.Commit()
	return events, nil
}

func (eventService *EventServiceImpl) FindById(ctx context.Context, event domain.Event) (domain.Event, error) {
	tx, err := eventService.DB.Begin()
	if err != nil {
		return event, err
	}
	events, err := eventService.EventRepository.FindById(ctx, tx, event)
	if err != nil {
		tx.Rollback()
		return events, err
	}
	tx.Commit()
	return events, nil
}