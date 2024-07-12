package middleware

import (
	"os"
	"strings"

	"github.com/flambra/helpers/hResp"
	"github.com/flambra/helpers/hToken"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	if os.Getenv("AUTH_MIDDLEWARE") == "disable" {
		return c.Next()
	}

	token := c.Get("Authorization")
	if token != "" {
		parts := strings.Split(token, " ")
		if len(parts) != 2 {
			return hResp.BadRequestResponse(c, "Token error")
		}
		scheme := parts[0]
		token = parts[1]
		if !strings.EqualFold(scheme, "Bearer") {
			return hResp.BadRequestResponse(c, "Token malformatted")
		}
	} else {
		var request hToken.Access

		if err := c.BodyParser(&request); err != nil {
			return hResp.BadRequestResponse(c, err.Error())
		}
		token = request.Token
	}

	claims, err := hToken.Parse(token)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	message, ok := claims["message"].(string)
	if !ok || message != "Token for Create User" {
		return hResp.UnauthorizedResponse(c, "Invalid token message")
	}

	return c.Next()
}
