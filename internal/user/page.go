package user

import (
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

func Page(c *fiber.Ctx) error {
	var response []domain.UserPageResponse
	paginator := hRepository.BuildPaginator(&response)

	var user []domain.User
	repo := hRepository.New(hDb.Get(), &user, c)

	err := c.QueryParser(paginator)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	var filter domain.UserPageFilter
	err = c.QueryParser(&filter)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err = repo.FindAllPaginating(&filter, paginator)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, paginator)
}
