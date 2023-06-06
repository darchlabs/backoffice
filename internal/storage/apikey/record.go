package apikey

import "time"

type Record struct {
	ID        string     `db:"id"`
	UserID    string     `db:"user_id"`
	Token     string     `db:"token"`
	ExpiresAt *time.Time `db:"expires_at"`
	CreatedAt *time.Time `db:"created_at"`
}
