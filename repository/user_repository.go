package repository

import (
	"context"
	"database/sql"

	"github.com/Anixy/event-api-golang/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx ,user domain.User) (domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	Delete(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	FindById(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
}