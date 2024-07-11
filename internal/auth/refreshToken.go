package auth

import (
	"strings"

	"github.com/flambra/helpers/hResp"
	"github.com/flambra/helpers/hToken"
	"github.com/gofiber/fiber/v2"
)

func RefreshToken(c *fiber.Ctx) error {
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
		token = request.RefreshToken
	}

	err := hToken.Validate(token)
	if err != nil {
		return hResp.UnauthorizedResponse(c, err.Error())
	}

	data, err := hToken.Parse(token)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	access, err := hToken.Create(data)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	response := hToken.Access{
		Token:        access.Token,
		RefreshToken: access.RefreshToken,
	}

	return hResp.SuccessResponse(c, &response)
}
