package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	task_errors "github.com/EkaterinaSerikova/todo-list/internal/domain/errors"
	"github.com/EkaterinaSerikova/todo-list/internal/domain/models"
)

// HTTP-обработчики для работы с задачами

func (s *ServerApi) getTasks(c *gin.Context) {
	tasks, err := s.tService.GetTasks()
	if err != nil {
		if errors.Is(err, task_errors.ErrTaskNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	if len(tasks) == 0 {
		c.JSON(http.StatusOK, []models.Task{})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (s *ServerApi) createTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindBodyWithJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	creatorId, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := s.tService.CreateTask(task, creatorId); err != nil {
		if errors.Is(err, task_errors.ErrConflict) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (s *ServerApi) getTaskById(c *gin.Context) {
	uid := c.Param("id")
	tasks, err := s.tService.GetTasks()
	if err != nil {
		if errors.Is(err, task_errors.ErrTaskNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	for _, task := range tasks {
		if task.UID == uid {
			c.JSON(http.StatusOK, task)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": task_errors.ErrTaskNotFound.Error()})
}

func (s *ServerApi) updateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks, err := s.tService.GetTasks()
	if err != nil {
		if errors.Is(err, task_errors.ErrTaskNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	for i, task := range tasks {
		if task.UID == id {
			tasks[i].Title = updatedTask.Title
			tasks[i].Description = updatedTask.Description
			tasks[i].Status = updatedTask.Status
			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": task_errors.ErrTaskNotFound.Error()})
}

func (s *ServerApi) deleteTask(c *gin.Context) {
	id := c.Param("id")

	tasks, err := s.tService.GetTasks()
	if err != nil {
		if errors.Is(err, task_errors.ErrTaskNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	for i, task := range tasks {
		if task.UID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			if err := s.tService.CreateTask(task, task.CreatorId); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save tasks"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": task_errors.ErrTaskNotFound.Error()})
}

func (s *ServerApi) saveTasks(c *gin.Context) {
	var tasks []models.Task
	if err := c.ShouldBindJSON(&tasks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.tService.SaveTasks(tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "tasks saved")
}
