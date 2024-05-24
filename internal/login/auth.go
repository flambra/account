package login

import (
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hAuth"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	var user domain.User
	var request domain.LoginAuthRequest
	repo := hRepository.New(hDb.Get(), &user, c)

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err := repo.GetWhere(fiber.Map{"email": request.Email})
	if err != nil {
		return hResp.UnauthorizedResponse(c, "Invalid email or password")
	}

	userAuth := hAuth.User{
		ID:    int(user.ID), // remove int()
		Email: user.Email,
	}

	token, err := hAuth.GenerateJWT(userAuth)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	response := domain.LoginAuthResponse{
		Token: token,
	}

	return hResp.SuccessResponse(c, response)
}
