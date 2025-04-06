package memstorage

import (
	"github.com/EkaterinaSerikova/todo-list/internal/domain/errors"
	"github.com/EkaterinaSerikova/todo-list/internal/domain/models"
)

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

func (m *MemStorage) LoginUser(user models.UserRequest) (models.User, error) {
	for _, us := range m.users {
		if us.Login == user.Login {
			return us, nil
		}
	}
	return models.User{}, errors.ErrUserNotFound
}

func (m *MemStorage) RegisterUser(user models.User) (string, error) {
	_, ok := m.users[user.UID]
	if ok {
		return "", errors.ErrUserAlreadyExists
	}

	m.users[user.UID] = user
	return user.UID, nil
}

func (m *MemStorage) GetUsers() ([]models.User, error) {
	var users []models.User
	if len(m.users) == 0 {
		return nil, errors.ErrEmptyUsersList
	}
	for id, user := range m.users {
		user.UID = id
		users = append(users, user)
	}
	return users, nil
}

func (m *MemStorage) GetUser(id string) (models.User, error) {
	user, ok := m.users[id]
	if !ok {
		return models.User{}, errors.ErrUserNotFound
	}
	return user, nil
}

func (m *MemStorage) SaveUser(user models.User) error {
	for _, t := range m.users {
		if t.Login == user.Login {
			return errors.ErrUserAlreadyExists
		}
	}
	m.users[user.UID] = user
	return nil
}

func (m *MemStorage) UpdateUser(user models.User) error {
	_, ok := m.users[user.UID]
	if !ok {
		return errors.ErrUserNotFound
	}
	m.users[user.UID] = user
	return nil
}

func (m *MemStorage) DeleteUser(id string) error {
	_, ok := m.users[id]
	if !ok {
		return errors.ErrUserNotFound
	}
	delete(m.users, id)
	return nil
}
