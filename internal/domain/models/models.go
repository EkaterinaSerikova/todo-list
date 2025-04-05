package models

import (
	"time"
)

type Task struct {
	UID         string    `json:"uid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DoneAt      time.Time `json:"done_at"`
}

type User struct {
	UID      string `json:"uid"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
