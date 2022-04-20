package web

type EventParticipantResponse struct {
	Event        EventResponse     `json:"event"`
	Participants []UserParticipant `json:"participants"`
}

type UserParticipant struct {
	Id            int    `json:"id"`
	ParticipantId int    `json:"participant_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
}