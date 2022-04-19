package web

type ParticipantResponse struct {
	User  UserResponse  `json:"user"`
	Event EventResponse `json:"event"`
}