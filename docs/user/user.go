package docs

import (
	_ "github.com/flambra/account/internal/domain"
	_ "github.com/flambra/helpers/hRepository"
	_ "github.com/flambra/helpers/hResp"
)

//	@title				Flambra Account API
//	@description		This API is for the Flambra Account service.
//	@externalDocs.url	https://github.com/flambra/account

// Create godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user with the provided details
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string						false	"Bearer <access token>"
//	@Param			user			body		domain.UserCreateRequest	true	"User Create Request"
//	@Success		201				{object}	domain.User
//	@Failure		400				{object}	hResp.DefaultResponse
//	@Failure		409				{object}	hResp.DefaultResponse
//	@Failure		500				{object}	hResp.DefaultResponse
//	@Router			/user [post]
func Create() {}

// Delete godoc
//
//	@Summary		Delete a user
//	@Description	Delete a user by ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	false	"Bearer <access token>"
//	@Param			id				path		int		true	"User ID"
//	@Success		200				{object}	domain.User
//	@Failure		400				{object}	hResp.DefaultResponse
//	@Failure		500				{object}	hResp.DefaultResponse
//	@Router			/user/{id} [delete]
func Delete() {}

// Read godoc
//
//	@Summary		Get a user
//	@Description	Get a user by ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	false	"Bearer <access token>"
//	@Param			id				path		int		true	"User ID"
//	@Success		200				{object}	domain.User
//	@Failure		400				{object}	hResp.DefaultResponse
//	@Failure		500				{object}	hResp.DefaultResponse
//	@Router			/user/{id} [get]
func Read() {}

// Update godoc
//
//	@Summary		Update a user
//	@Description	Update a user's details by ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string						false	"Bearer <access token>"
//	@Param			id				path		int							true	"User ID"
//	@Param			user			body		domain.UserUpdateRequest	true	"User Update Request"
//	@Success		200				{object}	domain.User
//	@Failure		400				{object}	hResp.DefaultResponse
//	@Failure		500				{object}	hResp.DefaultResponse
//	@Router			/user/{id} [put]
func Update() {}

// Page godoc
//
//	@Summary		Get a page
//	@Description	Get a page of user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	false	"Bearer <access token>"
//	@Param			page			query		int		false	"Page number"
//	@Param			limit			query		int		false	"Page limit"
//	@Param			sort			query		string	false	"Sort by"
//	@Param			name			query		string	false	"Search by Name"
//	@Param			address			query		string	false	"Search by Address"
//	@Param			usertype		query		string	false	"Search by UserType"
//	@Success		200				{object}	hRepository.Paginator
//	@Failure		400				{object}	hResp.DefaultResponse
//	@Failure		500				{object}	hResp.DefaultResponse
//	@Router			/user [get]
func Page() {}
