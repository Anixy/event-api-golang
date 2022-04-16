package helpers

import (
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
