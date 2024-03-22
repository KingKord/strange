package repository

import (
	"github.com/KingKord/strange/internal/model"
)

type Repository interface {
	AssignMeet(card model.Card) error
}
