package auth

import (
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hResp"
	"github.com/flambra/helpers/hToken"
	"github.com/gofiber/fiber/v2"
)

func RefreshToken(c *fiber.Ctx) error {
	var request domain.AuthRefreshTokenRequest

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err := hToken.Validate(request.RefreshToken)
	if err != nil {
		return hResp.UnauthorizedResponse(c, err.Error())
	}

	data, err := hToken.Parse(request.RefreshToken)
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
