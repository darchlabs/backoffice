package user

import (
	"time"

	"github.com/darchlabs/backoffice/internal/api/context"
	"github.com/darchlabs/backoffice/internal/storage"
	userdb "github.com/darchlabs/backoffice/internal/storage/user"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type userInsertQuery func(storage.QueryContext, *userdb.Record) error

type idGenerate func() string

type PostSignupHandler struct {
	idGenerate      idGenerate
	userInsertQuery userInsertQuery
}

type postSignupHandlerRequest struct {
	Email    string `db:"email"`
	Name     string `db:"name"`
	Password string `db:"password"`
}

type PostSignupHandlerResponse struct {
}

func (h *PostSignupHandler) Invoke(ctx *context.Ctx, c *fiber.Ctx) (interface{}, int, error) {
	var req postSignupHandlerRequest
	err := c.BodyParser(&req)
	if err != nil {
		return nil, fiber.StatusInternalServerError, errors.Wrap(
			err, "api: PostSignupHandler.Invoke c.BodyParser error",
		)
	}

	payload, status, err := h.invoke(ctx, &req)
	if err != nil {
		return nil, status, errors.Wrap(err, "user: PostSignupHandler.Invoke h.invoke error")
	}
	return payload, status, nil
}

func (h *PostSignupHandler) invoke(ctx *context.Ctx, req *postSignupHandlerRequest) (interface{}, int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fiber.StatusInternalServerError, errors.Wrap(
			err, "user: PostSignupHandler.invoke bcrypt.GenerateFromPassword error",
		)
	}

	// Create *userdb.Record
	record := &userdb.Record{
		ID:             h.idGenerate(),
		Email:          req.Email,
		Name:           req.Name,
		HashedPassword: string(hashedPassword),
		Verified:       false,
		CreatedAt:      time.Now(),
	}

	err = h.userInsertQuery(ctx.SqlStore, record)
	if err != nil {
		return nil, fiber.StatusInternalServerError, errors.Wrap(
			err, "user: PostSignupHandler.invoke h.userInsertQuery error",
		)
	}

	return nil, fiber.StatusCreated, nil
}
