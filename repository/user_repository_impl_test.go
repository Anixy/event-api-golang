package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Anixy/event-api-golang/helpers"
	"github.com/Anixy/event-api-golang/model/domain"
	"github.com/Anixy/event-api-golang/test"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/stretchr/testify/assert"
)

func TruncateUser(db *sql.DB)  {
	db.Exec("TRUNCATE TABLE users;")
}

func TestUserRepositoryImpl_Save(t *testing.T) {
	db := test.ConnectToDB()
	defer db.Close()
	TruncateUser(db)
	type args struct {
		ctx  context.Context
		user domain.User
	}
	tests := []struct {
		name           string
		userRepository UserRepository
		args           args
		want           domain.User
	}{
		// TODO: Add test cases.
		{
			name: "TEST CREATE FIRST USER",
			userRepository: NewUserRepositoryImpl(),
			args: args{
				ctx: context.Background(),
				user: domain.User{
					Name: "Budi",
					Email: "budi@example.com",
					Password: "secretpassword",
				},
			},
			want: domain.User{
				Id: 1,
			},
		},
		{
			name: "TEST CREATE SECOND USER",
			userRepository: NewUserRepositoryImpl(),
			args: args{
				ctx: context.Background(),
				user: domain.User{
					Name: "John",
					Email: "john@example.com",
					Password: "secretpassword",
				},
			},
			want: domain.User{
				Id: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err :=db.Begin()
			helpers.PanicIfError(err)
			user := tt.userRepository.Save(tt.args.ctx, tx, tt.args.user)
			err = tx.Commit()
			helpers.PanicIfError(err)
			assert.Equal(t, tt.want.Id, user.Id)
		})
	}
}

func TestUserRepositoryImpl_Update(t *testing.T) {
	db := test.ConnectToDB()
	defer db.Close()
	TruncateUser(db)
	//make user data 
	userRepository := NewUserRepositoryImpl()
	tx, err := db.Begin()
	helpers.PanicIfError(err)
	userRepository.Save(context.Background(), tx, domain.User{
		Name: "Budi",
		Email: "budi@example.com",
		Password: "secretpassword",
	})
	err = tx.Commit()
	helpers.PanicIfError(err)

	type args struct {
		ctx  context.Context
		user domain.User
	}
	tests := []struct {
		name           string
		userRepository UserRepository
		args           args
		want           domain.User
	}{
		{
			name: "TEST UPDATE USER",
			userRepository: NewUserRepositoryImpl(),
			args: args{
				ctx: context.Background(),
				user: domain.User{
					Id: 1,
					Name: "John",
					Email: "budi@example.com",
					Password: "secretpassword",
				},
			},
			want: domain.User{
				Id: 1,
				Name: "John",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := db.Begin()
			helpers.PanicIfError(err)
			user := tt.userRepository.Update(tt.args.ctx, tx, tt.args.user)
			err = tx.Commit()
			helpers.PanicIfError(err)
			assert.Equal(t, tt.want.Id, user.Id)
			assert.Equal(t, tt.want.Name, user.Name)
		})
	}
}

func TestUserRepositoryImpl_Delete(t *testing.T) {
	db := test.ConnectToDB()
	defer db.Close()
	TruncateUser(db)
	//make user data 
	userRepository := NewUserRepositoryImpl()
	tx, err := db.Begin()
	helpers.PanicIfError(err)
	userRepository.Save(context.Background(), tx, domain.User{
		Name: "Budi",
		Email: "budi@example.com",
		Password: "secretpassword",
	})
	err = tx.Commit()
	helpers.PanicIfError(err)

	type args struct {
		ctx  context.Context
		user domain.User
	}
	tests := []struct {
		name           string
		userRepository UserRepository
		args           args
		want           domain.User
	}{
		{
			name: "TEST DELETE USER",
			userRepository: NewUserRepositoryImpl(),
			args: args{
				ctx: context.Background(),
				user: domain.User{
					Id: 1,
					Name: "Budi",
					Email: "budi@example.com",
					Password: "secretpassword",
				},
			},
			want: domain.User{
				Id: 1,
				Name: "Budi",
				Email: "budi@example.com",
				Password: "secretpassword",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := db.Begin()
			helpers.PanicIfError(err)
			user := tt.userRepository.Delete(tt.args.ctx, tx, tt.args.user)
			err = tx.Commit()
			helpers.PanicIfError(err)
			assert.Equal(t, tt.want.Id, user.Id)
			assert.Equal(t, tt.want.Name, user.Name)
		})
	}
}

