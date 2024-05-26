package domain

import "github.com/golang-jwt/jwt"

type AuthTwoFactorGenerateRequest struct {
	Token  string `json:"token"`
	Method string `json:"method"`
}

type AuthTwoFactorValidateRequest struct {
	Token string `json:"token"`
	Code  string `json:"code"`
}

// UserClaims represents the JWT claims
type UserClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	jwt.StandardClaims
}
