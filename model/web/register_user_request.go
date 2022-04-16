package web

type RegisterUserRequest struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	VerifyPassword string `json:"verify_password"`
}