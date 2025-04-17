package memstorage

import (
	"github.com/EkaterinaSerikova/todo-list/internal/domain/errors"
	"github.com/EkaterinaSerikova/todo-list/internal/domain/models"
)

func (m *MemStorage) LoginUser(user models.UserRequest) (models.User, error) {
	for _, us := range m.users {
		if us.Login == user.Login {
			return us, nil
		}
	}
	return models.User{}, errors.ErrUserNotFound
}

func (m *MemStorage) RegisterUser(user models.User) (string, error) {
	for _, existingUser := range m.users {
		if existingUser.Name == user.Name {
			return "", errors.ErrUserAlreadyExists
		}
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
