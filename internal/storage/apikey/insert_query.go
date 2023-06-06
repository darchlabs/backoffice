package apikey

import (
	"github.com/darchlabs/backoffice/internal/storage"
	"github.com/pkg/errors"
)

func InsertQuery(tx storage.QueryContext, record *Record) error {
	_, err := tx.Exec(`
		INSERT INTO api_keys (id, user_id, token, expires_at, created_at)
		VALUES ($1, $2, $3, $4, $5);`,
		record.ID,
		record.UserID,
		record.Token,
		record.ExpiresAt,
		record.CreatedAt,
	)
	if err != nil {
		return errors.Wrap(err, "apikey: InsertQuery tx.Exec error")
	}

	return nil
}
