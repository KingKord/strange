package test_repo

import (
	"fmt"
	"github.com/KingKord/strange/internal/model"
	"github.com/KingKord/strange/internal/repository"
	"time"
)

type testRepo struct{}

func NewTestRepo() repository.Repository {
	return testRepo{}
}
func (testRepo) AssignMeet(card model.Card) error {
	if card.Name == "DB err" {
		return fmt.Errorf("DB err")
	}

	return nil
}

func (testRepo) DaySchedule(day time.Time, userID int) ([]model.Card, error) {
	if userID == 1234 {
		return []model.Card{
			{
				Name:        "super meeting",
				UserID:      1234,
				Description: "discussion about importance of high education",
			},
		}, nil
	}
	if userID == 4321 {
		return nil, fmt.Errorf("DB err")
	}
	return nil, nil
}
