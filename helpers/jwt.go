package helpers

import (
	"errors"
	"strings"
	"time"

	"github.com/Anixy/event-api-golang/model/domain"
	"github.com/golang-jwt/jwt"
)

type MyCustomClaims struct {
	User domain.User
	jwt.StandardClaims
}

func CreateJwtToken(user domain.User) string {
	mySigningKey := []byte("secretkey")

	// Create the Claims
	expiredTime := time.Now().Add(15*time.Minute).Unix()
	claims := MyCustomClaims{
		domain.User{
			Id: user.Id,
			Name: user.Name,
			Email: user.Email,
		},
		jwt.StandardClaims{
			ExpiresAt: expiredTime,
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

func GetJwtClaim(tokenString string) (domain.User, error) {
	user := domain.User{}
	claims := MyCustomClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("secretkey"), nil
	})
	if err != nil {
		return user, nil
	}
	return claims.User, nil
}

func GetJwtTokenFromBearer(bearerToken string) (string, error) {
	splitBearer := strings.Split(bearerToken, " ")
	if len(splitBearer) != 2 || splitBearer[0] != "Bearer" {
		return "", errors.New("not valid bearer token")
	}

	return splitBearer[1], nil
}