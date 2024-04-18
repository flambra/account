package middleware

import (
	"io"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

// Auth creates a middleware to verify the authentication token.
func Auth(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization token not provided",
		})
	}

	authURL := os.Getenv("AUTH_URL")
	if authURL == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Authentication service URL not configured",
		})
	}

	authReq, err := http.NewRequest("GET", authURL + "/client/verify", nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create request to authentication service",
		})
	}

	authReq.Header.Set("Authorization", token)
	client := &http.Client{}
	res, err := client.Do(authReq)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Failed to verify token",
		})
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
			"body":  string(body),
		})
	}

	return c.Next()
}
