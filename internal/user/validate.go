package user

import (
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hPassword"
	"github.com/flambra/helpers/hResp"
	"github.com/flambra/helpers/hValidate"
	"github.com/gofiber/fiber/v2"
)

func ValidateUserCreateRequest(request *domain.UserCreateRequest, c *fiber.Ctx) error {
	if request.TaxNumber == "" {
		return hResp.BadRequestResponse(c, "inform cpf")
	}

	if request.Phone == "" {
		return hResp.BadRequestResponse(c, "inform phone")
	}

	if request.Email == "" {
		return hResp.BadRequestResponse(c, "inform email")
	}

	err := hValidate.BirthDate(request.BirthDate)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err = validateEmail(request.Email, c)
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

	err = validatePassword(request.Password, c)
	if err != nil || c.Response().StatusCode() != 200 {
		return err
	}
	return nil
}

func ValidateUserUpdateRequest(request *domain.UserUpdateRequest, user *domain.User, c *fiber.Ctx) error {
	if request.Email != "" {
		request.Email = user.Email
		err := validateEmail(request.Email, c)
		if err != nil || c.Response().StatusCode() != 200 {
			return err
		}
	}

	var err error
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

	if request.Password != "" {
		err = validatePassword(request.Password, c)
		if err != nil || c.Response().StatusCode() != 200 {
			return err
		}
	}

	if !request.BirthDate.IsZero() {
		err := hValidate.BirthDate(request.BirthDate)
		if err != nil {
			return hResp.BadRequestResponse(c, err.Error())
		}
		user.BirthDate = request.BirthDate
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

func validatePassword(password string, c *fiber.Ctx) error {
	if password == "" {
		return nil
	}
	enteredPass := password
	err := hPassword.Validate(enteredPass)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	encryptedPassword, err := hPassword.Encrypt(enteredPass)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, "failed to encrypt user password")
	}
	password = encryptedPassword
	return nil
}
