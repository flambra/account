package user

import (
	"log"

	"github.com/flambra/account/database"
	"github.com/flambra/account/internal/access/domain"
	"github.com/flambra/account/internal/auth"
	"github.com/gofiber/fiber/v2"
)

func Update(c *fiber.Ctx) error {
	db := database.GetDB()

	var user domain.User
	var request UserCreateRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing client data",
		})
	}

	err := ValidateUserCreateRequest(&request, c)
	if err != nil || c.Response().StatusCode() != 200 {
		return err
	}

	var count int64
	db.Model(&domain.User{}).Where("(email = ? or tax_number = ?) and id != ?", request.Email, request.TaxNumber, user.ID).Count(&count)
	if count > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Email or Cpf already in use",
		})
	}

	hashedPassword, err := auth.EncryptPassword(request.Password)
	if err != nil {
		return err
	}

	user = domain.User{
		FirstName:      request.FirstName,
		LastName:       request.LastName,
		TaxNumber:      request.TaxNumber,
		Email:          request.Email,
		HashedPassword: hashedPassword,
		Phone:          request.Phone,
		Address:        request.Address,
		UserType:       request.UserType,
		Profile: domain.Profile{
			Description:     "",
			Skills:          "",
			Portfolio:       "",
			Specializations: "",
			Availability:    true,
			Languages:       "",
			Location:        "",
		},
	}

	if err := db.Save(&user).Error; err != nil {
		log.Fatalln(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error updating user in the database",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&user)
}
