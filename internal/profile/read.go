package profile

import (
	"strconv"

	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

func Read(c *fiber.Ctx) error {
	rawId := c.Params("id")
	if rawId == "" {
		return hResp.BadRequestResponse(c, "inform id")
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	var profile domain.Profile
	repoprofile := hRepository.New(hDb.Get(), &profile, c)

	err = repoprofile.GetById(id)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, &profile)
}
