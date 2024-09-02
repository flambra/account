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

	var users []domain.User
	repo := hRepository.New(hDb.Get(), &users, c)

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

	// Transforme `users` para `response`
	response = make([]domain.UserPageResponse, len(users))
	for i, user := range users {
		response[i] = domain.UserPageResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Address:   user.Address,
			UserType:  user.UserType,
		}
	}

	paginator.Rows = response
	return hResp.SuccessResponse(c, paginator)
}
