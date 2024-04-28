package user

import (
	"encoding/json"
	"strconv"

	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hPassword"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

type UserUpdateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	TaxNumber string `json:"tax_number"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	UserType  string `json:"user_type"`
}

func Update(c *fiber.Ctx) error {
	rawId := c.Params("id")
	if rawId == "" {
		return hResp.BadRequestResponse(c, "inform id")
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	var user domain.User
	var request UserUpdateRequest
	userRepo := hRepository.New(hDb.Get(), &user, c)

	err = userRepo.GetById(id)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err = ValidateUserUpdateRequest(&request, &user, c)
	if err != nil || c.Response().StatusCode() != 200 {
		return err
	}

	serialized, err := json.Marshal(&request)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	err = json.Unmarshal(serialized, &user)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	// var count int64
	// db.Model(&domain.User{}).Where("(email = ? or tax_number = ?) and id != ?", request.Email, request.TaxNumber, user.ID).Count(&count)
	// if count > 0 {
	// 	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
	// 		"error": "Email or Cpf already in use",
	// 	})
	// }

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

	err = userRepo.Save()
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, &user)
}
