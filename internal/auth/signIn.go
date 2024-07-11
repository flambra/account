package auth

import (
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/flambra/helpers/hToken"
	"github.com/gofiber/fiber/v2"
)

func SignIn(c *fiber.Ctx) error {
	var user domain.User
	var request domain.AuthSignInRequest
	repo := hRepository.New(hDb.Get(), &user, c)

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err := repo.GetWhere(fiber.Map{"email": request.Email})
	if err != nil {
		return hResp.UnauthorizedResponse(c, "Invalid email or password")
	}

	data := map[string]interface{}{
		"user_id": user.ID,
		"email":   user.Email,
		"phone":   user.Phone,
	}

	access, err := hToken.Create(data)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	response := hToken.Access{
		Token:        access.Token,
		RefreshToken: access.RefreshToken,
	}

	return hResp.SuccessResponse(c, &response)
}
