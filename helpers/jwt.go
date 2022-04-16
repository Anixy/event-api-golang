package helpers

import (
	"errors"
	"time"

	"github.com/Anixy/event-api-golang/model/domain"
	"github.com/golang-jwt/jwt"
)

func CreateJwtToken(user domain.User) string {
	mySigningKey := []byte("secretkey")

	type MyCustomClaims struct {
		User domain.User
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		domain.User{
			Name: user.Name,
			Email: user.Email,
		},
		jwt.StandardClaims{
			ExpiresAt: int64(15*time.Minute),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	PanicIfError(err)
	return ss
}

func VerifyJwtToken(tokenString string) error {
	// Token from another example.  This token is expired

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretkey"), nil
	})

	if token.Valid {
		return nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return errors.New("that's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return errors.New("timing is everything")
		} else {
			return err
		}
	} else {
		return err
	}
}