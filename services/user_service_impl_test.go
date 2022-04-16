package services

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Anixy/event-api-golang/model/domain"
	"github.com/Anixy/event-api-golang/model/web"
	"github.com/Anixy/event-api-golang/repository"
	"github.com/Anixy/event-api-golang/test"
	"github.com/stretchr/testify/assert"
)

func TruncateUser(db *sql.DB)  {
	db.Exec("TRUNCATE TABLE users;")
}

func TestUserServiceImpl_Register(t *testing.T) {
	db := test.ConnectToDB()
	TruncateUser(db)
	type args struct {
		ctx     context.Context
		request web.RegisterUserRequest
	}
	tests := []struct {
		name        string
		userService UserService
		args        args
		want        domain.User
	}{
		{
			name: "REGISTER USER TEST",
			userService: NewUserServiceImpl(repository.NewUserRepositoryImpl(), db),
			args: args{
				ctx: context.Background(),
				request: web.RegisterUserRequest{
					Name: "Budi",
					Email: "budi@gmail.com",
					Password: "secretpassword",
					VerifyPassword: "secretpassword",
				},
			},
			want: domain.User{
				Id: 1,
				Name: "Budi",
				Email: "budi@gmail.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, _ := tt.userService.Register(tt.args.ctx, tt.args.request)
			assert.Equal(t, tt.want.Id, user.Id)
			assert.Equal(t, tt.want.Name, user.Name)
			assert.Equal(t, tt.want.Email, user.Email)
		})
	}
}

func TestUserServiceImpl_Login(t *testing.T) {
	db := test.ConnectToDB()
	TruncateUser(db)
	userService := NewUserServiceImpl(repository.NewUserRepositoryImpl(), db)
	userService.Register(context.Background(), web.RegisterUserRequest{
		Name: "Budi",
		Email: "budi@gmail.com",
		Password: "secretpassword",
		VerifyPassword: "secretpassword",
	})
	type args struct {
		ctx     context.Context
		request web.LoginUserRequest
	}
	tests := []struct {
		name        string
		userService UserService
		args        args
		want        string
	}{
		{
			name: "TEST LOGIN USER",
			userService: userService,
			args: args{
				ctx: context.Background(),
				request: web.LoginUserRequest{
					Email: "budi@gmail.com",
					Password: "secretpassword",
				},
			},	
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, _ := tt.userService.Login(tt.args.ctx, tt.args.request)
			assert.NotEmpty(t, token)
		})
	}
}
