package auth

type AuthResponse struct {
	SessionToken string `json:"session_token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
