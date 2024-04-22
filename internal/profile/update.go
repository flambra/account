package user

import (
	"encoding/json"
	"log"

	"github.com/flambra/account/database"
	"github.com/flambra/account/internal/access/domain"
	"github.com/flambra/account/internal/auth"
	"github.com/flambra/account/pkg/helpers"
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
}

func Update(c *fiber.Ctx) error {
	var user domain.User
	var request ProfileUpdateRequest

	if err := c.QueryParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error parsing query",
		})
	}

	db := database.GetDB()

	result := db.Find(&user, "id = ?", request.ID)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database error",
		})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Client not found",
		})
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing client data",
		})
	}

	err := ValidateUserUpdateRequest(&request, &user, c)
	if err != nil || c.Response().StatusCode() != 200 {
		return err
	}

	serialized, err := json.Marshal(&request)
	if err != nil {
		return helpers.InternalServerErrorResponse(c, err.Error())
	}

	err = json.Unmarshal(serialized, &user)
	if err != nil {
		return helpers.InternalServerErrorResponse(c, err.Error())
	}

	if request.Password != "" {
		hashedPassword, err := auth.EncryptPassword(request.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error encrypting password",
			})
		}
		request.HashedPassword = string(hashedPassword)
	}

	user = domain.User{
		ID: request.ID,
	}

	if err := db.Model(&user).Updates(request).Error; err != nil {
		log.Fatalln(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error updating user in the database",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&request)
}
