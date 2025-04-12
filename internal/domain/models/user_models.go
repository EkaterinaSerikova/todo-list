package models

type User struct {
	UID      string `json:"uid" check:"unique"`
	Name     string `json:"name" validate:"required"`
	Login    string `json:"login" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
