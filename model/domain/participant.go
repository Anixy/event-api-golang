package domain

import "time"

type Participant struct {
	Id        int
	Event     Event
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}