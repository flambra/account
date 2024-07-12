package profile

import (
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/flambra/helpers/hTypes"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Find(c *fiber.Ctx) error {
	// var response []ProfilePublicResponse
	var profiles []domain.Profile
	var filter FindProfileFilter

	repo := hRepository.New(hDb.Get(), &profiles, c)
	profilePaginator := hRepository.BuildPaginator(&profiles)

	err := c.QueryParser(profilePaginator)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err = c.QueryParser(&filter)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err = repo.FindAllPaginating(&filter, profilePaginator)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, profilePaginator)

}

type FindProfileFilter struct {
	// Name        string     `query:"name"`
	// Email       string     `query:"email"`
	// TaxNumber   string     `query:"tax_number"`
	// Phone       string     `query:"phone"`
	InitialDate hTypes.Date `query:"initialDate"`
	FinalDate   hTypes.Date `query:"finalDate"`
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
