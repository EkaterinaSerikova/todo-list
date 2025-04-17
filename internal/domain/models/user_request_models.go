package models

type UserRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Login    string `json:"login" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}
