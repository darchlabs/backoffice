package user

import (
	"encoding/json"
	"time"

	"github.com/darchlabs/backoffice/internal/api/context"
	authdb "github.com/darchlabs/backoffice/internal/storage/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

type PostValidTokenHandler struct {
	secretKey              string
	authSelectByTokenQuery authSelectByTokenQuery
	userSelectByEmailQuery userSelectByEmailQuery
}

type PostValidTokenHandlerRequest struct {
	Token string `json:"token"`
}

type PostValidTokenHandlerResponse struct {
	UserID string `json:"user_id"`
}

func (h *PostValidTokenHandler) Invoke(ctx *context.Ctx, c *fiber.Ctx) (interface{}, int, error) {
	var req PostValidTokenHandlerRequest
	if err := json.Unmarshal(c.Body(), &req); err != nil {
		return nil, fiber.StatusBadRequest, errors.Wrap(err, "user: PostValidTokenHandler.Invoke json.Unmarshal error")
	}

	payload, status, err := h.invoke(ctx, &req)
	if err != nil {
		return nil, status, errors.Wrap(err, "user: PostValidTokenHandler.Invoke h.invoke error")
	}
	return payload, status, nil
}

func (h *PostValidTokenHandler) invoke(ctx *context.Ctx, req *PostValidTokenHandlerRequest) (interface{}, int, error) {
	// Validate the token
	auth, err := h.authSelectByTokenQuery(ctx.SqlStore, req.Token)
	if errors.Is(err, authdb.ErrNotFound) {
		return nil, fiber.StatusUnauthorized, errors.Wrap(err, "user: h.authSelectByTokenQuery error")
	}
	if err != nil {
		return nil, fiber.StatusInternalServerError, errors.Wrap(err, "h.selectByTokenQuery error")
	}

	claims, err := h.parseToken(auth.Token)
	if err != nil {
		return nil, fiber.StatusInternalServerError, errors.Wrap(err, "h.parseToken error")
	}

	userRecord, err := h.userSelectByEmailQuery(ctx.SqlStore, claims.Email)
	if err != nil {
		return nil, fiber.StatusInternalServerError, errors.Wrap(err, "h.userSelectByEmailQuery error")
	}

	// Perform additional checks on the claims, if necessary

	return &PostValidTokenHandlerResponse{UserID: userRecord.ID}, fiber.StatusOK, nil
}

func (h *PostValidTokenHandler) parseToken(tokenString string) (*customClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(h.secretKey), nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "jwt.ParseWithClaims error")
	}

	claims, ok := token.Claims.(*customClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Verify that the token is not expired
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expired")
	}

	return claims, nil
}
