package twoFactor

import (
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
)

func Update(claims domain.UserClaims, code string) error {
	var user domain.User
	repo := hRepository.New(hDb.Get(), &user, nil)

	user = domain.User{
		LastCode: code,
	}

	err := repo.Update(user, claims.UserID)
	if err != nil {
		return err
	}

	return nil
}
