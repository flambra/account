package domain

type LoginAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginAuthResponse struct {
	Token string `json:"token"`
}
