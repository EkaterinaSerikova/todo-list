package models

import "time"

type Task struct {
	UID         string    `json:"uid" validate:"required" check:"unique"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Status      string    `json:"status" validate:"required,oneof=Новая 'В процессе' Завершена"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DoneAt      time.Time `json:"done_at"`
	CreatorId   string    `json:"creator_id"`
}
