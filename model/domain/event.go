package domain

import "time"

type Event struct {
	Id        	int
	User      	User
	Title     	string
	StartDate 	time.Time
	EndDate 	time.Time
	Description string
	Type 		string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}