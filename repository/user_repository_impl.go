package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Anixy/event-api-golang/helpers"
	"github.com/Anixy/event-api-golang/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

func (userRepository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	sql := `INSERT INTO users(name, email, password, created_at, updated_at) VALUES(?,?,?,?,?)`
	createdAt := time.Now()
	result, err := tx.ExecContext(ctx, sql, user.Name, user.Email, user.Password, createdAt, createdAt)
	helpers.PanicIfError(err)
	id, err := result.LastInsertId()
	helpers.PanicIfError(err)
	user.Id = int(id)
	user.CreatedAt = createdAt
	user.UpdatedAt = createdAt
	return user
}

func (userRepository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	sql := `UPDATE users SET name = ?, email = ?, password = ?, updated_at = ? WHERE id = ?`
	updatedAt := time.Now()
	_, err := tx.ExecContext(ctx, sql, user.Name, user.Email, user.Password, updatedAt, user.Id)
	helpers.PanicIfError(err)
	helpers.PanicIfError(err)
	return user
}

func (userRepository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	sql := `DELETE FROM users WHERE id = ?`
	_, err := tx.ExecContext(ctx, sql, user.Id)
	helpers.PanicIfError(err)
	return user
}

func (userRepository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	sql := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = ? LIMIT 1`
	rows, err := tx.QueryContext(ctx, sql, user.Id)
	helpers.PanicIfError(err)

	if rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		return user, nil
	}else {
		return user, errors.New("user not found")
	}
}

func (userRepository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	sql := `SELECT id, name, email, password, created_at, updated_at FROM users`

	rows, err := tx.QueryContext(ctx, sql)
	helpers.PanicIfError(err)

	users := []domain.User{}

	for rows.Next() {
		user := domain.User{}
		rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		users = append(users, user)
	}

	return users
}

func (userRepository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	sql := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ? LIMIT 1`
	rows, err := tx.QueryContext(ctx, sql, user.Email)
	helpers.PanicIfError(err)
	if rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		return user, nil
	}else {
		return user, errors.New("user not found")
	}
}

