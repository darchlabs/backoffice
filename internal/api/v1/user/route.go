package user

import (
	"github.com/darchlabs/backoffice/internal/api/context"
	v1 "github.com/darchlabs/backoffice/internal/api/v1"
	userdb "github.com/darchlabs/backoffice/internal/storage/user"
	"github.com/google/uuid"
)

func Route(basePath string, ctx *context.Ctx) {
	// handlers
	postSignupHandler := &PostSignupHandler{
		userInsertQuery: userdb.InsertQuery,
		idGenerate:      uuid.NewString,
	}

	// route
	ctx.Server.Post(basePath, v1.HandleFunc(ctx, postSignupHandler.Invoke))
}
