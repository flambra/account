package profile

import (
	"encoding/json"
	"strconv"

	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

func Update(c *fiber.Ctx) error {
	rawId := c.Params("id")
	if rawId == "" {
		return hResp.BadRequestResponse(c, "inform id")
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	var profile domain.Profile
	var request domain.ProfileUpdateRequest
	repo := hRepository.New(hDb.Get(), &profile, c)

	err = repo.GetById(id)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	serialized, err := json.Marshal(&request)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	err = json.Unmarshal(serialized, &profile)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	// result := db.Find(&user, "id = ?", request.ID)
	// if result.Error != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": "Database error",
	// 	})
	// }
	// if result.RowsAffected == 0 {
	// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	// 		"message": "Client not found",
	// 	})
	// }

	err = repo.Save()
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, &profile)
}
