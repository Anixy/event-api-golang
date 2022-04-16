package web

type RegisterUserRequest struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required"`
	VerifyPassword string `json:"verify_password" binding:"required"`
}