func TestUserRepositoryImpl_FindById(t *testing.T) {
	db := test.ConnectToDB()
	defer db.Close()
	TruncateUser(db)
	//make user data 
	userRepository := NewUserRepositoryImpl()
	tx, err := db.Begin()
	helpers.PanicIfError(err)
	userRepository.Save(context.Background(), tx, domain.User{
		Name: "Budi",
		Email: "budi@example.com",
		Password: "secretpassword",
	})
	err = tx.Commit()
	helpers.PanicIfError(err)
	type args struct {
		ctx  context.Context
		user domain.User
	}
	tests := []struct {
		name           string
		userRepository UserRepository
		args           args
		want           domain.User
		wantErr        bool
	}{
		{
			name: "TEST GET USER BY ID",
			userRepository: NewUserRepositoryImpl(),
			args: args{
				ctx: context.Background(),
				user: domain.User{
					Id: 1,
				},
			},
			want: domain.User{
				Id: 1,
				Name: "Budi",
				Email: "budi@example.com",
				Password: "secretpassword",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := db.Begin()
			helpers.PanicIfError(err)
			user, err := tt.userRepository.FindById(tt.args.ctx, tx, tt.args.user)
			helpers.PanicIfError(err)
			err = tx.Commit()
			helpers.PanicIfError(err)
			assert.Equal(t, tt.want, user)
		})
	}
}

func TestUserRepositoryImpl_FindAll(t *testing.T) {
	db := test.ConnectToDB()
	defer db.Close()
	TruncateUser(db)
	//make user data 
	users := []domain.User{
		{
			Id: 1,
			Name: "Budi",
			Email: "budi@example.com",
			Password: "secretpassword",
		},
		{
			Id: 2,
			Name: "John",
			Email: "john@example.com",
			Password: "secretpassword",
		},
	}
	for _, user := range users {
		userRepository := NewUserRepositoryImpl()
		tx, err := db.Begin()
		helpers.PanicIfError(err)
		userRepository.Save(context.Background(), tx, user)
		err = tx.Commit()
		helpers.PanicIfError(err)
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name           string
		userRepository UserRepository
		args           args
		want           []domain.User
	}{
		{
			name: "GET ALL USER",
			userRepository: NewUserRepositoryImpl(),
			args: args{
				ctx: context.Background(),
			},
			want: users,
		},
	}
	for _, tt := range tests {
		tx, err := db.Begin()
		helpers.PanicIfError(err)
		users := tt.userRepository.FindAll(tt.args.ctx, tx)
		helpers.PanicIfError(err)
		err = tx.Commit()
		helpers.PanicIfError(err)
		assert.Equal(t, tt.want, users)
		assert.Equal(t, 2, len(users))
	}
}


func TestUserRepositoryImpl_FindByEmail(t *testing.T) {
	db := test.ConnectToDB()
	defer db.Close()
	TruncateUser(db)
	//make user data 
	userRepository := NewUserRepositoryImpl()
	tx, err := db.Begin()
	helpers.PanicIfError(err)
	userRepository.Save(context.Background(), tx, domain.User{
		Name: "Budi",
		Email: "budi@example.com",
		Password: "secretpassword",
	})
	err = tx.Commit()
	helpers.PanicIfError(err)
	type args struct {
		ctx  context.Context
		user domain.User
	}
	tests := []struct {
		name           string
		userRepository UserRepository
		args           args
		want           domain.User
		wantErr        bool
	}{
		{
			name: "TEST GET USER BY ID",
			userRepository: NewUserRepositoryImpl(),
			args: args{
				ctx: context.Background(),
				user: domain.User{
					Email: "budi@example.com",
				},
			},
			want: domain.User{
				Id: 1,
				Name: "Budi",
				Email: "budi@example.com",
				Password: "secretpassword",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := db.Begin()
			helpers.PanicIfError(err)
			user, err := tt.userRepository.FindByEmail(tt.args.ctx, tx, tt.args.user)
			helpers.PanicIfError(err)
			err = tx.Commit()
			helpers.PanicIfError(err)
			assert.Equal(t, tt.want, user)
		})
	}
}