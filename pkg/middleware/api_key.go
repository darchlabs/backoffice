package middleware

import (
	"github.com/darchlabs/backoffice/pkg/client"
	"github.com/gofiber/fiber/v2"
)

type ApiKey struct {
	client *client.Client
}

func NewApiKey(cl *client.Client) *Auth {
	return &Auth{client: cl}
}

func (a *ApiKey) Middleware(c *fiber.Ctx) error {
	// Extract the Authorization header
	apiKey := c.Get("X-Api-Key")

	// Check if the Authorization header is present and in the correct format
	if apiKey == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	// Call the ValidTokenWithCtx function to validate the token and get the user ID
	response, err := a.client.ValidApiKeyWithCtx(c.Context(), apiKey)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	// Attach the user ID to the request context
	c.Locals("user_id", response.UserID)

	// Continue to the next middleware or route handler
	return c.Next()
}
