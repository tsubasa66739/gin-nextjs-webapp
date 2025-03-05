package dto

type SignupInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"passeord" binding:"required,min=8"`
}
