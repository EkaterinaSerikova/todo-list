package dbstorage

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	repositoryErrors "github.com/EkaterinaSerikova/todo-list/internal/domain/errors"
	"github.com/EkaterinaSerikova/todo-list/internal/domain/models"
	"github.com/EkaterinaSerikova/todo-list/pkg/logger"
)

// расширение функциональности DBStorage, добавление методов для работы с пользователями в PostgreSQL

func (d *DBStorage) LoginUser(user models.UserRequest) (models.User, error) {
	log := logger.Get()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var dbUser models.User
	row := d.db.QueryRow(ctx, "SELECT * FROM users WHERE login = $1", user.Login)
	if err := row.Scan(&dbUser.UID, &dbUser.Name, &dbUser.Login, &dbUser.Password); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.User{}, repositoryErrors.ErrUserNotFound
		}
		log.Error().Err(err).Msg("failed to get user from db")
		return models.User{}, err
	}
	return dbUser, nil
}

func (d *DBStorage) RegisterUser(user models.User) (string, error) {
	log := logger.Get()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := d.db.Exec(
		ctx,
		"INSERT INTO users (uid, name, login, password) VALUES ($1, $2, $3, $4)",
		user.UID, user.Name, user.Login, user.Password,
	)
	if err != nil {
		log.Debug().Any("err", err).Str("error", err.Error()).Msg("error form pgx.Exec")
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			log.Debug().Any("pgErr", pgErr).Msg("error form pgx.PgError")
			if pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
				return "", repositoryErrors.ErrUserAlreadyExists
			}
		}
		return "", err
	}
	return user.UID, nil
}

func (d *DBStorage) GetUsers() ([]models.User, error) {
	log := logger.Get()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := d.db.Query(ctx, "SELECT * FROM users")
	if err != nil {
		log.Error().Err(err).Msg("failed to get users from db")
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UID, &user.Name, &user.Login, &user.Password); err != nil {
			log.Error().Err(err).Msg("failed to parse users from db")
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (d *DBStorage) GetUser(uid string) (models.User, error) {
	log := logger.Get()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	row := d.db.QueryRow(ctx, "SELECT * FROM users WHERE uid = $1", uid)
	err := row.Scan(&user.UID, &user.Name, &user.Login, &user.Password)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user from db")
		return models.User{}, err
	}
	return user, nil
}

func (d *DBStorage) UpdateUser(task models.User) error {
	log := logger.Get()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := d.db.Exec(
		ctx,
		"UPDATE users SET name = $1, login = $2, password = $3 WHERE uid = $4",
		task.Name, task.Login, task.Password, task.UID,
	)
	if err != nil {
		log.Error().Err(err).Msg("failed to update user in db")
		return err
	}
	return nil
}

func (d *DBStorage) DeleteUser(uid string) error {
	log := logger.Get()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := d.db.Exec(ctx, "DELETE FROM users WHERE uid = $1", uid)
	if err != nil {
		log.Error().Err(err).Msg("failed to delete user from db")
		return err
	}
	return nil
}
