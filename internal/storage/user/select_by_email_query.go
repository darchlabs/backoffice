package user

import (
	"github.com/darchlabs/backoffice/internal/storage"
	"github.com/pkg/errors"
)

func SelectByEmailQuery(tx storage.Transaction, email string) (*Record, error) {
	var record Record
	err := tx.Get(&record, `
		SELECT *
		FROM users
		WHERE email = $1;`,
		email,
	)
	if err != nil {
		return nil, errors.Wrap(err, "user: SelectByEmailQuery tx.Get error")
	}
	return &record, nil
}
