package user

import (
	"github.com/flambra/account/internal/auth"
	"github.com/flambra/account/pkg/helpers"
	"github.com/flambra/account/pkg/validate"
	"github.com/gofiber/fiber/v2"
)

func ValidateUserCreateRequest(request *UserCreateRequest, c *fiber.Ctx) error {
	if request.TaxNumber == "" {
		return helpers.BadRequestResponse(c, "inform cpf")
	}

	if request.Phone == "" {
		return helpers.BadRequestResponse(c, "inform phone")
	}

	if request.Email == "" {
		return helpers.BadRequestResponse(c, "inform email")
	}

	err := validateEmail(request.Email, c)
	if err != nil || c.Response().StatusCode() != 200 {
		return err
	}

	request.Phone, err = validate.Cellphone(request.Phone)
	if err != nil {
		return helpers.BadRequestResponse(c, err.Error())
	}

	request.TaxNumber, err = validate.CPF(request.TaxNumber)
	if err != nil {
		return helpers.BadRequestResponse(c, err.Error())
	}

	err = validatePassword(request, c)
	if err != nil || c.Response().StatusCode() != 200 {
		return err
	}
	return nil
}

// func ValidateUserUpdateRequest(request *UserUpdateRequest, user *domain.User, c *fiber.Ctx) error {
// 	if request.Mail == "" {
// 		request.Mail = user.Mail
// 	}

// 	if request.BusinessUnitID == 0 {
// 		request.BusinessUnitID = user.BusinessUnitID
// 	}

// 	var bu domain.BusinessUnit
// 	buRepo := repository.New(&bu, c)

// 	err := buRepo.GetById(request.BusinessUnitID)
// 	if err != nil {
// 		return helpers.HandleGormQueryError(c, err, "bu")
// 	}

// 	err = validateBu(&bu, c)
// 	if err != nil || c.Response().StatusCode() != 200 {
// 		return err
// 	}

// 	err = validateEmail(request.Mail, &bu, c)
// 	if err != nil || c.Response().StatusCode() != 200 {
// 		return err
// 	}

// 	if request.TaxNumber != "" {
// 		request.TaxNumber, err = validate.CPF(request.TaxNumber)
// 		if err != nil {
// 			return helpers.BadRequestResponse(c, err.Error())
// 		}
// 	}

// 	if request.Phone != "" {
// 		request.Phone, err = validate.Cellphone(request.Phone)
// 		if err != nil {
// 			return helpers.BadRequestResponse(c, err.Error())
// 		}
// 	}
// 	return nil
// }

func validateEmail(email string, c *fiber.Ctx) error {
	err := validate.Email(email)
	if err != nil {
		return helpers.BadRequestResponse(c, "invalid email")
	}
	return nil
}

func validatePassword(request *UserCreateRequest, c *fiber.Ctx) error {
	if request.Password == "" {
		return nil
	}
	enteredPass := request.Password
	err := auth.ValidatePassword(enteredPass)
	if err != nil {
		return helpers.BadRequestResponse(c, err.Error())
	}

	encryptedPassword, err := auth.EncryptPassword(enteredPass)
	if err != nil {
		return helpers.InternalServerErrorResponse(c, "failed to encrypt user password")
	}
	request.Password = encryptedPassword
	return nil
}
