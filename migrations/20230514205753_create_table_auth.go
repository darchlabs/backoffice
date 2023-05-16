package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateTableAuth, downCreateTableAuth)
}

func upCreateTableAuth(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`CREATE TYPE token_kind AS ENUM ('USER', 'FORGOT_PASSWORD', 'VERIFY_PASSWORD');`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		CREATE TABLE auth (
			user_id TEXT PRIMARY KEY NOT NULL REFERENCES users(id),
			token TEXT NOT NULL,
			blacklist BOOLEAN NOT NULL,
			kind TOKEN_KIND NOT NULL DEFAULT 'USER',
			created_at TIMESTAMPTZ NOT NULL,
			updated_at TIMESTAMPTZ,
			deleted_at TIMESTAMPTZ
		);`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateTableAuth(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE auth CASCADE;`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`DROP TYPE token_kind;`)
	return nil
}
