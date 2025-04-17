package memstorage

import (
	"github.com/EkaterinaSerikova/todo-list/internal/domain/models"
)

// реализует in-memory хранилище - альтернатива dbstorage, когда PostgreSQL недоступен
// хранит данные в оперативной памяти
// реализация тех же методов, что и у dbstorage, но без SQL

type MemStorage struct {
	tasks map[string]models.Task
	users map[string]models.User
}

func New() *MemStorage {
	return &MemStorage{
		tasks: make(map[string]models.Task),
		users: make(map[string]models.User),
	}
}
