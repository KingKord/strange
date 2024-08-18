package services

import (
	"context"
	"fmt"
	"time"

	"github.com/KingKord/strange/internal/model"
	"github.com/KingKord/strange/internal/repository"
)

type ScheduleService struct {
	repo repository.Repository
}

func NewScheduleService(repo repository.Repository) ScheduleService {
	return ScheduleService{
		repo: repo,
	}
}

func (s ScheduleService) AssignMeet(ctx context.Context, card model.Card) error {
	err := s.repo.AssignMeet(card)
	if err != nil {
		return fmt.Errorf("repo.AssignMeet: %w", err)
	}

	return nil
}

func (s ScheduleService) DaySchedule(ctx context.Context, day time.Time, userID int) ([]model.Card, error) {
	cards, err := s.repo.DaySchedule(day, userID)
	if err != nil {
		return nil, fmt.Errorf("repo.AssignMeet: %w", err)
	}

	return cards, nil
}
