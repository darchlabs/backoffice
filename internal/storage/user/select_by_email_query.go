package user

import (
	"github.com/darchlabs/backoffice/internal/storage"
	"github.com/pkg/errors"
)

type SelectByEmailQueryData struct {
	Email string
}

func SelectByEmailQuery(tx storage.Transaction, data *SelectByEmailQueryData) (*Record, error) {
	var record Record
	err := tx.Get(&record, `
		SELECT *
		FROM users
		WHERE email = $1;`,
		data.Email,
	)
	if err != nil {
		return nil, errors.Wrap(err, "user: SelectByEmailQuery tx.Get error")
	}
	return &record, nil
}
