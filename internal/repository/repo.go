package repository

import "database/sql"

type Repository interface {
}

type postgresDBRepo struct {
	DB *sql.DB
}

func NewPostgresRepo(conn *sql.DB) Repository {
	return &postgresDBRepo{
		DB: conn,
	}
}
