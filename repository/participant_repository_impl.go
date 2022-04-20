package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Anixy/event-api-golang/model/domain"
)

type ParticipantRepositoryImpl struct {
}

func NewParticipantRepositoryImpl() ParticipantRepository {
	return &ParticipantRepositoryImpl{}
}

func (participantRepository *ParticipantRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, participant domain.Participant) (domain.Participant, error) {
	sql := `INSERT INTO participants(user_id, event_id, created_at, updated_at) VALUES(?,?,?,?)`
	createdAt := time.Now()
	result, err := tx.ExecContext(ctx, sql, participant.User.Id, participant.Event.Id, createdAt, createdAt)
	if err != nil {
		return participant, err
	}
	id, err := result.LastInsertId()
	if err != nil{
		return participant, err
	}
	participant.Id = int(id)
	participant.CreatedAt = createdAt
	participant.UpdatedAt = createdAt
	return participant, nil
}

func (participantRepository *ParticipantRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, participant domain.Participant) (domain.Participant, error) {
	sql := `UPDATE participants SET user_id = ?, event_id = ?, updated_at = ? WHERE id = ?`
	updatedAt := time.Now()
	_, err := tx.ExecContext(ctx, sql, participant.User.Id, participant.Event.Id, updatedAt, participant.Id)
	if err != nil {
		return participant, err
	}
	return participant, nil
}

func (participantRepository *ParticipantRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, participant domain.Participant) (domain.Participant, error) {
	sql := `DELETE FROM participants WHERE id = ?`
	_, err := tx.ExecContext(ctx, sql, participant.Id)
	if err != nil {
		return participant, err
	}
	return participant, nil
}

func (participantRepository *ParticipantRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, participant domain.Participant) (domain.Participant, error) {
	sql := `SELECT 
	participants.id, 
	participants.created_at, 
	participants.updated_at,
	users.id, 
	users.name, 
	users.email, 
	users.password, 
	users.created_at, 
	users.updated_at,
	events.id,
	events.title,
	events.start_date, 
	events.end_date, 
	events.description, 
	events.type, 
	events.created_at, 
	events.updated_at, 
    eventuser.id, 
	eventuser.name, 
	eventuser.email, 
	eventuser.password, 
	eventuser.created_at, 
	eventuser.updated_at
	FROM participants 
	INNER JOIN users ON users.id = participants.user_id
	INNER JOIN events ON events.id = participants.event_id
    INNER JOIN users eventuser ON eventuser.id = events.user_id
	WHERE id = ? 
	LIMIT 1`
	rows, err := tx.QueryContext(ctx, sql, participant.Id)
	if err != nil {
		return participant, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(
			&participant.Id,
			&participant.CreatedAt,
			&participant.UpdatedAt,
			&participant.User.Id,
			&participant.User.Name,
			&participant.User.Email,
			&participant.User.Password,
			&participant.User.CreatedAt,
			&participant.User.UpdatedAt,
			&participant.Event.Id,
			&participant.Event.Title, 
			&participant.Event.StartDate, 
			&participant.Event.EndDate,
			&participant.Event.Description,
			&participant.Event.Type, 
			&participant.Event.CreatedAt, 
			&participant.Event.UpdatedAt,
			&participant.Event.User.Id,
			&participant.Event.User.Name,
			&participant.Event.User.Email,
			&participant.Event.User.Password,
			&participant.Event.User.CreatedAt,
			&participant.Event.User.UpdatedAt,
		)
		return participant, nil
	}else {
		return participant, errors.New("participant not found")
	}
}

func (participantRepository *ParticipantRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Participant, error) {
	sql := `SELECT 
	participants.id, 
	participants.created_at, 
	participants.updated_at,
	users.id, 
	users.name, 
	users.email, 
	users.password, 
	users.created_at, 
	users.updated_at,
	events.id,
	events.title,
	events.start_date, 
	events.end_date, 
	events.description, 
	events.type, 
	events.created_at, 
	events.updated_at, 
    eventuser.id, 
	eventuser.name, 
	eventuser.email, 
	eventuser.password, 
	eventuser.created_at, 
	eventuser.updated_at
	FROM participants 
	INNER JOIN users ON users.id = participants.user_id
	INNER JOIN events ON events.id = participants.event_id
    INNER JOIN users eventuser ON eventuser.id = events.user_id`

	rows, err := tx.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	participants := []domain.Participant{}

	for rows.Next() {
		participant := domain.Participant{}
		rows.Scan(
			&participant.Id,
			&participant.CreatedAt,
			&participant.UpdatedAt,
			&participant.User.Id,
			&participant.User.Name,
			&participant.User.Email,
			&participant.User.Password,
			&participant.User.CreatedAt,
			&participant.User.UpdatedAt,
			&participant.Event.Id,
			&participant.Event.Title, 
			&participant.Event.StartDate, 
			&participant.Event.EndDate,
			&participant.Event.Description,
			&participant.Event.Type, 
			&participant.Event.CreatedAt, 
			&participant.Event.UpdatedAt,
			&participant.Event.User.Id,
			&participant.Event.User.Name,
			&participant.Event.User.Email,
			&participant.Event.User.Password,
			&participant.Event.User.CreatedAt,
			&participant.Event.User.UpdatedAt,
		)
		participants = append(participants, participant)
	}
	return participants, nil
}

