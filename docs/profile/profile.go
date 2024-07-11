package docs

import (
	_ "github.com/flambra/account/internal/domain"
	_ "github.com/flambra/helpers/hResp"
)

// Read godoc
//
//	@Summary		Get a profile
//	@Description	Get a profile by ID
//	@Tags			Profile
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Profile ID"
//	@Success		200	{object}	domain.Profile
//	@Failure		400	{object}	hResp.DefaultResponse
//	@Failure		500	{object}	hResp.DefaultResponse
//	@Router			/profile/{id} [get]
func Read() {}

// Update godoc
//
//	@Summary		Update a profile
//	@Description	Update a profile's details by ID
//	@Tags			Profile
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Profile ID"
//	@Param			profile	body		domain.ProfileUpdateRequest	true	"Profile Update Request"
//	@Success		200		{object}	domain.Profile
//	@Failure		400		{object}	hResp.DefaultResponse
//	@Failure		500		{object}	hResp.DefaultResponse
//	@Router			/profile/{id} [put]
func Update() {}
// 	List godoc
//
//	@Summary		List profiles
//	@Description	List all profiles
//	@Tags			Profile
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		domain.Profile
//	@Failure		400	{object}	hResp.DefaultResponse
//	@Failure		500	{object}	hResp.DefaultResponse
//	@Router			/profile [get]
func List() {}