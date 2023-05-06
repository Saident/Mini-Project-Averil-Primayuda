package model

type LoginResponse struct {
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}