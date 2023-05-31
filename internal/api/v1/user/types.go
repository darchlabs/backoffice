package user

import (
	"github.com/darchlabs/backoffice/internal/storage"
	"github.com/darchlabs/backoffice/internal/storage/auth"
	"github.com/darchlabs/backoffice/internal/storage/user"
	userdb "github.com/darchlabs/backoffice/internal/storage/user"
)

type idGenerate func() string

type userInsertQuery func(storage.QueryContext, *userdb.Record) error

type userSelectByEmailQuery func(storage.Transaction, string) (*user.Record, error)

type authInsertQuery func(storage.QueryContext, *auth.Record) error

type authUpsertQuery func(storage.QueryContext, *auth.Record) error

type authSelectByTokenQuery func(storage.Transaction, string) (*auth.Record, error)
