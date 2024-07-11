package twoFactor

import (
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hError"
	"github.com/flambra/helpers/hRepository"
)

func Update(claims map[string]interface{}, code string) error {
	var user domain.User
	repo := hRepository.New(hDb.Get(), &user, nil)

	user = domain.User{
		LastCode: code,
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return hError.New("UserID not found in token data")
	}

	err := repo.Update(user, int(userID))
	if err != nil {
		return err
	}

	return nil
}
