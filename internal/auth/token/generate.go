package token

import (
	"os"
	"strconv"
	"time"

	"github.com/flambra/account/internal/domain"
	"github.com/golang-jwt/jwt"
)

func Generate(user domain.User) (string, error) {
	secretKey := os.Getenv("TOKEN_SECRET_KEY")
	durationStr := os.Getenv("TOKEN_DURATION")

	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(time.Minute * time.Duration(duration))
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"phone":   user.Phone,
		"exp":     expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}
