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

func Create(c *fiber.Ctx) error {
	var request domain.UserCreateRequest
	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err := ValidateUserCreateRequest(&request, c)
	if err != nil || c.Response().StatusCode() != 200 {
		return err
	}

	var user domain.User
	repo := hRepository.New(hDb.Get(), &user, c)
	if err = repo.GetWhere(fiber.Map{"email": request.Email}); err == nil {
		return hResp.StatusConflict(c, &user, "Email or Cpf already in use")
	}

	hashedPassword, err := hPassword.Encrypt(request.Password)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	user = domain.User{
		FirstName:      request.FirstName,
		LastName:       request.LastName,
		Email:          request.Email,
		HashedPassword: hashedPassword,
	}

	if err := repo.Create(); err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	if user.ID == 0 {
		return hResp.InternalServerErrorResponse(c, "Fail to create user.")
	}

	err = profile.Create(domain.Profile{UserID: user.ID}, c)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	user.Profile = &domain.Profile{UserID: user.ID}

	return hResp.SuccessCreated(c, &user)
}
