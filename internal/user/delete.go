package user

import (
	"strconv"

	"github.com/flambra/account/internal/domain"
	"github.com/flambra/account/internal/profile"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

func Delete(c *fiber.Ctx) error {
	rawId := c.Params("id")
	if rawId == "" {
		return hResp.BadRequestResponse(c, "inform id")
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	var user domain.User
	userRepo := hRepository.New(hDb.Get(), &user, c)

	err = userRepo.Delete(id)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	err = userRepo.GetDeleted(fiber.Map{"id": id})
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	err = profile.Delete(domain.Profile{UserID: user.ID}, c)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, &user)
}
