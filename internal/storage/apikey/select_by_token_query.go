package apikey

import (
	"database/sql"

	"github.com/darchlabs/backoffice/internal/storage"
	"github.com/pkg/errors"
)

var (
	ErrNotFound = errors.New("apikey: apikey not found")
)

func SelectByTokenQuery(tx storage.Transaction, token string) (*Record, error) {
	var record Record

	err := tx.Get(&record, `
		SELECT *
		FROM api_keys
		WHERE token = $1;`,
		token,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, errors.Wrap(err, "apikey: SelectByEmailQuery tx.Get error")
	}

	return &record, nil
}
