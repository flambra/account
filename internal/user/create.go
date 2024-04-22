package user

import (
	"log"

	"github.com/flambra/account/database"
	"github.com/flambra/account/internal/access/domain"
	"github.com/flambra/account/internal/auth"
	"github.com/gofiber/fiber/v2"
)

type UserCreateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	TaxNumber string `json:"tax_number"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	UserType  string `json:"user_type"`
}

func Create(c *fiber.Ctx) error {

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

	hashedPassword, err := auth.EncryptPassword(request.Password)
	if err != nil {
		return err
	}

	var count int64
	db := database.GetDB()
	db.Model(&domain.User{}).Where("email = ? or tax_number = ?", request.Email, request.TaxNumber).Count(&count)
	if count > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Email or Cpf already in use",
		})
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
	}

	if err := db.Create(&user).Error; err != nil {
		log.Fatalln(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating user in the database",
		})
	}

	profile := domain.Profile{
        UserID: user.ID,
    }

    if err := db.Create(&profile).Error; err != nil {
        log.Fatalln(err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error creating profile in the database",
        })
    }

	user.Profile.UserID = user.ID

	return c.Status(fiber.StatusCreated).JSON(&user)
}
