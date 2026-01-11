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
		return nil, fmt.Errorf("can't connect: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't ping: %w", err)
	}

	return &Storage{
		db: db,
	}, nil
}
