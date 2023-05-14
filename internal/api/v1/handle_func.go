package v1

import (
	"github.com/darchlabs/backoffice/internal/api/context"
	"github.com/gofiber/fiber/v2"
)

var StatusAlreadyProxied = 1000

type Handler func(*context.Ctx, *fiber.Ctx) (interface{}, int, error)

func HandleFunc(ctx *context.Ctx, fn Handler) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		payload, statusCode, err := fn(ctx, c)
		if err != nil {
			return c.Status(statusCode).JSON(map[string]string{
				"error": err.Error(),
			})
		}

		if statusCode == StatusAlreadyProxied {
			return nil
		}

		return c.Status(statusCode).JSON(payload)
	}
}

func HealthRoute(ctx *context.Ctx) {
	ctx.Server.Get("/api/v1/health", HandleFunc(
		ctx,
		func(_ *context.Ctx, _ *fiber.Ctx) (interface{}, int, error) {
			return map[string]string{"status": "running"}, fiber.StatusOK, nil
		}),
	)
}
