package user

import (
	"time"

	"github.com/darchlabs/backoffice/internal/api/context"
	"github.com/darchlabs/backoffice/internal/storage/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type PostLoginHandler struct {
	secretKey              string
	userSelectByEmailQuery userSelectByEmailQuery
	authUpsertQuery        authUpsertQuery
}

type postLoginHandlerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type postLoginHandlerResponse struct {
	Token string `json:"token"`
}

type customClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (h *PostLoginHandler) Invoke(ctx *context.Ctx, c *fiber.Ctx) (interface{}, int, error) {
	var req postLoginHandlerRequest
	err := c.BodyParser(&req)
	if err != nil {
		return nil, fiber.StatusInternalServerError, errors.Wrap(
			err, "api: PostLoginHandler.Invoke c.BodyParser error",
		)
	}

	payload, status, err := h.invoke(ctx, &req)
	if err != nil {
		return nil, status, errors.Wrap(err, "user: PostLoginHandler.Invoke h.invoke error")
	}

	return payload, status, nil
}

func (h *PostLoginHandler) invoke(ctx *context.Ctx, req *postLoginHandlerRequest) (interface{}, int, error) {
	user, err := h.userSelectByEmailQuery(ctx.SqlStore, req.Email)
	if err != nil {
		return nil, fiber.StatusInternalServerError, errors.Wrap(err, "user: PostLoginHandler.invoke h.userSelectByEmailAndPwdQuery error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password))
	if err != nil {
		return nil, fiber.StatusUnauthorized, errors.Wrap(err, "user: PostLoginHandler.invoke bcrypt.CompareHashAndPassword error")
	}

	claims := customClaims{
		Email: req.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * 365 * time.Hour).Unix(), // TODO: use better token valid interval
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(h.secretKey))
	if err != nil {
		return nil, fiber.StatusInternalServerError, errors.Wrap(err, "auth: PostLoginHandler.invoke token.SignedString error")
	}

	err = h.authUpsertQuery(ctx.SqlStore, &auth.Record{
		UserID:    user.ID,
		Token:     signedToken,
		Blacklist: false,
		Kind:      auth.TokenKindUser,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, fiber.StatusInternalServerError, errors.Wrap(err, "auth: PostLoginHandler.invoke h.authUpsertQuery error")
	}

	return &postLoginHandlerResponse{Token: signedToken}, fiber.StatusCreated, nil
}
