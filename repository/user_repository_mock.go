package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Anixy/event-api-golang/model/domain"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (userRepository *UserRepositoryMock) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	ret := userRepository.Called(ctx, tx, user)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(context.Context, *sql.Tx, domain.User) domain.User); ok {
		r0 = rf(ctx, tx, user)
	} else {
		r0 = user
	}
	return r0
}

func (userRepository *UserRepositoryMock) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	arguments := userRepository.Mock.Called(user)

	if arguments.Get(0) == nil{
		return user
	} else {
		user := arguments.Get(0).(domain.User)
		return user
	}
}

func (userRepository *UserRepositoryMock) Delete(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	arguments := userRepository.Mock.Called(user.Id)

	if arguments.Get(0) == nil{
		return user
	} else {
		user := arguments.Get(0).(domain.User)
		return user
	}
}

func (userRepository *UserRepositoryMock) FindById(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	arguments := userRepository.Mock.Called(user)

	if arguments.Get(0) == nil{
		return user, errors.New("user not found")
	} else {
		user := arguments.Get(0).(domain.User)
		return user, nil
	}
}

func (userRepository *UserRepositoryMock) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	arguments := userRepository.Mock.Called()

	if arguments.Get(0) == nil{
		return nil
	} else {
		user := arguments.Get(0).([]domain.User)
		return user
	}
}

func (userRepository *UserRepositoryMock) FindByEmail(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	arguments := userRepository.Mock.Called(user)

	if arguments.Get(0) == nil{
		return user, errors.New("user not found")
	} else {
		user := arguments.Get(0).(domain.User)
		return user, nil
	}
}