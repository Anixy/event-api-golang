package helpers

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/Anixy/event-api-golang/model/domain"
	"github.com/golang-jwt/jwt"
)

type MyCustomClaims struct {
	User domain.User	`json:"user"`
	jwt.StandardClaims
}

func CreateJwtToken(user domain.User) string {
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
	ss, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	PanicIfError(err)
	return ss
}

func VerifyJwtToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func GetJwtClaim(tokenString string) (domain.User, error) {
	user := domain.User{}
	claims := MyCustomClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
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

func CreateRefreshToken(token string) (string, error) {
	expired := time.Now().Add(time.Hour * 24).Unix()
	claims := jwt.MapClaims{
		"token": token,
		"exp": expired,
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := refreshToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return token, err
	}
	return token, nil
}

func ValidateRefreshToken(refreshToken string) (domain.User, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	user := domain.User{}
	if err != nil {
		return user, err
	}
	payload, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return user, errors.New("invalid token")
	}
	claims := jwt.MapClaims{}
	parser := jwt.Parser{}
	token, _, err = parser.ParseUnverified(payload["token"].(string), claims)
	if err != nil {
		return user, err
	}
	payload, ok = token.Claims.(jwt.MapClaims)
	if !ok {
		return user, errors.New("invalid token")
	}
	bytes, err :=json.Marshal(payload["user"])
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}