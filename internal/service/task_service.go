package service

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/EkaterinaSerikova/todo-list/internal/domain/models"
	"github.com/EkaterinaSerikova/todo-list/pkg/logger"
)

// реализация бизнес-логики для работы с задачами

// определение контракта для работы с хранилищем (БД или memory-хранилищем)
type Repository interface {
	GetTasks() ([]models.Task, error)
	GetTask(string) (models.Task, error)
	SaveTask(task models.Task) error
	SaveTasks([]models.Task) error
	UpdateTask(task models.Task) error
	DeleteTask(string) error

	LoginUser(models.UserRequest) (models.User, error)
	RegisterUser(user models.User) (string, error)

	GetUsers() ([]models.User, error)
	GetUser(string) (models.User, error)
	UpdateUser(task models.User) error
	DeleteUser(string) error
}

type TaskService struct {
	repo  Repository
	valid *validator.Validate
}

func NewTaskService(repo Repository) *TaskService {
	valid := validator.New()
	return &TaskService{repo: repo, valid: valid}
}

func (t *TaskService) CreateTask(task models.Task, creator_id string) error {
	tID := uuid.New().String()
	task.UID = tID
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now
	task.CreatorId = creator_id
	err := t.repo.SaveTask(task)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskService) GetTasks() ([]models.Task, error) {
	tasks, err := t.repo.GetTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *TaskService) GetTaskByID(id string) (*models.Task, error) {
	task, err := t.repo.GetTask(id)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *TaskService) UpdateTask(task models.Task) error {
	err := t.repo.UpdateTask(task)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskService) DeleteTask(id string) error {
	err := t.repo.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskService) SaveTasks(tasks []models.Task) error {
	log := logger.Get()
	for index := range tasks {
		id := uuid.New().String()
		tasks[index].UID = id
	}

	log.Debug().Any("tasks", tasks).Msg("tasks for saving to db")

	return t.repo.SaveTasks(tasks)
}
