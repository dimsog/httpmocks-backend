package repository

import (
	"context"
	"github.com/dimsog/httpmocks-backend/internal/types"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func Create(mock types.Mock, conn *pgx.Conn, log *slog.Logger) {
	id, _ := uuid.NewV7()

	_, err := conn.Exec(context.Background(), `
		INSERT INTO mocks (id, http_code, content_type, body, created_at) VALUES(
			$1, $2, $3, $4, now()
		)
	`, id.String(), 200, mock.ContentType, mock.Body)

	if err != nil {
		log.Error(err.Error())
	}
}
