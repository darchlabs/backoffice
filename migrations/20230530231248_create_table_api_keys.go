package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateTableApiKeys, downCreateTableApiKeys)
}

func upCreateTableApiKeys(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
		CREATE TABLE api_keys (
			id TEXT PRIMARY KEY NOT NULL,
			user_id TEXT NOT NULL REFERENCES users(id),
			token TEXT NOT NULL UNIQUE,
			expires_at TIMESTAMPTZ NOT NULL,
			created_at TIMESTAMPTZ NOT NULL
		);`)
	if err != nil {
		return err
	}

	return nil
}

func downCreateTableApiKeys(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
