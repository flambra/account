package user

import (
	"strconv"

	"github.com/flambra/account/database"
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/http"
	"github.com/flambra/helpers/repository"
	"github.com/gofiber/fiber/v2"
)

func Read(c *fiber.Ctx) error {
	rawId := c.Params("id")
	if rawId == "" {
		return http.BadRequestResponse(c, "inform id")
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		return http.BadRequestResponse(c, err.Error())
	}

	var user domain.User
	userRepo := repository.New(database.GetDB(), &user, c)

	userRepo.GetById(id)
	if err != nil {
		return http.InternalServerErrorResponse(c, err.Error())
	}

	return http.SuccessResponse(c, &user)
}