func (participantRepository *ParticipantRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, user domain.User) ([]domain.Participant, error) {
	sql := `SELECT 
	participants.id, 
	participants.created_at, 
	participants.updated_at,
	users.id, 
	users.name, 
	users.email, 
	users.password, 
	users.created_at, 
	users.updated_at,
	events.id,
	events.title,
	events.start_date, 
	events.end_date, 
	events.description, 
	events.type, 
	events.created_at, 
	events.updated_at, 
    eventuser.id, 
	eventuser.name, 
	eventuser.email, 
	eventuser.password, 
	eventuser.created_at, 
	eventuser.updated_at
	FROM participants 
	INNER JOIN users ON users.id = participants.user_id
	INNER JOIN events ON events.id = participants.event_id
    INNER JOIN users eventuser ON eventuser.id = events.user_id
	WHERE user_id = ? 
	LIMIT 1`
	rows, err := tx.QueryContext(ctx, sql, user.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	participants := []domain.Participant{}
	for rows.Next() {
		participant := domain.Participant{}
		rows.Scan(
			&participant.Id,
			&participant.CreatedAt,
			&participant.UpdatedAt,
			&participant.User.Id,
			&participant.User.Name,
			&participant.User.Email,
			&participant.User.Password,
			&participant.User.CreatedAt,
			&participant.User.UpdatedAt,
			&participant.Event.Id,
			&participant.Event.Title, 
			&participant.Event.StartDate, 
			&participant.Event.EndDate,
			&participant.Event.Description,
			&participant.Event.Type, 
			&participant.Event.CreatedAt, 
			&participant.Event.UpdatedAt,
			&participant.Event.User.Id,
			&participant.Event.User.Name,
			&participant.Event.User.Email,
			&participant.Event.User.Password,
			&participant.Event.User.CreatedAt,
			&participant.Event.User.UpdatedAt,
		)
		participants = append(participants, participant)
	}
	return participants, nil
}

func (participantRepository *ParticipantRepositoryImpl) FindByEventId(ctx context.Context, tx *sql.Tx, event domain.Event) ([]domain.Participant, error) {
	sql := `SELECT 
	participants.id, 
	participants.event_id,
	participants.created_at, 
	participants.updated_at,
	users.id, 
	users.name, 
	users.email, 
	users.password, 
	users.created_at, 
	users.updated_at
	FROM participants 
	INNER JOIN users ON users.id = participants.user_id
	WHERE event_id = ?`
	rows, err := tx.QueryContext(ctx, sql, event.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	participants := []domain.Participant{}
	for rows.Next() {
		participant := domain.Participant{}
		rows.Scan(
			&participant.Id,
			&participant.Event.Id,
			&participant.CreatedAt,
			&participant.UpdatedAt,
			&participant.User.Id,
			&participant.User.Name,
			&participant.User.Email,
			&participant.User.Password,
			&participant.User.CreatedAt,
			&participant.User.UpdatedAt,
		)
		participants = append(participants, participant)
	}
	fmt.Println(participants)
	return participants, nil
}

func (participantRepository *ParticipantRepositoryImpl) FindByUserIdAndEventId(ctx context.Context, tx *sql.Tx, user domain.User, event domain.Event) (domain.Participant, error) {
	sql := `SELECT 
	id,
	user_id,
	event_id,
	created_at, 
	updated_at
	FROM participants
	WHERE user_id = ? 
	AND event_id = ?
	LIMIT 1`
	participant := domain.Participant{}
	rows, err := tx.QueryContext(ctx, sql, user.Id, event.Id)
	if err != nil {
		return participant, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(
			&participant.Id,
			&participant.User.Id,
			&participant.Event.Id,
			&participant.CreatedAt,
			&participant.UpdatedAt,
		)
		return participant, nil
	}else {
		return participant, errors.New("participant not found")
	}
}