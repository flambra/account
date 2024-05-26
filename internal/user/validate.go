package user

import (
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hPassword"
	"github.com/flambra/helpers/hResp"
	"github.com/flambra/account/hValidate"
	"github.com/gofiber/fiber/v2"
)

func ValidateUserCreateRequest(request *UserCreateRequest, c *fiber.Ctx) error {
	if request.TaxNumber == "" {
		return hResp.BadRequestResponse(c, "inform cpf")
	}

	if request.Phone == "" {
		return hResp.BadRequestResponse(c, "inform phone")
	}

	if request.Email == "" {
		return hResp.BadRequestResponse(c, "inform email")
	}

	err := validateEmail(request.Email, c)
	if err != nil || c.Response().StatusCode() != 200 {
		return err
	}

	request.Phone, err = hValidate.Cellphone(request.Phone)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	request.TaxNumber, err = hValidate.CPF(request.TaxNumber)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err = validatePassword(request, c)
	if err != nil || c.Response().StatusCode() != 200 {
		return err
	}
	return nil
}

func ValidateUserUpdateRequest(request *UserUpdateRequest, user *domain.User, c *fiber.Ctx) error {
	if request.Email == "" {
		request.Email = user.Email
	}

	err := validateEmail(request.Email, c)
	if err != nil || c.Response().StatusCode() != 200 {
		return err
	}

	if request.TaxNumber != "" {
		request.TaxNumber, err = hValidate.CPF(request.TaxNumber)
		if err != nil {
			return hResp.BadRequestResponse(c, err.Error())
		}
	}

	if request.Phone != "" {
		request.Phone, err = hValidate.Cellphone(request.Phone)
		if err != nil {
			return hResp.BadRequestResponse(c, err.Error())
		}
	}
	return nil
}

func validateEmail(email string, c *fiber.Ctx) error {
	err := hValidate.Email(email)
	if err != nil {
		return hResp.BadRequestResponse(c, "invalid email")
	}
	return nil
}

func validatePassword(request *UserCreateRequest, c *fiber.Ctx) error {
	if request.Password == "" {
		return nil
	}
	enteredPass := request.Password
	err := hPassword.Validate(enteredPass)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	encryptedPassword, err := hPassword.Encrypt(enteredPass)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, "failed to encrypt user password")
	}
	request.Password = encryptedPassword
	return nil
}
