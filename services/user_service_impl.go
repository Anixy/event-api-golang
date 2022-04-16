package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Anixy/event-api-golang/helpers"
	"github.com/Anixy/event-api-golang/model/domain"
	"github.com/Anixy/event-api-golang/model/web"
	"github.com/Anixy/event-api-golang/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB sql.DB
}

func NewUserServiceImpl(userRepository repository.UserRepository, db *sql.DB) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB: *db,
	}
}

func (userService *UserServiceImpl) Register(ctx context.Context, request web.RegisterUserRequest) domain.User {
	if request.Password != request.VerifyPassword {
		err := errors.New("password and verify password not match")
		helpers.PanicIfError(err)
	}
	tx, err :=userService.DB.Begin()
	helpers.PanicIfError(err)
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	helpers.PanicIfError(err)
	user :=userService.UserRepository.Save(ctx, tx, domain.User{
		Name: request.Name,
		Email: request.Email,
		Password: string(bytes),
	})
	tx.Commit()
	return user
}

func (userService *UserServiceImpl) Login(ctx context.Context, request web.LoginUserRequest) string {
	tx, err := userService.DB.Begin()
	helpers.PanicIfError(err)
	user, err := userService.UserRepository.FindByEmail(ctx, tx, domain.User{
		Email: request.Email,
	})
	helpers.PanicIfError(err)
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err !=nil {
		err := errors.New("wrong password")
		helpers.PanicIfError(err)
	}
	token := helpers.CreateJwtToken(user)
	return token
}