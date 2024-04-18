package database

import (
	"log"

	"github.com/flambra/account/internal/access/domain"
)

func Migrate() error {
	err := GetDB().AutoMigrate(
		&domain.User{},
		&domain.Profile{},
	)
	if err != nil {
		return err
	}

	log.Println("Database migration completed successfully.")
	return nil
}
