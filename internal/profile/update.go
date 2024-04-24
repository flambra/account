package profile

import (
	"encoding/json"
	"strconv"

	"github.com/flambra/account/database"
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/http"
	"github.com/flambra/helpers/repository"
	"github.com/gofiber/fiber/v2"
)

type ProfileUpdateRequest struct {
	UserID          uint
	Description     string `json:"description,omitempty"`
	Skills          string `json:"skills,omitempty"`
	Portfolio       string `json:"portfolio,omitempty"`
	Specializations string `json:"specializations,omitempty"`
	Availability    string `json:"availability,omitempty"`
	Languages       string `json:"languages,omitempty"`
	Location        string `json:"location,omitempty"`
	PhoneType       string `json:"phone_type,omitempty"`
}

func Update(c *fiber.Ctx) error {
	rawId := c.Params("id")
	if rawId == "" {
		return http.BadRequestResponse(c, "inform id")
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		return http.BadRequestResponse(c, err.Error())
	}

	var profile domain.Profile
	var request ProfileUpdateRequest
	profileRepo := repository.New(database.GetDB(), &profile, c)

	err = profileRepo.GetById(id)
	if err != nil {
		return http.InternalServerErrorResponse(c, err.Error())
	}

	if err := c.BodyParser(&request); err != nil {
		return http.BadRequestResponse(c, err.Error())
	}

	serialized, err := json.Marshal(&request)
	if err != nil {
		return http.InternalServerErrorResponse(c, err.Error())
	}

	err = json.Unmarshal(serialized, &profile)
	if err != nil {
		return http.InternalServerErrorResponse(c, err.Error())
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

	err = profileRepo.Save()
	if err != nil {
		return http.InternalServerErrorResponse(c, err.Error())
	}

	return http.SuccessResponse(c, &profile)
}
