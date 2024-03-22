package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/KingKord/strange/internal/model"
	"github.com/KingKord/strange/internal/repository"
	"time"
)

const dbTimeout = time.Second * 3

type postgresDBRepo struct {
	DB *sql.DB
}

func NewPostgresRepo(conn *sql.DB) repository.Repository {
	return &postgresDBRepo{
		DB: conn,
	}
}

func (p postgresDBRepo) AssignMeet(card model.Card) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	checkQuery := `select count(*)
from schedule
where (
          ($1 between date_from and date_to)
              or ($2 between date_from and date_to)
              or ($1 < date_from and date_to < $2)
    )`
	var count int

	err := p.DB.QueryRowContext(ctx, checkQuery,
		time.Now(),
		time.Now(),
	).Scan(&count)
	if err != nil {
		return fmt.Errorf("DB.QueryRowContext: %w", err)
	}
	if count != 0 {
		return fmt.Errorf("couldn't assign meet due schedule conflict")
	}

	query := `insert into schedule (user_id, name, description, date_from, date_to)
				values ($1, $2, $3, $4, $5)`

	_, err = p.DB.QueryContext(ctx, query,
		card.UserID,
		card.Name,
		card.Description,
		card.From,
		card.To,
	)
	if err != nil {
		return err
	}

	return nil
}
