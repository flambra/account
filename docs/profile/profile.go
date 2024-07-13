package docs

import (
	_ "github.com/flambra/account/internal/domain"
	_ "github.com/flambra/helpers/hResp"
)

//	@title				Flambra Account API
//	@description		This API is for the Flambra Account service.
//	@externalDocs.url	https://github.com/flambra/account

// Read godoc
//
//	@Summary		Get a profile
//	@Description	Get a profile by ID
//	@Tags			Profile
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	false	"Bearer <access token>"
//	@Param			id				path		int		true	"Profile ID"
//	@Success		200				{object}	domain.Profile
//	@Failure		400				{object}	hResp.DefaultResponse
//	@Failure		500				{object}	hResp.DefaultResponse
//	@Router			/profile/{id} [get]
func Read() {}

// Update godoc
//
//	@Summary		Update a profile
//	@Description	Update a profile's details by ID
//	@Tags			Profile
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string						false	"Bearer <access token>"
//	@Param			id				path		int							true	"Profile ID"
//	@Param			profile			body		domain.ProfileUpdateRequest	true	"Profile Update Request"
//	@Success		200				{object}	domain.Profile
//	@Failure		400				{object}	hResp.DefaultResponse
//	@Failure		500				{object}	hResp.DefaultResponse
//	@Router			/profile/{id} [put]
func Update() {}

// List godoc
//
//	@Summary		List profiles
//	@Description	List all profiles
//	@Tags			Profile
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	false	"Bearer <access token>"
//	@Success		200				{array}		domain.Profile
//	@Failure		400				{object}	hResp.DefaultResponse
//	@Failure		500				{object}	hResp.DefaultResponse
//	@Router			/profile [get]
func List() {}
