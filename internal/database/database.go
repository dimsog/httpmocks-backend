package database

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func MustConnect(uri string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), uri)

	if err != nil {
		panic(err)
	}

	if err = conn.Ping(context.Background()); err != nil {
		panic(err)
	}

	return conn
}
