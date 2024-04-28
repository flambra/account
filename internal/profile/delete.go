package profile

import (
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

func Delete(profile domain.Profile, c *fiber.Ctx) error {
	profileRepo := hRepository.New(hDb.Get(), &profile, c)

	err := profileRepo.Delete(int(profile.UserID))
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	err = profileRepo.GetDeleted(fiber.Map{"user_id": int(profile.UserID)})
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return nil
}
