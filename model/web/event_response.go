package web

import "time"

type EventResponse struct {
	Id        	int				`json:"id"`
	Title     	string			`json:"title"`
	User      	UserResponse	`json:"user"`
	StartDate 	time.Time		`json:"start_date"`
	EndDate		time.Time		`json:"end_date"`
	Description string			`json:"description"`
	Type 		string			`json:"type"`
}