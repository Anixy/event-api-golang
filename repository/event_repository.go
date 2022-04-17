package repository

import (
	"context"
	"database/sql"

	"github.com/Anixy/event-api-golang/model/domain"
)

type EventRepository interface {
	Save(ctx context.Context, tx *sql.Tx, event domain.Event) (domain.Event, error)
	Update(ctx context.Context, tx *sql.Tx, event domain.Event) (domain.Event, error)
	Delete(ctx context.Context, tx *sql.Tx, event domain.Event) (domain.Event, error)
	FindById(ctx context.Context, tx *sql.Tx, event domain.Event) (domain.Event, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Event, error)
	FindByUserId(ctx context.Context, tx *sql.Tx, user domain.User) ([]domain.Event, error)
}