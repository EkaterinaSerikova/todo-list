package models

import (
	"time"
)

type Task struct {
	UID         string    `json:"uid" validate:"required" check:"unique"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Status      string    `json:"status" validate:"required,oneof=Новая 'В процессе' Завершена"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DoneAt      time.Time `json:"done_at"`
}

type User struct {
	UID      string `json:"uid" check:"unique"`
	Name     string `json:"name" validate:"required"`
	Login    string `json:"login" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Login    string `json:"login" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}
