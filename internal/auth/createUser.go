package auth

import (
	"github.com/flambra/helpers/hResp"
	"github.com/flambra/helpers/hToken"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	data := map[string]interface{}{
		"message": "Token for Create User",
	}

	access, err := hToken.Create(data)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	response := hToken.Access{
		Token: access.Token,
	}

	return hResp.SuccessResponse(c, &response)
}
