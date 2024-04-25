package middleware

import (
	token "github.com/flambra/helpers/access"
	"github.com/flambra/helpers/http"
	"github.com/gofiber/fiber/v2"
)

// Auth creates a middleware to validate the authentication token.
func Auth(c *fiber.Ctx) error {
	tokenAuthorizer := c.Get("Authorization")

	err := token.ValidateToken(tokenAuthorizer)
	if err != nil {
		return http.UnauthorizedResponse(c, err.Error())
	}

	return c.Next()
}
