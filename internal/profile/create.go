package profile

import (
	"github.com/flambra/account/database"
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/http"
	"github.com/flambra/helpers/repository"
	"github.com/gofiber/fiber/v2"
)

func Create(profile domain.Profile, c *fiber.Ctx) error {
	profileRepo := repository.New(database.GetDB(), &profile, c)

	err := profileRepo.Create()
	if err != nil {
		return http.InternalServerErrorResponse(c, err.Error())
	}

	return nil
}
