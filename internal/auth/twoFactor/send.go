package twoFactor

import (
	"os"

	"github.com/flambra/account/internal/auth/token"
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hAuth"
	"github.com/flambra/helpers/hReq"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

func Send(c *fiber.Ctx) error {
	var request domain.AuthTwoFactorGenerateRequest

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	claims, err := token.Validate(request.Token)
	if err != nil {
		return hResp.UnauthorizedResponse(c, err.Error())
	}

	code := GenerateCode()

	url := os.Getenv("SENDER_URL")
	authoritazion, err := hAuth.GetToken()
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}
	senderRequest := domain.SenderRequest{
		Variables: map[string]interface{}{
			"pin_code": code,
		},
	}

	switch request.Method {
	case "email":
		senderRequest.To = claims.Email
		senderRequest.TemplateName = os.Getenv("SENDER_EMAIL_TEMPLATE_NAME")
		url += "/email/send"
	case "sms":
		senderRequest.To = claims.Phone
		senderRequest.TemplateName = os.Getenv("SENDER_SMS_TEMPLATE_NAME")
		url += "/sms/send"
	default:
		return hResp.BadRequestResponse(c, "Invalid method")
	}

	req := hReq.Request{
		Url:           url,
		Authorization: authoritazion,
		Body:          senderRequest,
	}

	_, err = req.Post()
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	if err := Update(*claims, code); err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, fiber.Map{
		"message": "2FA code sent successfully",
	})
}
