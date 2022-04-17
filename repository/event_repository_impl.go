package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Anixy/event-api-golang/model/domain"
)

type EventRepositoryImpl struct {
}

func NewEventRepositoryImpl() EventRepository {
	return &EventRepositoryImpl{}
}

func (eventRepository *EventRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, event domain.Event) (domain.Event, error) {
	sql := `INSERT INTO events(user_id, title, start_date, end_date, description, type, created_at, updated_at) VALUES(?,?,?,?,?,?,?,?)`
	createdAt := time.Now()
	result, err := tx.ExecContext(ctx, sql, event.User.Id, event.Title, event.StartDate, event.EndDate, event.Description, event.Type, createdAt, createdAt)
	if err != nil {
		return event, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return event, err
	}
	event.Id = int(id)
	event.CreatedAt = createdAt
	event.UpdatedAt = createdAt
	return event, nil
}

func (eventRepository *EventRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, event domain.Event) (domain.Event, error) {
	sql := `UPDATE events SET title = ?, start_date = ?, end_date = ?, description = ?, type = ?, updated_at = ? WHERE id = ?`
	updatedAt := time.Now()
	_, err := tx.ExecContext(ctx, sql, event.Title, event.StartDate, event.EndDate, event.Description, event.Type, updatedAt, event.Id)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (eventRepository *EventRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, event domain.Event) (domain.Event, error) {
	sql := `DELETE FROM events WHERE id = ?`
	_, err := tx.ExecContext(ctx, sql, event.Id)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (eventRepository *EventRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, event domain.Event) (domain.Event, error) {
	sql := `SELECT id, user_id, title, start_date, end_date, description, type, created_at, updated_at FROM events WHERE id = ? LIMIT 1`
	rows, err := tx.QueryContext(ctx, sql, event.Id)
	if err != nil {
		return event, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(
			&event.Id, 
			&event.User.Id, 
			&event.Title, 
			&event.StartDate, 
			&event.EndDate, 
			&event.Description, 
			&event.Type, 
			&event.CreatedAt, 
			&event.UpdatedAt,
		)
		return event, nil
	} else {
		return event, errors.New("user not found")
	}
}

func (eventRepository *EventRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Event, error) {
	sql := `SELECT id, user_id, title, start_date, end_date, description, type, created_at, updated_at FROM events`

	rows, err := tx.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []domain.Event{}

	for rows.Next() {
		event := domain.Event{}
		rows.Scan(
			&event.Id,
			&event.User.Id, 
			&event.Title, 
			&event.StartDate, 
			&event.EndDate,
			&event.Description,
			&event.Type, 
			&event.CreatedAt, 
			&event.UpdatedAt,
		)
		events = append(events, event)
	}
	return events, nil
}

func (eventRepository *EventRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, user domain.User) ([]domain.Event, error) {
	sql := `SELECT id, user_id, title, start_date, end_date, description, type, created_at, updated_at FROM events WHERE user_id = ? LIMIT 1`
	rows, err := tx.QueryContext(ctx, sql, user.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []domain.Event{}

	for rows.Next() {
		event := domain.Event{}
		rows.Scan(
			&event.Id,
			&event.User.Id, 
			&event.Title, 
			&event.StartDate, 
			&event.EndDate,
			&event.Description,
			&event.Type, 
			&event.CreatedAt, 
			&event.UpdatedAt,
		)
		events = append(events, event)
	}
	return events, nil
}
