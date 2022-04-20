package web

type LoginUserResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}