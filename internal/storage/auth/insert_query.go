package auth

import (
	"github.com/darchlabs/backoffice/internal/storage"
	"github.com/pkg/errors"
)

func InsertQuery(tx storage.QueryContext, record *Record) error {
	_, err := tx.Exec(`
		INSERT INTO auth (user_id, token, blacklist, kind, created_at)
		VALUES ($1, $2, $3, $4, %5);`,
		record.UserID,
		record.Token,
		record.Blacklist,
		record.Kind,
		record.CreatedAt,
	)
	if err != nil {
		return errors.Wrap(err, "auth: InsertQuery tx.Exec error")
	}

	return nil
}
