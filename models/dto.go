package models

// UserRegisterRequest represents user registration data
type UserRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

// UserLoginRequest represents user login data
type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AuthResponse represents the authentication response
type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// MovieRequest represents movie creation/update data
type MovieRequest struct {
	Title    string  `json:"title" binding:"required"`
	Director string  `json:"director" binding:"required"`
	Year     int     `json:"year" binding:"required"`
	Plot     string  `json:"plot"`
	Genre    string  `json:"genre"`
	Rating   float32 `json:"rating"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}
