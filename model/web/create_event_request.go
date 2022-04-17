package web

import (
	"encoding/json"
	"time"
)

type myTime time.Time

var _ json.Unmarshaler = &myTime{}

func (mt *myTime) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}
	*mt = myTime(t)
	return nil
}

type CreateEventRequest struct {
	Title     	string 		`json:"title" binding:"required"`
	StartDate 	time.Time	`json:"start_date" binding:"required"`
	EndDate 	time.Time	`json:"end_date" binding:"required"`
	Description string		`json:"description" binding:"required"`
	Type 		string 		`json:"type" binding:"required,oneof=online offline"`
}