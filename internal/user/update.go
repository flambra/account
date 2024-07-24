package user

import (
	"strconv"
	"time"

	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hLog"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

func Update(c *fiber.Ctx) error {
	defer hLog.Performance(time.Now(), "Update")
	rawId := c.Params("id")
	if rawId == "" {
		return hResp.BadRequestResponse(c, "inform id")
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	var user domain.User
	var request domain.UserUpdateRequest
	repo := hRepository.New(hDb.Get(), &user, c)

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err = ValidateUserUpdateRequest(&request, &user, c)
	if err != nil || c.Response().StatusCode() != 200 {
		return err
	}

	user = domain.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Address:   request.Address,
		UserType:  request.UserType,
	}

	err = repo.Update(&user, id)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, &user)
}
