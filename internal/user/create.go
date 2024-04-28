package user

import (
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/account/internal/profile"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hPassword"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
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
	userRepo := hRepository.New(hDb.Get(), &user, c)

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err := ValidateUserCreateRequest(&request, c)
	if err != nil || c.Response().StatusCode() != 200 {
		return err
	}

	var count int64
	db := hDb.Get()
	db.Model(&domain.User{}).Where("email = ? or tax_number = ?", request.Email, request.TaxNumber).Count(&count)
	if count > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Email or Cpf already in use",
		})
	}

	hashedPassword, err := hPassword.Encrypt(request.Password)
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
	}

	err = userRepo.Create()
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	err = profile.Create(domain.Profile{UserID: user.ID}, c)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	user.Profile.UserID = user.ID

	return hResp.SuccessCreated(c, &user)
}
