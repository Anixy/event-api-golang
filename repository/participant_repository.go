package repository

import (
	"context"
	"database/sql"

	"github.com/Anixy/event-api-golang/model/domain"
)

type ParticipantRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.Participant) (domain.Participant, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.Participant) (domain.Participant, error)
	Delete(ctx context.Context, tx *sql.Tx, user domain.Participant) (domain.Participant, error)
	FindById(ctx context.Context, tx *sql.Tx, user domain.Participant) (domain.Participant, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Participant, error)
	FindByUserId(ctx context.Context, tx *sql.Tx, user domain.User) ([]domain.Participant, error)
	FindByEventId(ctx context.Context, tx *sql.Tx, event domain.Event) ([]domain.Participant, error)
	FindByUserIdAndEventId(ctx context.Context, tx *sql.Tx, user domain.User, event domain.Event) (domain.Participant, error)
}