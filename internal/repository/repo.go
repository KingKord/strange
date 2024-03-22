package repository

import (
	"github.com/KingKord/strange/internal/model"
	"time"
)

type Repository interface {
	AssignMeet(card model.Card) error
	DaySchedule(day time.Time, userID int) ([]model.Card, error)
}
