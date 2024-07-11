package twoFactor

import (
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/flambra/helpers/hToken"
	"github.com/gofiber/fiber/v2"
)

func Validate(c *fiber.Ctx) error {
	var request domain.AuthTwoFactorValidateRequest

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	claims, err := hToken.Parse(request.Token)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return hResp.InternalServerErrorResponse(c, "UserID not found in token data")
	}

	var user domain.User
	repo := hRepository.New(hDb.Get(), &user, nil)
	err = repo.GetById(int(userID))
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	if user.LastCode != request.Code {
		return hResp.UnauthorizedResponse(c, "Invalid 2FA code")
	}

	return hResp.SuccessResponse(c, fiber.Map{
		"message": "2FA code validated successfully",
	})
}
