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

func (userService *UserServiceImpl) Register(ctx context.Context, request web.RegisterUserRequest) (domain.User, error) {
	user := domain.User{
		Name: request.Name,
		Email: request.Email,
	}
	if request.Password != request.VerifyPassword {
		return user, errors.New("password and verify password not match") 
	}
	tx, err :=userService.DB.Begin()
	if err != nil {
		return user, err
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		return user, err
	}
	user.Password = string(bytes)
	user, err =userService.UserRepository.Save(ctx, tx, user)
	if err != nil {
		return user, err
	}
	tx.Commit()
	return user, nil
}

func (userService *UserServiceImpl) Login(ctx context.Context, request web.LoginUserRequest) (string, error) {
	tx, err := userService.DB.Begin()
	if err != nil {
		return "", err
	}
	user, err := userService.UserRepository.FindByEmail(ctx, tx, domain.User{
		Email: request.Email,
	})
	if err != nil {
		return "", err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err !=nil {
		return "", errors.New("wrong password")
		
	}
	token := helpers.CreateJwtToken(user)
	return token, nil
}