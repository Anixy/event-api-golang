package services

import (
	"context"

	"github.com/Anixy/event-api-golang/model/domain"
	"github.com/Anixy/event-api-golang/model/web"
)

type UserService interface {
	Register(ctx context.Context, request web.RegisterUserRequest) domain.User
	Login(ctx context.Context, request web.LoginUserRequest) string
}