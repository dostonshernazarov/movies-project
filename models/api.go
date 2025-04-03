package models

import "time"

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token     string `json:"token"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type MovieRequest struct {
	Title    string  `json:"title" binding:"required"`
	Director string  `json:"director" binding:"required"`
	Year     int     `json:"year" binding:"required"`
	Plot     string  `json:"plot"`
	Genre    string  `json:"genre"`
	Rating   float32 `json:"rating"`
}

type MovieCreateResponse struct {
	Title     string  `json:"title"`
	Director  string  `json:"director"`
	Year      int     `json:"year"`
	Plot      string  `json:"plot"`
	Genre     string  `json:"genre"`
	Rating    float32 `json:"rating"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type MovieResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Director  string    `json:"director"`
	Year      int       `json:"year"`
	Plot      string    `json:"plot"`
	Genre     string    `json:"genre"`
	Rating    float32   `json:"rating"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Movies struct {
	Movies []MovieResponse `json:"movies"`
}

type UserRegisterResponse struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
