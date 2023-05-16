package user

import (
	"fmt"

	"github.com/darchlabs/backoffice/internal/api/context"
	v1 "github.com/darchlabs/backoffice/internal/api/v1"
	authdb "github.com/darchlabs/backoffice/internal/storage/auth"
	userdb "github.com/darchlabs/backoffice/internal/storage/user"
	"github.com/google/uuid"
)

func Route(basePath string, ctx *context.Ctx) {
	// handlers
	postSignupHandler := &PostSignupHandler{
		userInsertQuery: userdb.InsertQuery,
		idGenerate:      uuid.NewString,
	}

	postLoginHandler := &PostLoginHandler{
		secretKey:              ctx.App.Config.SecretKey,
		userSelectByEmailQuery: userdb.SelectByEmailQuery,
		authUpsertQuery:        authdb.UpsertQuery,
	}

	// route
	ctx.Server.Post(fmt.Sprintf("%s/signup", basePath), v1.HandleFunc(ctx, postSignupHandler.Invoke))
	ctx.Server.Post(fmt.Sprintf("%s/login", basePath), v1.HandleFunc(ctx, postLoginHandler.Invoke))
}
