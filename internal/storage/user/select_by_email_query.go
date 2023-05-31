package user

import (
	"database/sql"

	"github.com/darchlabs/backoffice/internal/storage"
	"github.com/pkg/errors"
)

var (
	ErrNotFound = errors.New("user: user not found")
)

func SelectByEmailQuery(tx storage.Transaction, email string) (*Record, error) {
	var record Record
	err := tx.Get(&record, `
		SELECT *
		FROM users
		WHERE email = $1;`,
		email,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, errors.Wrap(err, "user: SelectByEmailQuery tx.Get error")
	}
	return &record, nil
}
