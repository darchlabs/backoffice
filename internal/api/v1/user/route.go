package user

import (
	"fmt"
	"net/http"

	"github.com/darchlabs/backoffice/internal/api/context"
	v1 "github.com/darchlabs/backoffice/internal/api/v1"
	"github.com/darchlabs/backoffice/internal/storage/apikey"
	authdb "github.com/darchlabs/backoffice/internal/storage/auth"
	userdb "github.com/darchlabs/backoffice/internal/storage/user"
	"github.com/darchlabs/backoffice/pkg/client"
	"github.com/darchlabs/backoffice/pkg/middleware"
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

	postValidTokenHandler := &PostValidTokenHandler{
		secretKey:              ctx.App.Config.SecretKey,
		authSelectByTokenQuery: authdb.SelectByTokenQuery,
		userSelectByEmailQuery: userdb.SelectByEmailQuery,
	}

	postApiKeyHandler := &PostApiKeyHandler{
		secretKey:              ctx.App.Config.SecretKey,
		idGenerate:             uuid.NewString,
		apikeyInsertQuery:      apikey.InsertQuery,
		authSelectByTokenQuery: authdb.SelectByTokenQuery,
	}

	postValidApiKeyHandler := &PostValidApiKeyHandler{
		secretKey:                ctx.App.Config.SecretKey,
		apikeySelectByTokenQuery: apikey.SelectByTokenQuery,
	}

	// setup middleware
	cl := client.New(&client.Config{
		Client:  http.DefaultClient,
		BaseURL: fmt.Sprintf("http://0.0.0.0:%s", ctx.App.Config.ApiServerPort),
	})

	auth := middleware.NewAuth(cl)

	// route
	ctx.Server.Post(fmt.Sprintf("%s/signup", basePath), v1.HandleFunc(ctx, postSignupHandler.Invoke))
	ctx.Server.Post(fmt.Sprintf("%s/login", basePath), v1.HandleFunc(ctx, postLoginHandler.Invoke))
	ctx.Server.Post(fmt.Sprintf("%s/tokens", basePath), v1.HandleFunc(ctx, postValidTokenHandler.Invoke))
	ctx.Server.Post(
		fmt.Sprintf("%s/api-key", basePath),
		auth.Middleware,
		v1.HandleFunc(ctx, postApiKeyHandler.Invoke),
	)
	ctx.Server.Post(fmt.Sprintf("%s/valid-api-key", basePath), v1.HandleFunc(ctx, postValidApiKeyHandler.Invoke))
}
