package services

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/KingKord/strange/internal/model"
	"github.com/KingKord/strange/internal/repository"
	"github.com/KingKord/strange/internal/repository/postgres"
	"github.com/KingKord/strange/internal/repository/test_repo"
)

func TestNewScheduleService(t *testing.T) {
	type args struct {
		repo repository.Repository
	}
	repo := postgres.NewPostgresRepo(nil)
	tests := []struct {
		name string
		args args
		want ScheduleService
	}{
		{
			name: "happy path",
			args: args{
				repo: repo,
			},
			want: ScheduleService{
				repo: repo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScheduleService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScheduleService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScheduleService_AssignMeet(t *testing.T) {
	type fields struct {
		repo repository.Repository
	}
	type args struct {
		ctx  context.Context
		card model.Card
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "happy path",
			fields: fields{
				repo: test_repo.NewTestRepo(),
			},
			args: args{
				ctx:  context.Background(),
				card: model.Card{},
			},
			wantErr: false,
		},
		{
			name:   "got DB err",
			fields: fields{repo: test_repo.NewTestRepo()},
			args: args{
				ctx: context.Background(),
				card: model.Card{
					Name: "DB err",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ScheduleService{
				repo: tt.fields.repo,
			}
			if err := s.AssignMeet(tt.args.ctx, tt.args.card); (err != nil) != tt.wantErr {
				t.Errorf("AssignMeet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestScheduleService_DaySchedule(t *testing.T) {
	type fields struct {
		repo repository.Repository
	}
	type args struct {
		ctx    context.Context
		day    time.Time
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Card
		wantErr bool
	}{
		{
			name:   "happy path",
			fields: fields{repo: test_repo.NewTestRepo()},
			args: args{
				ctx:    context.Background(),
				userID: 1234,
			},
			want: []model.Card{
				{
					Name:        "super meeting",
					UserID:      1234,
					Description: "discussion about importance of high education",
				},
			},
			wantErr: false,
		},
		{
			name:   "DB err",
			fields: fields{repo: test_repo.NewTestRepo()},
			args: args{
				ctx:    context.Background(),
				userID: 4321,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ScheduleService{
				repo: tt.fields.repo,
			}
			got, err := s.DaySchedule(tt.args.ctx, tt.args.day, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DaySchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DaySchedule() got = %v, want %v", got, tt.want)
			}
		})
	}
}
