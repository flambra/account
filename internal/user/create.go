package user

import (
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/account/internal/profile"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {

	var user domain.User
	var request domain.UserCreateRequest
	userRepo := hRepository.New(hDb.Get(), &user, c)

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err := ValidateUserCreateRequest(&request, c)
	if err != nil || c.Response().StatusCode() != 200 {
		return err
	}

	if err = userRepo.GetWhere(fiber.Map{"email": request.Email, "tax_number": request.TaxNumber}); err == nil {
		return hResp.StatusConflict(c, &user, "Email or Cpf already in use")
	}

	user = domain.User{
		FirstName:      request.FirstName,
		LastName:       request.LastName,
		TaxNumber:      request.TaxNumber,
		Email:          request.Email,
		HashedPassword: request.Password,
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
