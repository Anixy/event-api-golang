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
	ParticipantRepository repository.ParticipantRepository
	DB *sql.DB
}

func NewEventServiceImpl(evenRepository repository.EventRepository, userRepository repository.UserRepository, participantRepository repository.ParticipantRepository, db *sql.DB) EventService {
	return &EventServiceImpl{
		EventRepository: evenRepository,
		UserRepository: userRepository,
		ParticipantRepository: participantRepository,
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

func (eventService *EventServiceImpl) Update(ctx context.Context, request web.CreateOrUpdateEventRequest, eventId int, user domain.User) (domain.Event, error) {
	eventRequest := domain.Event{
		Id: eventId,
		Title: request.Title,
		StartDate: request.StartDate,
		EndDate: request.EndDate,
		Description: request.Description,
		Type: request.Type,
	}
	tx, err := eventService.DB.Begin()
	if err != nil {
		return eventRequest, err
	}

	event, err := eventService.EventRepository.FindById(ctx, tx, eventRequest)
	if err != nil {
		tx.Rollback()
		return event, err
	}
	if event.User.Id != user.Id {
		tx.Rollback()
		return event, errors.New("not allowed request")
	}
	eventRequest.User = user
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

func (eventService *EventServiceImpl) Delete(ctx context.Context, eventiId int, user domain.User) (domain.Event, error) {
	event := domain.Event{
		Id: eventiId,
	}
	tx, err := eventService.DB.Begin()
	if err != nil {
		return event, err
	}
	event, err = eventService.EventRepository.FindById(ctx, tx, event)
	if err != nil {
		tx.Rollback()
	}
	if event.User.Id != user.Id {
		tx.Rollback()
		return event, errors.New("not allowed request")
	}
	event.User = user
	event, err = eventService.EventRepository.Delete(ctx, tx, event)
	if err != nil {
		tx.Rollback()
		return event, err
	}
	tx.Commit()
	return event, nil
}

func (eventService *EventServiceImpl) FindByUserId(ctx context.Context, user domain.User) ([]domain.Event, error) {
	tx, err := eventService.DB.Begin()
	if err != nil {
		return nil, err
	}
	events, err := eventService.EventRepository.FindByUserId(ctx, tx, user)
	if err != nil {
		tx.Rollback()
		return events, err
	}
	tx.Commit()
	return events, nil
}

func (eventService *EventServiceImpl) RegisterParticipant(ctx context.Context, participant domain.Participant) (domain.Participant, error) {
	tx, err := eventService.DB.Begin()
	if err != nil {
		return participant, err
	}
	user, err := eventService.UserRepository.FindById(ctx, tx, participant.User)
	if err != nil {
		tx.Rollback()
		return participant, err
	}
	participant.User = user
	event, err := eventService.EventRepository.FindById(ctx, tx, participant.Event)
	if err != nil {
		tx.Rollback()
		return participant, err
	}
	participant.Event = event
	if participant.User.Id == participant.Event.User.Id {
		tx.Rollback()
		return participant, errors.New("cannot register to your event")
	}
	_, err = eventService.ParticipantRepository.FindByUserIdAndEventId(ctx, tx, participant.User, participant.Event)
	if err == nil{
		tx.Rollback()
		return participant, errors.New("you already registered in this event")
	}
	if err != nil && err.Error() != "participant not found" {
		tx.Rollback()
		return participant, err
	}
	participant, err = eventService.ParticipantRepository.Save(ctx, tx, participant)
	if err != nil {
		tx.Rollback()
		return participant, err
	}
	tx.Commit()
	return participant, nil
}

func (eventService *EventServiceImpl) FindParticipantByEventId(ctx context.Context, eventId int) (domain.Event,[]domain.Participant, error) {
	event := domain.Event{
		Id: eventId,
	}
	tx, err := eventService.DB.Begin()
	if err != nil {
		return event, nil, err
	}
	event, err = eventService.EventRepository.FindById(ctx, tx, event)
	if err != nil {
		return event, nil, err
	}
	participants , err := eventService.ParticipantRepository.FindByEventId(ctx, tx, event)
	if err != nil {
		return event, nil, err
	}
	return event, participants, nil
}


