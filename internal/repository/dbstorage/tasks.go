package dbstorage

import (
	"context"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/EkaterinaSerikova/todo-list/internal/domain/models"
	"github.com/EkaterinaSerikova/todo-list/pkg/logger"
)

// расширение функциональности DBStorage, добавление методов для работы с задачами в PostgreSQL

func (d *DBStorage) GetTasks() ([]models.Task, error) {
	log := logger.Get()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := d.db.Query(ctx, "SELECT * FROM tasks")
	if err != nil {
		log.Error().Err(err).Msg("failed to get tasks from db")
		return nil, err
	}

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(
			&task.UID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt, &task.DoneAt); err != nil {
			log.Error().Err(err).Msg("failed to parse tasks from db")
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (d *DBStorage) GetTask(id string) (models.Task, error) {
	log := logger.Get()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var task models.Task
	row := d.db.QueryRow(ctx, "SELECT * FROM tasks WHERE uid = $1", id)
	err := row.Scan(&task.UID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt, &task.DoneAt)
	if err != nil {
		log.Error().Err(err).Msg("failed to get task from db")
		return models.Task{}, err
	}
	return task, nil
}

func (d *DBStorage) SaveTask(task models.Task) error {
	log := logger.Get()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := d.db.Exec(
		ctx,
		"INSERT INTO tasks (uid, title, description, created_at, updated_at, done_at, creator_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		task.UID, task.Title, task.Description, task.Status, task.CreatedAt, task.UpdatedAt, task.DoneAt, task.CreatorId)
	if err != nil {
		log.Error().Err(err).Msg("failed to save task to db")
		return err
	}

	return nil
}

func (d *DBStorage) SaveTasks(tasks []models.Task) error {
	log := logger.Get()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := d.db.Begin(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Debug().Err(err).Msg("failed to rollback transaction")
		}
	}()

	_, err = tx.Prepare(ctx, "save_task", "INSERT INTO tasks (tid, title, description, status, creator_id) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		log.Error().Err(err).Msg("failed to prepare statement")
		return err
	}

	for _, task := range tasks {
		_, err := tx.Exec(ctx, "save_task", task.UID, task.Title, task.Description, task.Status, task.CreatorId)
		if err != nil {
			log.Error().Err(err).Msg("failed to save task to db")
			return err
		}
	}

	return tx.Commit(ctx)
}

func (d *DBStorage) UpdateTask(task models.Task) error {
	log := logger.Get()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := d.db.Exec(ctx, "UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = $4 WHERE uid = $5",
		task.Title, task.Description, task.Status, task.UpdatedAt, task.UID)
	if err != nil {
		log.Error().Err(err).Msg("failed to update task in db")
		return err
	}
	return nil
}

func (d *DBStorage) DeleteTask(id string) error {
	log := logger.Get()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := d.db.Exec(ctx, "DELETE FROM tasks WHERE uid = $1", id)
	if err != nil {
		log.Error().Err(err).Msg("failed to delete task from db")
		return err
	}
	return nil
}
