package helpers

import (
	"fmt"
	"testing"

	"github.com/Anixy/event-api-golang/model/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateJwtToken(t *testing.T) {
	type args struct {
		user domain.User
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TEST CREATE TOKEN",
			args: args{
				user: domain.User{
					Id:       1,
					Name:     "Budi",
					Email:    "budi@example.com",
					Password: "secretpassword",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := CreateJwtToken(tt.args.user)
			assert.NotZero(t, len(token))
		})
	}
}

func TestVerifyJwtToken(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TEST VALID JWT TOKEN",
			args: args{
				tokenString: CreateJwtToken(domain.User{
					Id:       1,
					Name:     "Budi",
					Email:    "budi@example.com",
					Password: "secretpassword",
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := VerifyJwtToken(tt.args.tokenString)
			assert.NoError(t, err)
		})
	}
}

func TestGetJwtClaim(t *testing.T) {
	user := domain.User{
		Id:       1,
		Name:     "Budi",
		Email:    "budi@example.com",
		Password: "secretpassword",
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.User
		wantErr bool
	}{
		{
			name: "TEST GET JWT CLAIMS",
			args: args{
				tokenString: CreateJwtToken(user),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRes, err := GetJwtClaim(tt.args.tokenString)
			assert.NoError(t, err)
			assert.Equal(t, user.Id, userRes.Id)
			assert.Equal(t, user.Name, userRes.Name)
			assert.Equal(t, user.Email, userRes.Email)
		})
	}
}

func TestValidateRefreshToken(t *testing.T) {
	user := domain.User{
		Id:       1,
		Name:     "Budi",
		Email:    "budi@example.com",
		Password: "secretpassword",
	}
	token := CreateJwtToken(user)
	err := VerifyJwtToken(token)
	if err != nil {
		fmt.Println(err.Error())
	}
	user, err = GetJwtClaim(token)
	if err != nil {
		fmt.Println(err.Error())
	}
	refreshToken, err := CreateRefreshToken(token)
	assert.Nil(t, err)
	user, err = ValidateRefreshToken(refreshToken)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(user)
	// fmt.Println(err.Error())	
}
