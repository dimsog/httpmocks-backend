package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5"
)

type Storage struct {
	db *sql.DB
}

func New(uri string) (*Storage, error) {
	db, err := sql.Open("postgres", uri)

	if err != nil {
		return nil, fmt.Errorf("can't open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to db: %w", err)
	}

	return &Storage{
		db: db,
	}, nil
}
