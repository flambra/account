package docs

import (
	_ "github.com/flambra/account/internal/domain"
	_ "github.com/flambra/helpers/hResp"
	_ "github.com/flambra/helpers/hToken"
)

// SignIn godoc
//
//	@Summary		Sign in a user
//	@Description	Sign in a user with email and password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			auth	body		domain.AuthSignInRequest	true	"Auth Sign In Request"
//	@Success		200		{object}	hToken.Access				"Access token generated successfully"
//	@Failure		400		{object}	hResp.DefaultResponse
//	@Failure		401		{object}	hResp.DefaultResponse
//	@Failure		500		{object}	hResp.DefaultResponse
//	@Router			/auth/signin [post]
func SignIn() {}

// RefreshToken godoc
//
//	@Summary		Refresh an access token
//	@Description	Refresh an access token using a refresh token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string			false	"Bearer <access token>"
//	@Param			request			body		hToken.Access	false	"Refresh Token Request"
//	@Success		200				{object}	hToken.Access
//	@Failure		400				{object}	hResp.DefaultResponse
//	@Failure		401				{object}	hResp.DefaultResponse
//	@Failure		500				{object}	hResp.DefaultResponse
//	@Router			/auth/refreshtoken [post]
func RefreshToken() {}

// Send godoc
//
//	@Summary		Send 2FA code
//	@Description	Send a 2FA code to the user's email or phone
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string								false	"Bearer <access token>"
//	@Param			request			body		domain.AuthTwoFactorGenerateRequest	true	"TwoFactor Generate Request"
//	@Success		200				{object}	hResp.DefaultResponse				"2FA code sent successfully"
//	@Failure		400				{object}	hResp.DefaultResponse
//	@Failure		401				{object}	hResp.DefaultResponse
//	@Failure		500				{object}	hResp.DefaultResponse
//	@Router			/auth/twofactor/send [post]
func Send() {}

// Validate godoc
//
//	@Summary		Validate 2FA code
//	@Description	Validate the 2FA code provided by the user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string								false	"Bearer <access token>"
//	@Param			request			body		domain.AuthTwoFactorValidateRequest	true	"TwoFactor Validate Request"
//	@Success		200				{object}	map[string]string					"2FA code validated successfully"
//	@Failure		400				{object}	hResp.DefaultResponse
//	@Failure		401				{object}	hResp.DefaultResponse
//	@Failure		500				{object}	hResp.DefaultResponse
//	@Router			/auth/twofactor/validate [post]
func Validate() {}
