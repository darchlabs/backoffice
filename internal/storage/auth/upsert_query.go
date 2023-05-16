package auth

import (
	"github.com/darchlabs/backoffice/internal/storage"
	"github.com/pkg/errors"
)

func UpsertQuery(tx storage.QueryContext, record *Record) error {
	_, err := tx.Exec(`
		INSERT INTO auth (user_id, token, blacklist, kind, created_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (user_id)
		DO UPDATE
		SET token = $2
		WHERE auth.user_id = $1;`,
		record.UserID,
		record.Token,
		record.Blacklist,
		record.Kind,
		record.CreatedAt,
	)
	if err != nil {
		return errors.Wrap(err, "auth: UpsertQuery tx.Exec error")
	}
	return nil
}
