package profile

import (
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

func Create(profile domain.Profile, c *fiber.Ctx) error {
	repo := hRepository.New(hDb.Get(), &profile, c)

	err := repo.Create()
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return nil
}
