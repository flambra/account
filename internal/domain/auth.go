package domain

type AuthSignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthSignInResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type AuthTwoFactorGenerateRequest struct {
	Token  string `json:"token"`
	Method string `json:"method"`
}

type AuthTwoFactorValidateRequest struct {
	Token string `json:"token"`
	Code  string `json:"code"`
}
