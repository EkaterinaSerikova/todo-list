package memstorage

import (
	"github.com/EkaterinaSerikova/todo-list/internal/domain/errors"
	"github.com/EkaterinaSerikova/todo-list/internal/domain/models"
)

func (m *MemStorage) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	if len(m.tasks) == 0 {
		return nil, errors.ErrEmptyTasksList
	}
	for id, task := range m.tasks {
		task.UID = id
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (m *MemStorage) GetTask(id string) (models.Task, error) {
	task, ok := m.tasks[id]
	if !ok {
		return models.Task{}, errors.ErrTaskNotFound
	}
	return task, nil
}

func (m *MemStorage) SaveTask(task models.Task) error {
	for _, t := range m.tasks {
		if t.Title == task.Title {
			return errors.ErrTaskAlreadyExist
		}
	}
	m.tasks[task.UID] = task
	return nil
}

func (m *MemStorage) UpdateTask(task models.Task) error {
	_, ok := m.tasks[task.UID]
	if !ok {
		return errors.ErrTaskNotFound
	}
	m.tasks[task.UID] = task
	return nil
}

func (m *MemStorage) DeleteTask(id string) error {
	_, ok := m.tasks[id]
	if !ok {
		return errors.ErrTaskNotFound
	}
	delete(m.tasks, id)
	return nil
}
