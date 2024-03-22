package model

import "time"

// Card struct
type Card struct {
	Id          int
	Name        string
	UserID      int
	Description string
	From        time.Time
	To          time.Time
}
