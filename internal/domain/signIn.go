package domain

type AuthSignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthSignInResponse struct {
	Token string `json:"token"`
}
