package repository

import (
	"time"

	"github.com/KingKord/strange/internal/model"
)

type Repository interface {
	AssignMeet(card model.Card) error
	DaySchedule(day time.Time, userID int) ([]model.Card, error)
}
