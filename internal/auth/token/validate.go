package token

import (
	"fmt"
	"os"

	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hError"
	"github.com/golang-jwt/jwt"
)

func Validate(tokenString string) (*domain.UserClaims, error) {
	secretKey := os.Getenv("TOKEN_SECRET_KEY")
	token, err := jwt.ParseWithClaims(tokenString, &domain.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, hError.New("Invalid or expired token")
	}

	claims, ok := token.Claims.(*domain.UserClaims)
	if !ok || !token.Valid {
		return nil, hError.New("Invalid token claims")
	}

	return claims, nil
}
