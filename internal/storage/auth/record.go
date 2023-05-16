package auth

import "time"

type TokenKind string

var (
	TokenKindUser      TokenKind = "USER"
	TokenKindForgotPwd TokenKind = "FORGOT_PASSWORD"
	TokenKindVerifyPwd TokenKind = "VERIFY_PASSWORD"
)

type Record struct {
	UserID    string     `db:"user_id"`
	Token     string     `db:"token"`
	Blacklist bool       `db:"blacklist"`
	Kind      TokenKind  `db:"kind"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
