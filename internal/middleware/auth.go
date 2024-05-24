package middleware

import (
	"os"

	"github.com/flambra/helpers/hAuth"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

// Auth creates a middleware to hValidate the authentication token.
func Auth(c *fiber.Ctx) error {
	if os.Getenv("AUTH_MIDDLEWARE") == "disable" {
		return c.Next()
	}

	token := c.Get("Authorization")

	err := hAuth.ValidateToken(token)
	if err != nil {
		return hResp.UnauthorizedResponse(c, err.Error())
	}

	return c.Next()
}
