package profile

import (
	"github.com/flambra/account/database"
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/http"
	"github.com/flambra/helpers/repository"
	"github.com/flambra/helpers/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Find(c *fiber.Ctx) error {
	// var response []ProfilePublicResponse
	var profiles []domain.Profile
	var filter FindProfileFilter

	profileRepo := repository.New(database.GetDB(), &profiles, c)
	profilePaginator := repository.BuildPaginator(&profiles)

	err := c.QueryParser(profilePaginator)
	if err != nil {
		return http.BadRequestResponse(c, err.Error())
	}

	err = c.QueryParser(&filter)
	if err != nil {
		return http.BadRequestResponse(c, err.Error())
	}

	err = profileRepo.FindAllPaginating(&filter, profilePaginator)
	if err != nil {
		return http.InternalServerErrorResponse(c, err.Error())
	}

	return http.SuccessResponse(c, profilePaginator)

}

type FindProfileFilter struct {
	// Name        string     `query:"name"`
	// Email       string     `query:"email"`
	// TaxNumber   string     `query:"tax_number"`
	// Phone       string     `query:"phone"`
	InitialDate types.Date `query:"initialDate"`
	FinalDate   types.Date `query:"finalDate"`
}

func (d *FindProfileFilter) Apply(db *gorm.DB) *gorm.DB {
	if !d.InitialDate.Default().IsZero() {
		db = db.Where("created_at >= ?", d.InitialDate.Default())
	}

	if !d.FinalDate.Default().IsZero() {
		db = db.Where("created_at < ?", d.FinalDate.Default())
	}

	// if d.Name != "" {
	// 	likeQuery := "%" + d.Name + "%"
	// 	db = db.Where("name ILIKE ?", likeQuery)
	// }

	// if d.Email != "" {
	// 	likeQuery := "%" + d.Email + "%"
	// 	db = db.Where("email ILIKE ?", likeQuery)
	// }

	// if d.TaxNumber != "" {
	// 	likeQuery := "%" + d.TaxNumber + "%"
	// 	db = db.Where("tax_number ILIKE ?", likeQuery)
	// }

	// if d.Phone != "" {
	// 	likeQuery := "%" + d.Phone + "%"
	// 	db = db.Where("phone ILIKE ?", likeQuery)
	// }

	return db
}
